一招教你无阻塞读写通道
----------------

### 文件介绍
```
├── block
│   ├── read_buffer_block.go
│   ├── read_unbuffer_block.go
│   ├── write_buufer_block.go
│   └── write_unbuffer_block.go
├── select_timer_unblock.go
└── select_unblock.go
```

- `block`：目录，4种情况的阻塞示例代码
- `select_unblock.go`：使用select的default语句实现无阻塞读写
- `select_timer_unblock.go`：使用select和超时实现无阻塞读写


正文
-----------------

介绍Golang并发的模型写了几篇了，但一直没有以**channel**为主题进行介绍，今天就给大家聊一聊channel，channel的基本使用非常简单，想必大家都已了解，所以直接来个进阶点的：**介绍channel的阻塞情况，以及给你一个必杀技，立马解决阻塞问题，实用性高**。



### 阻塞场景

**无论是有缓存通道、无缓冲通道都存在阻塞的情况**。阻塞场景共4个，有缓存和无缓冲各2个。

**无缓冲通道**的特点是，发送的数据需要被读取后，发送才会完成，它**阻塞场景**：

1. 通道中无数据，但执行读通道。
2. 通道中无数据，向通道写数据，但无协程读取。


```go
// 场景1
func ReadNoDataFromNoBufCh() {
	noBufCh := make(chan int)

	<-noBufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// 场景2
func WriteNoBufCh() {
	ch := make(chan int)

	ch <- 1
	fmt.Println("write success no block")
	
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
```

*注：示例代码中的Output注释代表函数的执行结果，每一个函数都由于阻塞在通道操作而无法继续向下执行，最后报了死锁错误。*

**有缓存通道**的特点是，有缓存时可以向通道中写入数据后直接返回，缓存中有数据时可以从通道中读到数据直接返回，这时有缓存通道是不会阻塞的，它**阻塞场景是**：

1. 通道的缓存无数据，但执行读通道。
2. 通道的缓存已经占满，向通道写数据，但无协程读。

```go
// 场景1
func ReadNoDataFromBufCh() {
	bufCh := make(chan int, 1)

	<-bufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}

// 场景2
func WriteBufChButFull() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100

	ch <- 1
	fmt.Println("write success no block")
	
	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
```



### 使用Select实现无阻塞读写

select是执行选择操作的一个结构，它里面有一组case语句，它会执行其中无阻塞的那一个，如果都阻塞了，那就等待其中一个不阻塞，进而继续执行，**它有一个default语句，该语句是永远不会阻塞的，我们可以借助它实现无阻塞的操作**。如果不了解，不想多了解一下select可以先看下这2篇文章：
- [Golang并发模型：轻松入门select](https://mp.weixin.qq.com/s/ACh-TGlPo72r4e6pbh52vg)
- [Golang并发模型：select进阶](https://mp.weixin.qq.com/s/ZfBcxvqiyks_s7cAD-zGCw)

下面**示例代码是使用select修改后的无缓冲通道和有缓冲通道的读写**，以下函数可以直接通过main函数调用，其中的Ouput的注释是运行结果，从结果能看出，在通道不可读或者不可写的时候，不再阻塞等待，而是直接返回。

```go
// 无缓冲通道读
func ReadNoDataFromNoBufChWithSelect() {
	bufCh := make(chan int)

	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}

	// Output:
	// channel has no data
}

// 有缓冲通道读
func ReadNoDataFromBufChWithSelect() {
	bufCh := make(chan int, 1)

	if v, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("read: %d\n", v)
	}

	// Output:
	// channel has no data
}

// select结构实现通道读
func ReadWithSelect(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
		return x, nil
	default:
		return 0, errors.New("channel has no data")
	}
}

// 无缓冲通道写
func WriteNoBufChWithSelect() {
	ch := make(chan int)
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}

	// Output:
	// channel blocked, can not write
}

// 有缓冲通道写
func WriteBufChButFullWithSelect() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100
	if err := WriteChWithSelect(ch); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("write success")
	}

	// Output:
	// channel blocked, can not write
}

// select结构实现通道写
func WriteChWithSelect(ch chan int) error {
	select {
	case ch <- 1:
		return nil
	default:
		return errors.New("channel blocked, can not write")
	}
}
```

### 使用Select+超时改善无阻塞读写

**使用default实现的无阻塞通道阻塞有一个缺陷：当通道不可读或写的时候，会即可返回**。实际场景，更多的需求是，我们希望尝试读一会数据，或者尝试写一会数据，如果实在没法读写再返回，程序继续做其它的事情。

**使用定时器替代default**可以解决这个问题，**给通道增加读写数据的容忍时间**，如果500ms内无法读写，就即刻返回。示例代码修改一下会是这样：

```go
func ReadWithSelect(ch chan int) (x int, err error) {
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case x = <-ch:
		return x, nil
	case <-timeout.C:
		return 0, errors.New("read time out")
	}
}

func WriteChWithSelect(ch chan int) error {
	timeout := time.NewTimer(time.Microsecond * 500)

	select {
	case ch <- 1:
		return nil
	case <-timeout.C:
		return errors.New("write time out")
	}
}
```

结果就会变成超时返回：

```text
read time out
write time out
read time out
write time out
```



### 示例源码

本文所有示例源码存储在Github，点击阅读原文查看：https://github.com/Shitaibin/golang_step_by_step/tree/master/channel/unblock_channel



这篇文章了channel的阻塞情况，以及解决阻塞的2种办法：
1. 使用select的default语句，在channel不可读写时，即可返回
2. 使用select+定时器，在超时时间内，channel不可读写，则返回

希望这篇文章对你的channel读写有所启发。


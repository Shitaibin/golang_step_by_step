# 目录介绍
- master：主要代码
- fan_model_slow：介绍fan不一定能提升性能
- optimize_fan_model：优化FAN模式的代码
- buffer_size：对buffer缓冲区大小的测试

========================

前一篇文章[《Golang并发模型：轻松入门流水线模型》](https://segmentfault.com/a/1190000017142506)，介绍了流水线模型的概念，这篇文章是流水线模型进阶，介绍FAN-IN和FAN-OUT，FAN模式可以让我们的流水线模型更好的利用Golang并发，提高软件性能。但FAN模式不一定是万能，不见得能提高程序的性能，甚至还不如普通的流水线。我们先介绍下FAN模式，再看看它怎么提升性能的，它是不是万能的。



# FAN-IN和FAN-OUT模式

Golang的并发模式灵感来自现实世界，这些模式是通用的，毫无例外，FAN模式也是对当前世界的模仿。**以汽车组装为例，汽车生产线上有个阶段是给小汽车装4个轮子，可以把这个阶段任务交给4个人同时去做，这4个人把轮子都装完后，再把汽车移动到生产线下一个阶段。这个过程中，就有任务的分发，和任务结果的收集。其中任务分发是FAN-OUT，任务收集是FAN-IN。**

- **FAN-OUT模式：多个goroutine从同一个通道读取数据，直到该通道关闭。**OUT是一种张开的模式，所以又被称为扇出，可以用来分发任务。
- **FAN-IN模式：1个goroutine从多个通道读取数据，直到这些通道关闭。**IN是一种收敛的模式，所以又被称为扇入，用来收集处理的结果。

![fan-in和fan-out.png](https://upload-images.jianshu.io/upload_images/10901752-727b047a9808439d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


# FAN-IN和FAN-OUT实践

我们这次试用FAN-OUT和FAN-IN，解决[《Golang并发模型：轻松入门流水线模型》](https://segmentfault.com/a/1190000017142506?_ea=5178632)中提到的问题：计算一个整数切片中元素的平方值并把它打印出来。

- `producer()`保持不变，负责生产数据。
- `squre()`也不变，负责计算平方值。
- 修改`main()`，启动3个square，这3个squre从producer生成的通道读数据，**这是FAN-OUT**。
- 增加`merge()`，入参是3个square各自写数据的通道，给这3个通道分别启动1个协程，把数据写入到自己创建的通道，并返回该通道，**这是FAN-IN**。

[FAN模式流水线示例](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/master/hi_fan_example.go)：
```go
package main

import (
	"fmt"
	"sync"
)

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
    // FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer(1, 2, 3, 4)
	
    // FAN-OUT
    c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d ", ret)
	}
	fmt.Println()
}
```

3个squre协程**并发**运行，结果顺序是无法确定的，所以你得到的结果，不一定与下面的相同。

```go
➜  awesome git:(master) ✗ go run hi.go
  1   4  16   9 
```

# FAN模式真能提升性能吗？

相信你心里已经有了答案，可以的。我们还是使用老问题，对比一下简单的流水线和FAN模式的流水线，修改下代码，增加程序的执行时间：

- `produer()`使用参数生成指定数量的数据。
- `square()`增加阻塞操作，睡眠1s，模拟阶段的运行时间。
- `main()`关闭对结果数据的打印，降低结果处理时的IO对FAN模式的对比。

[普通流水线](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/master/hi_simple.go)：
```go
// hi_simple.go

package main

import (
	"fmt"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
            // simulate
            time.Sleep(time.Second)
		}
	}()

	return out
}

func main() {
	in := producer(10)
	ch := square(in)

	// consumer
	for _ = range ch {
	}
}
```


使用[FAN模式的流水线](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/master/hi_fan.go)：
```go
// hi_fan.go
package main

import (
	"sync"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
			// simulate
			time.Sleep(time.Second)
		}
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := producer(10)

	// FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for _ = range merge(c1, c2, c3) {
	}
}
```



多次测试，每次结果近似，结果如下：

- FAN模式利用了7%的CPU，而普通流水线CPU只使用了3%，**FAN模式能够更好的利用CPU，提供更好的并发，提高Golang程序的并发性能。**
- FAN模式耗时10s，普通流水线耗时4s。**在协程比较费时时，FAN模式可以减少程序运行时间，同样的时间，可以处理更多的数据。**

```bash
➜  awesome git:(master) ✗ time go run hi_simple.go
go run hi_simple.go  0.17s user 0.18s system 3% cpu 10.389 total
➜  awesome git:(master) ✗ 
➜  awesome git:(master) ✗ time go run hi_fan.go
go run hi_fan.go  0.17s user 0.16s system 7% cpu 4.288 total
```



**也可以使用Benchmark进行测试，看2个类型的执行时间，结论相同**。为了节约篇幅，这里不再介绍，[方法和结果贴在Gist](https://gist.github.com/Shitaibin/9593a18989b6c81bb3aae5ccdf9b6470)了，想看的朋友瞄一眼，或自己动手搞搞。



# FAN模式一定能提升性能吗？

FAN模式可以提高并发的性能，那我们是不是可以都使用FAN模式？

不行的，因为**FAN模式不一定能提升性能。**

依然使用之前的问题，再次修改下代码，其他不变：

- `squre()`去掉耗时。
- `main()`增加producer()的入参，让producer生产10,000,000个数据。

[简单版流水线修改代码](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/fan_model_slow/hi_simple.go)：
```go
// hi_simple.go

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func main() {
	in := producer(10000000)
	ch := square(in)

	// consumer
	for _ = range ch {
	}
}
```


[FAN模式流水线修改代码](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/fan_model_slow/hi_fan.go)：
```go
// hi-fan.go
package main

import (
	"sync"
)

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()

	return out
}

func main() {
	in := producer(10000000)

	// FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for _ = range merge(c1, c2, c3) {
	}
}
```

结果，可以跑多次，结果近似：

```bash
➜  awesome git:(master) ✗ time go run hi-simple.go    
go run hi-simple.go  9.96s user 5.93s system 168% cpu 9.424 total
➜  awesome git:(master) ✗ time go run hi-fan.go        
go run hi-fan.go  23.35s user 11.51s system 297% cpu 11.737 total
```

从这个结果，我们能看到2点。

- FAN模式可以提高CPU利用率。
- **FAN模式不一定能提升效率，降低程序运行时间。**

# 优化FAN模式

既然FAN模式不一定能提高性能，如何优化？

**不同的场景优化不同，要依具体的情况，解决程序的瓶颈。**

我们当前程序的瓶颈在FAN-IN，squre函数很快就完成，merge函数它把3个数据写入到1个通道的时候出现了瓶颈，**适当使用带缓冲通道可以提高程序性能**，[再修改下代码](https://github.com/Shitaibin/golang_pipeline_step_by_step/blob/optimize_fan_model/hi_fan_buffered.go)

- `merge()`中的`out`修改为：

  ```go
  out := make(chan int, 100)
  ```

结果：

```bash
➜  awesome git:(master) ✗ time go run hi_fan-buffer.go 
go run hi-fan-buffer.go  19.85s user 8.19s system 323% cpu 8.658 total
```

使用带缓存通道后，程序的性能有了较大提升，**CPU利用率提高到323%，提升了8%，运行时间从11.7降低到8.6，降低了26%。**


FAN模式的特点很简单，相信你已经掌握了，如果记不清了[看这里](#FAN-IN和FAN-OUT模式)，本文所有代码浏览该[Github仓库]()。

FAN模式很有意思，并且能提高Golang并发的性能，如果想以后运用自如，用到自己的项目中去，还是要写写自己的Demo，快去实践一把。


下一篇，写流水线中协程的“优雅退出”，欢迎关注。


> 如果这篇文章对你有帮助，请点个赞/喜欢，让我知道我的写作是有价值的，感谢。


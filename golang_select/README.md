之前的文章都提到过，**Golang的并发模型都来自生活，select也不例外**。举个例子：我们都知道一句话，“吃饭睡觉打豆豆”，这一句话里包含了3件事：
1. 妈妈喊你吃饭，你去吃饭。
2. 时间到了，要睡觉。
3. 没事做，打豆豆。

在Golang里，select就是干这个事的：到吃饭了去吃饭，该睡觉了就睡觉，没事干就打豆豆。

结束发散，我们看下select的功能，以及它能做啥。

### select功能
**在多个通道上进行读或写操作，让函数可以处理多个事情，但1次只处理1个。以下特性也都必须熟记于心：**
1. 每次执行select，都会只执行其中1个case或者执行default语句。
2. 当没有case或者default可以执行时，select则阻塞，等待直到有1个case可以执行。
2. 当有多个case可以执行时，则随机选择1个case执行。
4. `case`后面跟的必须是读或者写通道的操作，否则编译出错。

select长下面这个样子，由`select`和`case`组成，`default`不是必须的，如果没其他事可做，可以省略`default`。
```golang
func main() {
	readCh := make(chan int, 1)
	writeCh := make(chan int, 1)

	y := 1
	select {
	case x := <-readCh:
		fmt.Printf("Read %d\n", x)
	case writeCh <- y:
		fmt.Printf("Write %d\n", y)
	default:
		fmt.Println("Do what you want")
	}
}
```

我们创建了`readCh`和`writeCh`2个通道：
1. `readCh`中没有数据，所以`case x := <-readCh`读不到数据，所以这个case不能执行。
2. `writeCh`是带缓冲区的通道，它里面是空的，可以写入1个数据，所以`case writeCh <- y`可以执行。
3. 有`case`可以执行，所以`default`不会执行。

这个测试的结果是
```bash
$ go run example.go
Write 1
```

### 用打豆豆实践select

来，我们看看select怎么实现打豆豆：`eat()`函数会启动1个协程，该协程先睡几秒，事件不定，然后喊你吃饭，`main()`函数中的`sleep`是个定时器，每3秒喊你吃1次饭，`select`则处理3种情况：
1. 从`eatCh`中读到数据，代表有人喊我吃饭，我要吃饭了。
2. 从`sleep.C`中读到数据，代表闹钟时间到了，我要睡觉。
3. `default`是，没人喊我吃饭，也不到时间睡觉，我就打豆豆。

```golang
import (
	"fmt"
	"time"
	"math/rand"
)

func eat() chan string {
	out := make(chan string)
	go func (){
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "Mom call you eating"
		close(out)
	}()
	return out
}


func main() {
	eatCh := eat()
	sleep := time.NewTimer(time.Second * 3)
	select {
	case s := <-eatCh:
		fmt.Println(s)
	case <- sleep.C:
		fmt.Println("Time to sleep")
	default:
		fmt.Println("Beat DouDou")
	}
}
```

由于前2个case都要等待一会，所以都不能执行，所以执行`default`，运行结果一直是打豆豆：
```bash
$ go run x.go
Beat DouDou
```
现在我们不打豆豆了，你把`default`和下面的打印注释掉，多运行几次，有时候会吃饭，有时候会睡觉，比如这样：
```bash
$ go run x.go
Mom call you eating
$ go run x.go
Time to sleep
$ go run x.go
Time to sleep
```


**select很简单但功能很强大，它让golang的并发功能变的更强大。这篇文章写的啰嗦了点，重点是为下一篇文章做铺垫，下一篇我们将介绍下select的高级用法。**

**select的应用场景很多，让我总结一下，放在下一篇文章中吧。**


### 并发系列文章推荐

- [Golang并发模型：轻松入门流水线模型](https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483671&idx=1&sn=1706ffa6deee44a367c34ef84448f55f&scene=21#wechat_redirect)
-  [Golang并发模型：轻松入门流水线FAN模式](https://mp.weixin.qq.com/s?__biz=Mzg3MTA0NDQ1OQ==&mid=2247483680&idx=1&sn=de463ebbd088c0acf6c2f0b5f179f38d&scene=21#wechat_redirect)
-  [Golang并发模型：并发协程的优雅退出](https://mp.weixin.qq.com/s/RjomKnfwCTy7tC9gbpPxCQ)
- [Golang并发模型：轻松入门select](https://mp.weixin.qq.com/s/ACh-TGlPo72r4e6pbh52vg)

> 1. 如果这篇文章对你有帮助，请点个赞/喜欢，鼓励我持续分享，感谢。
> 2. [我的文章列表，点此可查看](http://lessisbetter.site/2018/12/11/gongzhonghao-articles/)
> 3. 如果喜欢本文，随意转载，但请保留此[原文链接](https://mp.weixin.qq.com/s/ACh-TGlPo72r4e6pbh52vg)。


![一起学Golang-分享有料的Go语言技术](https://upload-images.jianshu.io/upload_images/10901752-0de86c464c34a5f7.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/258/)


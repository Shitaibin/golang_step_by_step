Golang并发模型：轻松入门协程池
-------------------

goroutine是非常轻量的，不会暂用太多资源，基本上有多少任务，我们可以开多少goroutine去处理。但有时候，我们还是想控制一下。

比如，我们有A、B两类工作，不想把太多资源花费在B类务上，而是花在A类任务上。对于A，我们可以来1个开一个goroutine去处理，对于B，我们可以使用一个协程池，协程池里有5个线程去处理B类任务，这样B消耗的资源就不会太多。

控制使用资源并不是协程池目的，**使用协程池是为了更好并发、程序鲁棒性、容错性等**。废话少说，快速入门协程池才是这篇文章的目的。

**协程池指的是预先分配固定数量的goroutine处理相同的任务**，和线程池是类似的，不同点是协程池中处理任务的是协程，线程池中处理任务的是线程。

### 最简单的协程池模型

![简单协程池模型](http://cnd.lessisbetter.site/2018-12-simple-goroutine-pool.png)


上面这个图展示了最简单的协程池的样子。先**把协程池作为一个整体看，它使用2个通道，左边的`jobCh`是任务通道，任务会从这个通道中流进来，右边的`retCh`是结果通道，协程池处理任务后得到的结果会写入这个通道**。至于协程池中，有多少协程处理任务，这是外部不关心的。

看一下协程池内部，图中画了5个goroutine，实际goroutine的数量是依具体情况而定的。**协程池内每个协程都从`jobCh`读任务、处理任务，然后将结果写入到`retCh`。**


### 示例
模型看懂了，看个小例子吧。

![示例代码1](http://cdn.lessisbetter.site/2018-12-goroutine-pool-code-1.png)

`workerPool()`会创建1个简单的协程池，协程的数量可以通入参数`n`执行，并且还指定了`jobCh`和`retCh`两个参数。

`worker()`是协程池中的协程，入参分布是它的ID、job通道和结果通道。使用`for-range`从`jobCh`读取任务，直到`jobCh`关闭，然后一个最简单的任务：生成1个字符串，证明自己处理了某个任务，并把字符串作为结果写入`retCh`。

![示例代码2](http://cdn.lessisbetter.site/2018-12-goroutine-pool-code-2.png)

`main()`启动`genJob`获取存放任务的通道`jobCh`，然后创建`retCh`，它的缓存空间是200，并使用`workerPool`启动一个有5个协程的协程池。1s之后，关闭`retCh`，然后开始从`retCh`中读取协程池处理结果，并打印。

`genJob`启动一个协程，并生产n个任务，写入到`jobCh`。


示例运行结果如下，一共产生了10个任务，显示大部分工作都被`worker 2`这个协程抢走了，如果我们设置的任务成千上万，协程池长时间处理任务，每个协程处理的工作数量就会均衡很多。
```bash
➜ go run simple_goroutine_pool.go
worker 2 processed job: 4
worker 2 processed job: 5
worker 2 processed job: 6
worker 2 processed job: 7
worker 2 processed job: 8
worker 2 processed job: 9
worker 0 processed job: 1
worker 3 processed job: 2
worker 4 processed job: 3
worker 1 processed job: 0
```
### 回顾
最简单的协程池模型就这么简单，再回头看下协程池及周边由哪些组成：
1. **协程池内的一定数量的协程。**
2. **任务队列**，即`jobCh`，存在协程池不能立即处理任务的情况，所以需要队列把任务先暂存。
3. **结果队列**，即`retCh`，同上，协程池处理任务的结果，也存在不能被下游立刻提取的情况，要暂时保存。

**协程池最简要（核心）的逻辑是所有协程从任务读取任务，处理后把结果存放到结果队列。**

### Go并发系列文章

1. [Golang并发模型：轻松入门流水线模型](https://mp.weixin.qq.com/s/YB5XZ5NatniHSYBQ3AHONw)
1. [Golang并发模型：轻松入门流水线FAN模式](https://mp.weixin.qq.com/s/68FGjm7PFN5VbVF0zL-PlQ)
1. [Golang并发模型：并发协程的优雅退出](https://mp.weixin.qq.com/s/RjomKnfwCTy7tC9gbpPxCQ)
1. [Golang并发模型：轻松入门select](https://mp.weixin.qq.com/s/ACh-TGlPo72r4e6pbh52vg)
1. [Golang并发模型：select进阶](https://mp.weixin.qq.com/s/ZfBcxvqiyks_s7cAD-zGCw)
1. [Golang并发模型：轻松入门协程池](https://mp.weixin.qq.com/s/fINhzg3eNi9YFi5qZ_JzGA)

> 1. 如果这篇文章对你有帮助，请点个赞/喜欢，鼓励我持续分享，感谢。
> 2. 如果喜欢本文，随意转载，但请保留此[原文链接](http://lessisbetter.site/2018/12/20/golang-simple-goroutine-pool/)。
> 3. [博客文章列表，点此可查看](http://lessisbetter.site/2018/12/11/gongzhonghao-articles/)




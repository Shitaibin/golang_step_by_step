package main

import "fmt"

// main.main
func main() {
	fmt.Println("Hello scheduler")
}

// Output
// ➜  one_routine git:(master) ✗ go build .
// ➜  one_routine git:(master) ✗ GODEBUG=schedtrace=1000 scheddetail=1 ./one_routine1
// SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler

// 上面各项含义，来自[也谈Gotrouine调度器](https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/)：
// SCHED：调试信息输出标志字符串，代表本行是goroutine scheduler的输出；
// 6016ms：即从程序启动到输出这行日志的时间；
// gomaxprocs: P的数量；
// idleprocs: 处于idle状态的P的数量；通过gomaxprocs和idleprocs的差值，我们就可知道执行go代码的P的数量；
// threads: os threads的数量，包含scheduler使用的m数量，加上runtime自用的类似sysmon这样的thread的数量；
// spinningthreads: 处于自旋状态的os thread数量；
// idlethread: 处于idle状态的os thread的数量；
// runqueue=1： go scheduler全局队列中G的数量；
// [3 4 0 10]: 分别为4个P的local queue中的G的数量。

// 详细版本
// ➜  one_routine1 git:(master) ✗ GODEBUG=schedtrace=1000,scheddetail=1 ./one_routine
// SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=1 schedtick=0 syscalltick=0 m=3 runqsize=0 gfreecnt=0
//   P1: status=1 schedtick=1 syscalltick=0 m=2 runqsize=0 gfreecnt=0
//   P2: status=1 schedtick=0 syscalltick=0 m=4 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=2 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=true blocked=false lockedg=-1
//   M3: p=0 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M2: p=1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=true blocked=false lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=1
//   G1: status=1(chan receive) m=-1 lockedm=0
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
// Hello scheduler

// 解释：
// 1. P0, P1, P2这3个P处于1(_Prunning)，其他P的状态为0(_Pidle)，分别和M0,M2,M4绑定
// 2. M4、M3、M2分别绑定了P2，P0和P1，curg代表当前运行的G，此时没有协程“正在”运行，所以M上的G都为-1，M4和M2处于自旋状态
// 3. G1、G2、G3，状态分别是1、4、4，1代表运行(_Grunning)，4代表等待(_Gwaiting)，处于当前状态的原因是后面的字符串，m都为-1代表不运行在M上，（讲不通，G1在运行，但没有绑定的M）

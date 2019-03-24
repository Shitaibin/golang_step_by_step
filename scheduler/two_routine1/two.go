package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义一个函数
	gr := func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("Hello scheduler")
		}
	}

	// 启动goroutine
	go gr()

	// 等待接收
	time.Sleep(time.Second * 7)
}

// Output
// ➜  one_routine2 git:(master) ✗ go build .
// ➜  one_routine2 git:(master) ✗ GODEBUG=schedtrace=1000  ./one_routine2
// SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
// SCHED 1001ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler
// SCHED 2002ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler
// SCHED 3004ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler
// SCHED 4005ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler
// SCHED 5013ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
// Hello scheduler

// 更详细版本
// ➜  two_routine1 git:(master) ✗ GODEBUG=schedtrace=1000,scheddetail=1 ./two_routine1
// SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=1 schedtick=0 syscalltick=0 m=0 runqsize=0 gfreecnt=0
//   P1: status=1 schedtick=2 syscalltick=0 m=2 runqsize=0 gfreecnt=0
//   P2: status=1 schedtick=0 syscalltick=0 m=3 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M3: p=2 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=true blocked=false lockedg=-1
//   M2: p=1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=0 curg=-1 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 helpgc=0 spinning=false blocked=false lockedg=1
//   G1: status=1(chan receive) m=-1 lockedm=0
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
// Hello scheduler
// SCHED 1006ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=2 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=4 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=1 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=33 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(sleep) m=-1 lockedm=-1
//   G6: status=3() m=0 lockedm=-1 // G6、G33都在系统调用，分别占用了M0和M2，系统调用时M必然是blocked状态？不一定，其他的代码示例就非这样
//   G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
//   G33: status=3() m=2 lockedm=-1 //
// Hello scheduler
// SCHED 2012ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=2 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=4 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=4 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=17 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(sleep) m=-1 lockedm=-1
//   G6: status=3() m=0 lockedm=-1
//   G17: status=3(timer goroutine (idle)) m=2 lockedm=-1
//   G33: status=4(timer goroutine (idle)) m=-1 lockedm=-1
// Hello scheduler
// SCHED 3017ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=2 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=4 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=33 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(sleep) m=-1 lockedm=-1
//   G6: status=3() m=0 lockedm=-1
//   G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
//   G33: status=3(timer goroutine (idle)) m=2 lockedm=-1
// Hello scheduler
// SCHED 4024ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=2 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=17 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(sleep) m=-1 lockedm=-1
//   G6: status=3() m=0 lockedm=-1
//   G17: status=3(timer goroutine (idle)) m=2 lockedm=-1
//   G33: status=4(timer goroutine (idle)) m=-1 lockedm=-1
// Hello scheduler
// SCHED 5035ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=9 m=-1 runqsize=0 gfreecnt=1
//   P2: status=0 schedtick=1 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=6() m=-1 lockedm=-1
//   G6: status=3() m=0 lockedm=-1
//   G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
//   G33: status=4(timer goroutine (idle)) m=-1 lockedm=-1
// SCHED 6047ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=4 syscalltick=9 m=-1 runqsize=0 gfreecnt=1
//   P2: status=0 schedtick=1 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=6 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=6() m=-1 lockedm=-1 // G5退出
//   G6: status=3() m=0 lockedm=-1
//   G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1
//   G33: status=4(timer goroutine (idle)) m=-1 lockedm=-1
// ➜  two_routine1 git:(master) ✗

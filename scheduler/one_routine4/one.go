package main

import (
	"fmt"
	"time"
)

// main.main
func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello scheduler")
	}
}

// Output
// ➜  one_routine2 git:(master) ✗ go build .
// ➜  one_routine2 git:(master) ✗ GODEBUG=schedtrace=1000 ./one_routine2
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

// ➜  one_routine2 git:(master) ✗ GODEBUG=schedtrace=1000,scheddetail=1 ./one_routine2
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
// SCHED 1002ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=2 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   为何所有的M都没绑定P了？M0进行系统调用，释放P这是正确的，因为其他P也没有“运行中的”G，所以就没有把M和P绑定到一起？
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=5 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=3() m=0 lockedm=-1 // 3代表系统调用，和M0绑定，M0也证明绑定了G5
// Hello scheduler
// SCHED 2003ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=2 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=3 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   依然是M0做系统调用，M0和G17绑定了
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=17 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(timer goroutine (idle)) m=-1 lockedm=-1 // G5是定时器
//   G17: status=3() m=0 lockedm=-1
// Hello scheduler
// SCHED 3006ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=9 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=2 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=3 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=5 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=3(timer goroutine (idle)) m=0 lockedm=-1
//   G17: status=4(timer goroutine (idle)) m=-1 lockedm=-1 // G17也是定时器？
// Hello scheduler
// SCHED 4014ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=1 syscalltick=9 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=2 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=6 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=17 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=4(timer goroutine (idle)) m=-1 lockedm=-1
//   G17: status=3(timer goroutine (idle)) m=0 lockedm=-1
// Hello scheduler
// Hello scheduler

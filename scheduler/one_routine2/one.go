package main

import "fmt"

func main() {
	for {
		fmt.Println("Hello scheduler")
	}
}

// 部分结果
// ➜  one_routine2 git:(master) ✗ GODEBUG=schedtrace=1000,scheddetail=1 ./one_routine2 > log
// SCHED 0ms: gomaxprocs=8 idleprocs=6 threads=3 spinningthreads=1 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P1: status=1 schedtick=0 syscalltick=0 m=2 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M2: p=1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=1
//   G1: status=4(chan receive) m=-1 lockedm=0
//   G2: status=1() m=-1 lockedm=-1
//   G3: status=1() m=-1 lockedm=-1
// SCHED 1005ms: gomaxprocs=8 idleprocs=7 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P1: status=2 schedtick=11 syscalltick=240808 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=1 curg=1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   G1: status=3(chan receive) m=0 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G17: status=4(finalizer wait) m=-1 lockedm=-1
// SCHED 2013ms: gomaxprocs=8 idleprocs=7 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P1: status=1 schedtick=16 syscalltick=469074 m=0 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=1 curg=1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   G1: status=3(chan receive) m=0 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G17: status=4(finalizer wait) m=-1 lockedm=-1

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello scheduler")
	time.Sleep(time.Second)
}

// Output
// ➜  one_routine2 git:(master) ✗ go build .
// ➜  one_routine2 git:(master) ✗ GODEBUG=schedtrace=1000,scheddetail=1 ./one_routine2
// SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=4 spinningthreads=1 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   P0: status=1 schedtick=0 syscalltick=0 m=3 runqsize=0 gfreecnt=0 // 只有2个M在运行
//   P1: status=1 schedtick=0 syscalltick=0 m=2 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M3、M2绑定P，但没有G，这是啥状态？P的队列空了
//   M3: p=0 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M2: p=1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=1
//   G1: status=1(chan receive) m=-1 lockedm=0
//   G2: status=1() m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
// Hello scheduler
// SCHED 1001ms: gomaxprocs=8 idleprocs=8 threads=5 spinningthreads=0 idlethreads=3 runqueue=0 gcwaiting=0 nmidlelocked=0 stopwait=0 sysmonwait=0
//   所有P都是Idle状态，闲下来了
//   P0: status=0 schedtick=1 syscalltick=7 m=-1 runqsize=0 gfreecnt=0
//   P1: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P2: status=0 schedtick=1 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P3: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P4: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P5: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P6: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   P7: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
//   M0绑定了G5，定时器为系统调用，并且定时器是运行在新goroutine中的
//   M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M3: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=-1
//   M1: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   M0: p=-1 curg=5 mallocing=0 throwing=0 preemptoff= locks=2 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
//   G1: status=4(sleep) m=-1 lockedm=-1
//   G2: status=4(force gc (idle)) m=-1 lockedm=-1
//   G3: status=4(GC sweep wait) m=-1 lockedm=-1
//   G4: status=4(finalizer wait) m=-1 lockedm=-1
//   G5: status=3() m=0 lockedm=-1

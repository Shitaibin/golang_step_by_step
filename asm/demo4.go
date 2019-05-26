package main

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	<-ch
}

// ➜  asm git:(master) ✗ go tool compile -S demo4.go
// "".main STEXT size=128 args=0x0 locals=0x28
// 	0x0000 00000 (demo4.go:3)	TEXT	"".main(SB), ABIInternal, $40-0
// 	0x0000 00000 (demo4.go:3)	MOVQ	(TLS), CX
// 	0x0009 00009 (demo4.go:3)	CMPQ	SP, 16(CX)
// 	0x000d 00013 (demo4.go:3)	JLS	121
//
// 	0x000f 00015 (demo4.go:3)	SUBQ	$40, SP
// 	0x0013 00019 (demo4.go:3)	MOVQ	BP, 32(SP)
// 	0x0018 00024 (demo4.go:3)	LEAQ	32(SP), BP
//
// 	0x001d 00029 (demo4.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
// 	0x001d 00029 (demo4.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
// 	0x001d 00029 (demo4.go:3)	FUNCDATA	$3, gclocals·1cf923758aae2e428391d1783fe59973(SB)
//  创建channel
// 	0x001d 00029 (demo4.go:4)	PCDATA	$2, $1
// 	0x001d 00029 (demo4.go:4)	PCDATA	$0, $0
// 	0x001d 00029 (demo4.go:4)	LEAQ	type.chan int(SB), AX
// 	0x0024 00036 (demo4.go:4)	PCDATA	$2, $0
// 	0x0024 00036 (demo4.go:4)	MOVQ	AX, (SP)
// 	0x0028 00040 (demo4.go:4)	MOVQ	$0, 8(SP)
// 	0x0031 00049 (demo4.go:4)	CALL	runtime.makechan(SB)
// 	0x0036 00054 (demo4.go:4)	PCDATA	$2, $1
// 	0x0036 00054 (demo4.go:4)	MOVQ	16(SP), AX
// 	0x003b 00059 (demo4.go:4)	PCDATA	$2, $0
// 	0x003b 00059 (demo4.go:4)	PCDATA	$0, $1
// 	0x003b 00059 (demo4.go:4)	MOVQ	AX, "".ch+24(SP)
//  创建goroutine
// 	0x0040 00064 (demo4.go:5)	MOVL	$8, (SP)
// 	0x0047 00071 (demo4.go:5)	PCDATA	$2, $2
// 	0x0047 00071 (demo4.go:5)	LEAQ	"".main.func1·f(SB), CX
// 	0x004e 00078 (demo4.go:5)	PCDATA	$2, $0
// 	0x004e 00078 (demo4.go:5)	MOVQ	CX, 8(SP)
// 	0x0053 00083 (demo4.go:5)	CALL	runtime.newproc(SB)
//  main从channel读，不会处理goroutine
// 	0x0058 00088 (demo4.go:8)	PCDATA	$2, $1
// 	0x0058 00088 (demo4.go:8)	PCDATA	$0, $0
// 	0x0058 00088 (demo4.go:8)	MOVQ	"".ch+24(SP), AX
// 	0x005d 00093 (demo4.go:8)	PCDATA	$2, $0
// 	0x005d 00093 (demo4.go:8)	MOVQ	AX, (SP)
// 	0x0061 00097 (demo4.go:8)	MOVQ	$0, 8(SP)
// 	0x006a 00106 (demo4.go:8)	CALL	runtime.chanrecv1(SB)
//  栈收缩
// 	0x006f 00111 (demo4.go:9)	MOVQ	32(SP), BP
// 	0x0074 00116 (demo4.go:9)	ADDQ	$40, SP
// 	0x0078 00120 (demo4.go:9)	RET
// 	0x0079 00121 (demo4.go:9)	NOP
// 	0x0079 00121 (demo4.go:3)	PCDATA	$0, $-1
// 	0x0079 00121 (demo4.go:3)	PCDATA	$2, $-1
// 	0x0079 00121 (demo4.go:3)	CALL	runtime.morestack_noctxt(SB)
// 	0x007e 00126 (demo4.go:3)	JMP	0
// 	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 6a 48  eH..%....H;a.vjH
// 	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 8d 05  ..(H.l$ H.l$ H..
// 	0x0020 00 00 00 00 48 89 04 24 48 c7 44 24 08 00 00 00  ....H..$H.D$....
// 	0x0030 00 e8 00 00 00 00 48 8b 44 24 10 48 89 44 24 18  ......H.D$.H.D$.
// 	0x0040 c7 04 24 08 00 00 00 48 8d 0d 00 00 00 00 48 89  ..$....H......H.
// 	0x0050 4c 24 08 e8 00 00 00 00 48 8b 44 24 18 48 89 04  L$......H.D$.H..
// 	0x0060 24 48 c7 44 24 08 00 00 00 00 e8 00 00 00 00 48  $H.D$..........H
// 	0x0070 8b 6c 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 80  .l$ H..(........
//  这是上面涉及的函数
// 	rel 5+4 t=16 TLS+0
// 	rel 32+4 t=15 type.chan int+0
// 	rel 50+4 t=8 runtime.makechan+0
// 	rel 74+4 t=15 "".main.func1·f+0
// 	rel 84+4 t=8 runtime.newproc+0
// 	rel 107+4 t=8 runtime.chanrecv1+0
// 	rel 122+4 t=8 runtime.morestack_noctxt+0
//  以下是子goroutine的执行
// "".main.func1 STEXT size=72 args=0x8 locals=0x18
// 	0x0000 00000 (demo4.go:5)	TEXT	"".main.func1(SB), ABIInternal, $24-8
//  统一是goroutine栈guard
// 	0x0000 00000 (demo4.go:5)	MOVQ	(TLS), CX
// 	0x0009 00009 (demo4.go:5)	CMPQ	SP, 16(CX)
// 	0x000d 00013 (demo4.go:5)	JLS	65
//
// 	0x000f 00015 (demo4.go:5)	SUBQ	$24, SP
// 	0x0013 00019 (demo4.go:5)	MOVQ	BP, 16(SP)
// 	0x0018 00024 (demo4.go:5)	LEAQ	16(SP), BP
// 	0x001d 00029 (demo4.go:5)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
// 	0x001d 00029 (demo4.go:5)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
// 	0x001d 00029 (demo4.go:5)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//  向channel写1
// 	0x001d 00029 (demo4.go:6)	PCDATA	$2, $1
// 	0x001d 00029 (demo4.go:6)	PCDATA	$0, $1
// 	0x001d 00029 (demo4.go:6)	MOVQ	"".ch+32(SP), AX
// 	0x0022 00034 (demo4.go:6)	PCDATA	$2, $0
// 	0x0022 00034 (demo4.go:6)	MOVQ	AX, (SP)
// 	0x0026 00038 (demo4.go:6)	PCDATA	$2, $1
// 	0x0026 00038 (demo4.go:6)	LEAQ	"".statictmp_0(SB), AX
// 	0x002d 00045 (demo4.go:6)	PCDATA	$2, $0
// 	0x002d 00045 (demo4.go:6)	MOVQ	AX, 8(SP)
// 	0x0032 00050 (demo4.go:6)	CALL	runtime.chansend1(SB)
//  goroutine的栈收缩
// 	0x0037 00055 (demo4.go:7)	MOVQ	16(SP), BP
// 	0x003c 00060 (demo4.go:7)	ADDQ	$24, SP
// 	0x0040 00064 (demo4.go:7)	RET
//  goroutine栈分裂
// 	0x0041 00065 (demo4.go:7)	NOP
// 	0x0041 00065 (demo4.go:5)	PCDATA	$0, $-1
// 	0x0041 00065 (demo4.go:5)	PCDATA	$2, $-1
// 	0x0041 00065 (demo4.go:5)	CALL	runtime.morestack_noctxt(SB)
// 	0x0046 00070 (demo4.go:5)	JMP	0

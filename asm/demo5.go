package main

func main() {
	// var ch chan int
	ch := make(chan int)
	close(ch)
	<-ch
}

// ➜ asm git:(master) ✗ go tool compile -S demo5.go| grep -A 10 'demo5'
// 0x0000 00000 (demo5.go:3)	TEXT	"".main(SB), ABIInternal, $40-0
// 0x0000 00000 (demo5.go:3)	MOVQ	(TLS), CX
// 0x0009 00009 (demo5.go:3)	CMPQ	SP, 16(CX)
// 0x000d 00013 (demo5.go:3)	JLS	106
// 0x000f 00015 (demo5.go:3)	SUBQ	$40, SP
// 0x0013 00019 (demo5.go:3)	MOVQ	BP, 32(SP)
// 0x0018 00024 (demo5.go:3)	LEAQ	32(SP), BP
// 0x001d 00029 (demo5.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
// 0x001d 00029 (demo5.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
// 0x001d 00029 (demo5.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
// 0x001d 00029 (demo5.go:5)	PCDATA	$2, $1
// 0x001d 00029 (demo5.go:5)	PCDATA	$0, $0
// 0x001d 00029 (demo5.go:5)	LEAQ	type.chan int(SB), AX
// 0x0024 00036 (demo5.go:5)	PCDATA	$2, $0
// 0x0024 00036 (demo5.go:5)	MOVQ	AX, (SP)
// 0x0028 00040 (demo5.go:5)	MOVQ	$0, 8(SP)
// 0x0031 00049 (demo5.go:5)	CALL	runtime.makechan(SB)
// 0x0036 00054 (demo5.go:5)	PCDATA	$2, $1
// 0x0036 00054 (demo5.go:5)	MOVQ	16(SP), AX
// 0x003b 00059 (demo5.go:5)	PCDATA	$0, $1
// 0x003b 00059 (demo5.go:5)	MOVQ	AX, "".ch+24(SP)
//
// 0x0040 00064 (demo5.go:6)	PCDATA	$2, $0
// 0x0040 00064 (demo5.go:6)	MOVQ	AX, (SP)
// 0x0044 00068 (demo5.go:6)	CALL	runtime.closechan(SB)
// 与demo4.go中的<-ch翻译得来的伪汇编是相同的
// 0x0049 00073 (demo5.go:7)	PCDATA	$2, $1
// 0x0049 00073 (demo5.go:7)	PCDATA	$0, $0
// 0x0049 00073 (demo5.go:7)	MOVQ	"".ch+24(SP), AX
// 0x004e 00078 (demo5.go:7)	PCDATA	$2, $0
// 0x004e 00078 (demo5.go:7)	MOVQ	AX, (SP)
// 0x0052 00082 (demo5.go:7)	MOVQ	$0, 8(SP)
// 0x005b 00091 (demo5.go:7)	CALL	runtime.chanrecv1(SB)
//
// 0x0060 00096 (demo5.go:8)	MOVQ	32(SP), BP
// 0x0065 00101 (demo5.go:8)	ADDQ	$40, SP
// 0x0069 00105 (demo5.go:8)	RET
// 0x006a 00106 (demo5.go:8)	NOP
// 0x006a 00106 (demo5.go:3)	PCDATA	$0, $-1
// 0x006a 00106 (demo5.go:3)	PCDATA	$2, $-1
// 0x006a 00106 (demo5.go:3)	CALL	runtime.morestack_noctxt(SB)
// 0x006f 00111 (demo5.go:3)	JMP	0

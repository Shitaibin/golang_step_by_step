package main

func main() {
	go func() {}()
}

// ➜  asm git:(master) ✗ go tool compile -S demo3.go
// "".main STEXT size=70 args=0x0 locals=0x18
// 	0x0000 00000 (demo3.go:3)	TEXT	"".main(SB), ABIInternal, $24-0
// 	0x0000 00000 (demo3.go:3)	MOVQ	(TLS), CX
// 	0x0009 00009 (demo3.go:3)	CMPQ	SP, 16(CX)
// 	0x000d 00013 (demo3.go:3)	JLS	63
//
// 	0x000f 00015 (demo3.go:3)	SUBQ	$24, SP
// 	0x0013 00019 (demo3.go:3)	MOVQ	BP, 16(SP)
// 	0x0018 00024 (demo3.go:3)	LEAQ	16(SP), BP
//
// 	0x001d 00029 (demo3.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x001d 00029 (demo3.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 	0x001d 00029 (demo3.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//
// 	0x001d 00029 (demo3.go:4)	PCDATA	$2, $0
// 	0x001d 00029 (demo3.go:4)	PCDATA	$0, $0
// 	0x001d 00029 (demo3.go:4)	MOVL	$0, (SP)
// 	0x0024 00036 (demo3.go:4)	PCDATA	$2, $1
//  启动1个协程
// 	0x0024 00036 (demo3.go:4)	LEAQ	"".main.func1·f(SB), AX
// 	0x002b 00043 (demo3.go:4)	PCDATA	$2, $0
// 	0x002b 00043 (demo3.go:4)	MOVQ	AX, 8(SP)
// 	0x0030 00048 (demo3.go:4)	CALL	runtime.newproc(SB)
//  栈收缩
// 	0x0035 00053 (demo3.go:5)	MOVQ	16(SP), BP
// 	0x003a 00058 (demo3.go:5)	ADDQ	$24, SP
//
// 	0x003e 00062 (demo3.go:5)	RET
// 	0x003f 00063 (demo3.go:5)	NOP
//
// 	0x003f 00063 (demo3.go:3)	PCDATA	$0, $-1
// 	0x003f 00063 (demo3.go:3)	PCDATA	$2, $-1
// 	0x003f 00063 (demo3.go:3)	CALL	runtime.morestack_noctxt(SB)
// 	0x0044 00068 (demo3.go:3)	JMP	0

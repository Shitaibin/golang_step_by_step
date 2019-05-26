package main

func main() {
	var sl []int
	sl = append(sl, 1)
}

// ➜  asm git:(master) ✗ go tool compile -S demo6_slice1.go
// "".main STEXT size=100 args=0x0 locals=0x48
// 0x0000 00000 (demo6_slice1.go:3)	TEXT	"".main(SB), ABIInternal, $72-0
// 0x0000 00000 (demo6_slice1.go:3)	MOVQ	(TLS), CX
// 0x0009 00009 (demo6_slice1.go:3)	CMPQ	SP, 16(CX)
// 0x000d 00013 (demo6_slice1.go:3)	JLS	93
// 0x000f 00015 (demo6_slice1.go:3)	SUBQ	$72, SP
// 0x0013 00019 (demo6_slice1.go:3)	MOVQ	BP, 64(SP)
// 0x0018 00024 (demo6_slice1.go:3)	LEAQ	64(SP), BP
// 0x001d 00029 (demo6_slice1.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x001d 00029 (demo6_slice1.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x001d 00029 (demo6_slice1.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//
// 没有第4行的事，第5行才为为sl分配内存空间
// 0x001d 00029 (demo6_slice1.go:5)	PCDATA	$2, $1
// 0x001d 00029 (demo6_slice1.go:5)	PCDATA	$0, $0
// 0x001d 00029 (demo6_slice1.go:5)	LEAQ	type.int(SB), AX
// 0x0024 00036 (demo6_slice1.go:5)	PCDATA	$2, $0
// 0x0024 00036 (demo6_slice1.go:5)	MOVQ	AX, (SP)
// 0x0028 00040 (demo6_slice1.go:5)	XORPS	X0, X0
// 0x002b 00043 (demo6_slice1.go:5)	MOVUPS	X0, 8(SP)
// 0x0030 00048 (demo6_slice1.go:5)	MOVQ	$0, 24(SP)
// 0x0039 00057 (demo6_slice1.go:5)	MOVQ	$1, 32(SP)
// 0x0042 00066 (demo6_slice1.go:5)	CALL	runtime.growslice(SB)
// 0x0047 00071 (demo6_slice1.go:5)	PCDATA	$2, $1
// 0x0047 00071 (demo6_slice1.go:5)	MOVQ	40(SP), AX
// 0x004c 00076 (demo6_slice1.go:5)	PCDATA	$2, $0
// 0x004c 00076 (demo6_slice1.go:5)	MOVQ	$1, (AX)
//
// 0x0053 00083 (demo6_slice1.go:6)	MOVQ	64(SP), BP
// 0x0058 00088 (demo6_slice1.go:6)	ADDQ	$72, SP
// 0x005c 00092 (demo6_slice1.go:6)	RET
// 0x005d 00093 (demo6_slice1.go:6)	NOP
// 0x005d 00093 (demo6_slice1.go:3)	PCDATA	$0, $-1
// 0x005d 00093 (demo6_slice1.go:3)	PCDATA	$2, $-1
// 0x005d 00093 (demo6_slice1.go:3)	CALL	runtime.morestack_noctxt(SB)
// 0x0062 00098 (demo6_slice1.go:3)	JMP	0

package main

func main() {
	sl := make([]int, 0)
	sl = append(sl, 1)
}

// ➜  asm git:(master) ✗ go tool compile -S demo6_slice2.go
// "".main STEXT size=101 args=0x0 locals=0x48
// 0x0000 00000 (demo6_slice2.go:3)	TEXT	"".main(SB), ABIInternal, $72-0
// 0x0000 00000 (demo6_slice2.go:3)	MOVQ	(TLS), CX
// 0x0009 00009 (demo6_slice2.go:3)	CMPQ	SP, 16(CX)
// 0x000d 00013 (demo6_slice2.go:3)	JLS	94
// 0x000f 00015 (demo6_slice2.go:3)	SUBQ	$72, SP
// 0x0013 00019 (demo6_slice2.go:3)	MOVQ	BP, 64(SP)
// 0x0018 00024 (demo6_slice2.go:3)	LEAQ	64(SP), BP
// 0x001d 00029 (demo6_slice2.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x001d 00029 (demo6_slice2.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x001d 00029 (demo6_slice2.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
//
// 0x001d 00029 (demo6_slice2.go:5)	PCDATA	$2, $1
// 0x001d 00029 (demo6_slice2.go:5)	PCDATA	$0, $0
// 0x001d 00029 (demo6_slice2.go:5)	LEAQ	type.int(SB), AX
// 0x0024 00036 (demo6_slice2.go:5)	PCDATA	$2, $0
// 0x0024 00036 (demo6_slice2.go:5)	MOVQ	AX, (SP)
// 0x0028 00040 (demo6_slice2.go:5)	PCDATA	$2, $1
// 0x0028 00040 (demo6_slice2.go:5)	LEAQ	""..autotmp_1+64(SP), AX
// 0x002d 00045 (demo6_slice2.go:5)	PCDATA	$2, $0
// 0x002d 00045 (demo6_slice2.go:5)	MOVQ	AX, 8(SP)
// 0x0032 00050 (demo6_slice2.go:5)	XORPS	X0, X0
// 0x0035 00053 (demo6_slice2.go:5)	MOVUPS	X0, 16(SP)
// 0x003a 00058 (demo6_slice2.go:5)	MOVQ	$1, 32(SP)
// 0x0043 00067 (demo6_slice2.go:5)	CALL	runtime.growslice(SB)
// 0x0048 00072 (demo6_slice2.go:5)	PCDATA	$2, $1
// 0x0048 00072 (demo6_slice2.go:5)	MOVQ	40(SP), AX
// 0x004d 00077 (demo6_slice2.go:5)	PCDATA	$2, $0
// 0x004d 00077 (demo6_slice2.go:5)	MOVQ	$1, (AX)
//
// 0x0054 00084 (demo6_slice2.go:6)	MOVQ	64(SP), BP
// 0x0059 00089 (demo6_slice2.go:6)	ADDQ	$72, SP
// 0x005d 00093 (demo6_slice2.go:6)	RET
// 0x005e 00094 (demo6_slice2.go:6)	NOP
// 0x005e 00094 (demo6_slice2.go:3)	PCDATA	$0, $-1
// 0x005e 00094 (demo6_slice2.go:3)	PCDATA	$2, $-1
// 0x005e 00094 (demo6_slice2.go:3)	CALL	runtime.morestack_noctxt(SB)
// 0x0063 00099 (demo6_slice2.go:3)	JMP	0

package main

func main() {
	sl := make([]int, 0, 1)
	sl = append(sl, 1)
}

// ➜  asm git:(master) ✗ go tool compile -S demo6_slice3.go
// "".main STEXT nosplit size=1 args=0x0 locals=0x0
// 0x0000 00000 (demo6_slice3.go:3)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $0-0
// 0x0000 00000 (demo6_slice3.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x0000 00000 (demo6_slice3.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x0000 00000 (demo6_slice3.go:3)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x0000 00000 (demo6_slice3.go:6)	RET
// 0x0000 c3

// 关闭优化
// ➜  asm git:(master) ✗ go tool compile -N -S demo6_slice3.go
// "".main STEXT nosplit size=98 args=0x0 locals=0x28
// 0x0000 00000 (demo6_slice3.go:3)	TEXT	"".main(SB), NOSPLIT|ABIInternal, $40-0
// 0x0000 00000 (demo6_slice3.go:3)	SUBQ	$40, SP
// 0x0004 00004 (demo6_slice3.go:3)	MOVQ	BP, 32(SP)
// 0x0009 00009 (demo6_slice3.go:3)	LEAQ	32(SP), BP
// 0x000e 00014 (demo6_slice3.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
// 0x000e 00014 (demo6_slice3.go:3)	FUNCDATA	$1, gclocals·54241e171da8af6ae173d69da0236748(SB)
// 0x000e 00014 (demo6_slice3.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
// 被优化
// 0x000e 00014 (demo6_slice3.go:4)	PCDATA	$2, $0
// 0x000e 00014 (demo6_slice3.go:4)	PCDATA	$0, $0
// 0x000e 00014 (demo6_slice3.go:4)	MOVQ	$0, ""..autotmp_1(SP)
// 0x0016 00022 (demo6_slice3.go:4)	PCDATA	$2, $1
// 0x0016 00022 (demo6_slice3.go:4)	LEAQ	""..autotmp_1(SP), AX
// 0x001a 00026 (demo6_slice3.go:4)	TESTB	AL, (AX)
// 0x001c 00028 (demo6_slice3.go:4)	JMP	30
// 0x001e 00030 (demo6_slice3.go:4)	PCDATA	$2, $-2
// 0x001e 00030 (demo6_slice3.go:4)	PCDATA	$0, $-2
// 0x001e 00030 (demo6_slice3.go:4)	JMP	32
// 0x0020 00032 (demo6_slice3.go:4)	PCDATA	$2, $1
// 0x0020 00032 (demo6_slice3.go:4)	PCDATA	$0, $0
// 0x0020 00032 (demo6_slice3.go:4)	MOVQ	AX, "".sl+8(SP)
// 0x0025 00037 (demo6_slice3.go:4)	MOVQ	$0, "".sl+16(SP)
// 0x002e 00046 (demo6_slice3.go:4)	MOVQ	$1, "".sl+24(SP)
// 0x0037 00055 (demo6_slice3.go:5)	JMP	57
// 0x0039 00057 (demo6_slice3.go:5)	MOVQ	$1, ""..autotmp_1(SP)
// 0x0041 00065 (demo6_slice3.go:5)	PCDATA	$2, $0
// 0x0041 00065 (demo6_slice3.go:5)	MOVQ	AX, "".sl+8(SP)
// 0x0046 00070 (demo6_slice3.go:5)	MOVQ	$1, "".sl+16(SP)
// 0x004f 00079 (demo6_slice3.go:5)	MOVQ	$1, "".sl+24(SP)
//
// 0x0058 00088 (demo6_slice3.go:6)	MOVQ	32(SP), BP
// 0x005d 00093 (demo6_slice3.go:6)	ADDQ	$40, SP
// 0x0061 00097 (demo6_slice3.go:6)	RET

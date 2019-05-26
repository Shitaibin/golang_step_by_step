package main

func main() {
	println("hello world")
}

// ➜  asm git:(master) ✗ go tool compile -S demo1.go
// "".main STEXT size=81 args=0x0 locals=0x18
//         0x0000 00000 (demo1.go:3)       TEXT    "".main(SB), ABIInternal, $24-0
// 		   前导：栈分裂。TLS是当前g的指针，当前g的栈空间不够用时，跳到74行执行栈分裂
//         0x0000 00000 (demo1.go:3)       MOVQ    (TLS), CX
//         0x0009 00009 (demo1.go:3)       CMPQ    SP, 16(CX)
//         0x000d 00013 (demo1.go:3)       JLS     74
// 		   调用下一个函数前，当前函数的栈向下增长
//         0x000f 00015 (demo1.go:3)       SUBQ    $24, SP
//         0x0013 00019 (demo1.go:3)       MOVQ    BP, 16(SP)
//         0x0018 00024 (demo1.go:3)       LEAQ    16(SP), BP
//          GC相关
//         0x001d 00029 (demo1.go:3)       FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//         0x001d 00029 (demo1.go:3)       FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//         0x001d 00029 (demo1.go:3)       FUNCDATA        $3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
// 		   调用println，设置被掉函数参数，参数入栈
//         0x001d 00029 (demo1.go:4)       PCDATA  $2, $0
//         0x001d 00029 (demo1.go:4)       PCDATA  $0, $0
//         0x001d 00029 (demo1.go:4)       CALL    runtime.printlock(SB)
//         0x0022 00034 (demo1.go:4)       PCDATA  $2, $1
//         0x0022 00034 (demo1.go:4)       LEAQ    go.string."hello world\n"(SB), AX
//         0x0029 00041 (demo1.go:4)       PCDATA  $2, $0
//         0x0029 00041 (demo1.go:4)       MOVQ    AX, (SP)
//         0x002d 00045 (demo1.go:4)       MOVQ    $12, 8(SP)
//         0x0036 00054 (demo1.go:4)       CALL    runtime.printstring(SB)
//         0x003b 00059 (demo1.go:4)       CALL    runtime.printunlock(SB)
//  	   被掉函数执行完，main函数收缩栈
//         0x0040 00064 (demo1.go:5)       MOVQ    16(SP), BP
//         0x0045 00069 (demo1.go:5)       ADDQ    $24, SP
//         0x0049 00073 (demo1.go:5)       RET
// 		   没错，这就是74行，的栈分裂
//         0x004a 00074 (demo1.go:5)       NOP
//         0x004a 00074 (demo1.go:3)       PCDATA  $0, $-1
//         0x004a 00074 (demo1.go:3)       PCDATA  $2, $-1
//         0x004a 00074 (demo1.go:3)       CALL    runtime.morestack_noctxt(SB) ;栈分裂调用morestack_noctxt
//         0x004f 00079 (demo1.go:3)       JMP     0 // 分裂完成后，从0开始执行

// ➜  asm git:(master) ✗ ls
// demo1.go demo1.o  demo2.go demo3.go
// ➜  asm git:(master) ✗ go tool objdump demo1.o
// TEXT %22%22.main(SB) gofile../Users/shitaibin/Workspace/golang_step_by_step/asm/demo1.go
//   栈分裂判断
//   demo1.go:3            0x35d                   65488b0c2500000000      MOVQ GS:0, CX           [5:9]R_TLS_LE
//   demo1.go:3            0x366                   483b6110                CMPQ 0x10(CX), SP
//   demo1.go:3            0x36a                   763b                    JBE 0x3a7
// 	 main函数栈向下增长
//   demo1.go:3            0x36c                   4883ec18                SUBQ $0x18, SP
//   demo1.go:3            0x370                   48896c2410              MOVQ BP, 0x10(SP)
//   demo1.go:3            0x375                   488d6c2410              LEAQ 0x10(SP), BP
//   demo1.go:4            0x37a                   e800000000              CALL 0x37f              [1:5]R_CALL:runtime.printlock
//   demo1.go:4            0x37f                   488d0500000000          LEAQ 0(IP), AX          [3:7]R_PCREL:go.string."hello world\n"
//   demo1.go:4            0x386                   48890424                MOVQ AX, 0(SP)
//   demo1.go:4            0x38a                   48c74424080c000000      MOVQ $0xc, 0x8(SP)
//   demo1.go:4            0x393                   e800000000              CALL 0x398              [1:5]R_CALL:runtime.printstring
//   demo1.go:4            0x398                   e800000000              CALL 0x39d              [1:5]R_CALL:runtime.printunlock
// 	 栈收缩
//   demo1.go:5            0x39d                   488b6c2410              MOVQ 0x10(SP), BP
//   demo1.go:5            0x3a2                   4883c418                ADDQ $0x18, SP
//   demo1.go:5            0x3a6                   c3                      RET
// 	 栈扩容
//   demo1.go:3            0x3a7                   e800000000              CALL 0x3ac              [1:5]R_CALL:runtime.morestack_noctxt
//   demo1.go:3            0x3ac                   ebaf                    JMP %22%22.main(SB)

// 上面的方式有个22%感觉别扭啊，请教了一下曹大，dump可执行程序就会显示包名了。
// ➜  asm git:(master) ✗ go tool objdump demo1 |  grep  -A 20 "demo1"
// TEXT main.main(SB) /Users/shitaibin/Workspace/golang_step_by_step/asm/demo1.go
// demo1.go:3		0x104ea70		65488b0c2530000000	MOVQ GS:0x30, CX
// demo1.go:3		0x104ea79		483b6110		CMPQ 0x10(CX), SP
// demo1.go:3		0x104ea7d		763b			JBE 0x104eaba
// demo1.go:3		0x104ea7f		4883ec18		SUBQ $0x18, SP
// demo1.go:3		0x104ea83		48896c2410		MOVQ BP, 0x10(SP)
// demo1.go:3		0x104ea88		488d6c2410		LEAQ 0x10(SP), BP
// demo1.go:4		0x104ea8d		e85e4efdff		CALL runtime.printlock(SB)
// demo1.go:4		0x104ea92		488d050fee0100		LEAQ go.string.*+2376(SB), AX
// demo1.go:4		0x104ea99		48890424		MOVQ AX, 0(SP)
// demo1.go:4		0x104ea9d		48c74424080c000000	MOVQ $0xc, 0x8(SP)
// demo1.go:4		0x104eaa6		e87557fdff		CALL runtime.printstring(SB)
// demo1.go:4		0x104eaab		e8c04efdff		CALL runtime.printunlock(SB)
// demo1.go:5		0x104eab0		488b6c2410		MOVQ 0x10(SP), BP
// demo1.go:5		0x104eab5		4883c418		ADDQ $0x18, SP
// demo1.go:5		0x104eab9		c3			RET
// demo1.go:3		0x104eaba		e8f184ffff		CALL runtime.morestack_noctxt(SB)
// demo1.go:3		0x104eabf		ebaf			JMP main.main(SB)

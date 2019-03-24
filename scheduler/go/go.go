package main

func main() {
	go func() {}()
}

// Output
// go build -o go go.go
// ➜  go git:(master) ✗ go tool objdump -s 'main\.main' go
// TEXT main.main(SB) /Users/shitaibin/Workspace/golang_step_by_step/scheduler/go/go.go
//   go.go:3		0x104e020		65488b0c2530000000	MOVQ GS:0x30, CX
//   go.go:3		0x104e029		483b6110		CMPQ 0x10(CX), SP
//   go.go:3		0x104e02d		7630			JBE 0x104e05f
//   go.go:3		0x104e02f		4883ec18		SUBQ $0x18, SP
//   go.go:3		0x104e033		48896c2410		MOVQ BP, 0x10(SP)
//   go.go:3		0x104e038		488d6c2410		LEAQ 0x10(SP), BP
//   go.go:4		0x104e03d		c7042400000000	MOVL $0x0, 0(SP)
//   go.go:4		0x104e044		488d05d50c0200	LEAQ go.func.*+68(SB), AX
//   go.go:4		0x104e04b		4889442408		MOVQ AX, 0x8(SP)
//   go.go:4		0x104e050		e89bcefdff		CALL runtime.newproc(SB) // go
//   go.go:5		0x104e055		488b6c2410		MOVQ 0x10(SP), BP
//   go.go:5		0x104e05a		4883c418		ADDQ $0x18, SP
//   go.go:5		0x104e05e		c3				RET
//   go.go:3		0x104e05f		e85c89ffff		CALL runtime.morestack_noctxt(SB)
//   go.go:3		0x104e064		ebba			JMP main.main(SB)

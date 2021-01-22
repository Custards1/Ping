package bytecode

const (
	Add  = byte(iota)
	Sub  = byte(iota)
	Mul  = byte(iota)
	Div  = byte(iota)
	Push = byte(iota)
	Pop  = byte(iota)
	Cmp  = byte(iota)
	Je   = byte(iota)
	Jne  = byte(iota)
	Jmp  = byte(iota)
	Jl   = byte(iota)
	Jle  = byte(iota)
	Jg   = byte(iota)
	Jre  = byte(iota)
	Call = byte(iota)
)

type Instruction struct {
	inst byte
	arg  interface{}
}

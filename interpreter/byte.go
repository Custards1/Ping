package interpreter

type OpCode byte

const (
	OpAdd  = OpCode(iota)
	OpSub  = OpCode(iota)
	OpMul  = OpCode(iota)
	OpDiv  = OpCode(iota)
	OpPush = OpCode(iota)
	OpPop  = OpCode(iota)
	OpCmp  = OpCode(iota)
	OpJe   = OpCode(iota)
	OpJne  = OpCode(iota)
	OpJmp  = OpCode(iota)
	OpJl   = OpCode(iota)
	OpJle  = OpCode(iota)
	OpJg   = OpCode(iota)
	OpJre  = OpCode(iota)
	OpCall = OpCode(iota)
)

type Instruction struct {
	inst OpCode
	arg  interface{}
}

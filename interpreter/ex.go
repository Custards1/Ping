package interpreter

type Executor struct {
}

func (e *Executor) Execute(ops *[]Instruction) error {
	max := len(*ops)
	b := *ops
	for i := 0; i < max; i++ {
		switch b[i].inst {
		case OpAdd:
			break

		}
	}
	return nil
}

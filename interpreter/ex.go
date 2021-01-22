package interpreter

type Executor struct {
	stack *Stack
	store *ObjectStore
}

func (e *Executor) Execute(ops *[]Instruction) error {
	*e.stack.max = uintptr(len(*ops))
	b := ops
	for {
		i, ok := e.stack.Inc()
		if !ok {
			break
		}
		a := *b
		switch a[i].inst {
		case OpAdd:
			if err := e.stack.Add(); err != nil {
				return err
			}
			break
		case OpSub:
			if err := e.stack.Sub(); err != nil {
				return err
			}
			break
		case OpDiv:
			if err := e.stack.Div(); err != nil {
				return err
			}
			break
		case OpMul:
			if err := e.stack.Mul(); err != nil {
				return err
			}
			break
		case OpStore:
			if err := e.stack.Store(e.store); err != nil {
				return err
			}
			break
		case OpCall:
			switch a[i].arg
			if err := e.stack.Call(e.store); err != nil {
				return err
			}
			break
		default:
			break
		}

	}
	return nil
}

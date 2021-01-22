package interpreter

import (
	"errors"
	"math/big"
	"sync"
)

var (
	ErrorStackEmpty = errors.New("Stack Empty")
	ErrorInvalidOp  = errors.New("Invalid Operation")
)

type CmpResult int

const (
	CmpLe = CmpResult(-1)
	CmpEq = CmpResult(0)
	CmpGr = CmpResult(1)
)

type Stack struct {
	items  []*PingObj
	lock   sync.RWMutex
	caller *Stack
	cmp    CmpResult
	sp     *uintptr
	max    *uintptr
	args   *[]*PingObj
}

// New creates a new ItemStack
func NewStack(sp *uintptr, args *[]*PingObj) *Stack {
	s := Stack{sp: sp, args: args}
	s.items = make([]*PingObj, 0, 16)

	return &s
}
func NewSubStack(caller *Stack, sp *uintptr, args *[]*PingObj) *Stack {
	s := Stack{sp: sp, args: args}
	s.items = make([]*PingObj, 0, 16)
	s.caller = caller
	return &s
}

func (s *Stack) Push(t *PingObj) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}
func (s *Stack) Inc() (uintptr, bool) {
	a := *s.sp
	*s.sp += 1
	ok := true
	if a >= *s.max {
		switch s.caller {
		case nil:
			ok = false
			break
		default:
			s = s.caller
			break
		}
		*s.sp = 1
		return 0, ok
	}
	return a, ok
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() (*PingObj, error) {
	s.lock.Lock()
	item, err := s.pop()
	s.lock.Unlock()
	return item, err
}
func (s *Stack) pop() (*PingObj, error) {
	a := len(s.items)
	if a == 0 {
		return nil, ErrorStackEmpty
	}
	item := s.items[a-1]
	s.items = s.items[0 : a-1]
	return item, nil
}

type IntNumberOperation func(eax *big.Int, ebx *big.Int) error
type FloatNumberOperation func(eax *big.Float, ebx *big.Float) error

func (s *Stack) NumOp(op IntNumberOperation, fop FloatNumberOperation) error {
	s.lock.Lock()
	a, e := s.pop()
	if e != nil {
		s.lock.Unlock()
		return e
	}
	b, v := s.pop()
	if v != nil {
		s.lock.Unlock()
		return v
	}
	s.lock.Unlock()

	if a.otype != TypeNumber || b.otype != TypeNumber {
		return ErrorInvalidOp
	}
	ac := a.obj
	bc := b.obj
	switch ac.(type) {
	case *big.Int:
		switch bc.(type) {
		case *big.Int:
			res := ac.(*big.Int)
			c := bc.(*big.Int)
			op(res, c)
			a.SetNumber(res)
			break
		case *big.Float:
			resa := ac.(*big.Int)
			res := new(big.Float).SetInt(resa)
			c := bc.(*big.Float)
			fop(res, c)
			a.SetNumber(res)
		default:
			return ErrorInvalidOp

		}
		break
	case *big.Float:
		switch bc.(type) {
		case *big.Int:

			resa := bc.(*big.Int)
			res := new(big.Float).SetInt(resa)
			c := ac.(*big.Float)
			fop(c, res)
			a.SetNumber(c)
			break
		case *big.Float:
			res := ac.(*big.Float)
			c := bc.(*big.Float)
			fop(res, c)
			a.SetNumber(res)
		default:
			return ErrorInvalidOp

		}
		break
	default:
		return ErrorInvalidOp

	}
	return nil
}
func (s *Stack) Add() error {
	return s.NumOp(func(eax *big.Int, ebx *big.Int) error {
		eax = eax.Add(eax, ebx)
		return nil
	}, func(eax *big.Float, ebx *big.Float) error {
		eax = eax.Add(eax, ebx)
		return nil
	})
}
func (s *Stack) Sub() error {
	return s.NumOp(func(eax *big.Int, ebx *big.Int) error {
		eax = eax.Sub(eax, ebx)
		return nil
	}, func(eax *big.Float, ebx *big.Float) error {
		eax = eax.Sub(eax, ebx)
		return nil
	})
}
func (s *Stack) Mul() error {
	return s.NumOp(func(eax *big.Int, ebx *big.Int) error {
		eax = eax.Sub(eax, ebx)
		return nil
	}, func(eax *big.Float, ebx *big.Float) error {
		eax = eax.Sub(eax, ebx)
		return nil
	})
}
func (s *Stack) Div() error {
	return s.NumOp(func(eax *big.Int, ebx *big.Int) error {
		eax = eax.Div(eax, ebx)
		return nil
	}, func(eax *big.Float, ebx *big.Float) error {
		eax = eax.Quo(eax, ebx)
		return nil
	})
}
func (s *Stack) Store(o *ObjectStore) error {
	s.lock.Lock()
	a, e := s.pop()
	if e != nil {
		s.lock.Unlock()
		return e
	}
	b, v := s.pop()
	if v != nil {
		s.lock.Unlock()
		return v
	}
	s.lock.Unlock()
	if a == nil || b == nil {
		return ErrorStackEmpty
	}
	b.Dec(o)
	a.hashid = b.hashid
	o.Add(a)
	return nil
}
func (s *Stack) Call(amount uint) error {
	sp := uintptr(0)
	//stack := NewSubStack(s, &sp, nil)

	s.lock.Lock()
	var l *[]*PingObj
	var err error
	switch amount {
	case 0, 1:
		break
	default:
		baf := int(amount) - 1
		*l = make([]*PingObj, 0, baf)
		baf--

		for u := baf; u >= 0; u-- {
			(*l)[u], err = s.pop()
			if err != nil {
				return err
			}
		}
		break
	}
	fun, err := s.pop()
	s.lock.Unlock()
	if err != nil {
		return err
	}
	if fun == nil || fun.obj == nil || fun.otype != TypeFunctionNative {
		return ErrorInvalidOp
	}
	switch fun.obj.(type) {
	case *Instruction:
		v := fun.obj.(*Instruction)
		s = NewSubStack(s, &sp, l)
		return nil
	default:
		break
	}

	return ErrorInvalidOp
}

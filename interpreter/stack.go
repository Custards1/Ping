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

type Stack struct {
	items  []Pong
	lock   sync.RWMutex
	caller *Stack
}

// New creates a new ItemStack
func NewStack() *Stack {
	s := Stack{}
	s.items = make([]Pong, 0, 16)
	return &s
}
func NewSubStack(caller *Stack) *Stack {
	s := Stack{}
	s.items = make([]Pong, 0, 16)
	s.caller = caller
	return &s
}

func (s *Stack) Push(t Pong) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// Pop removes an Item from the top of the stack
func (s *Stack) Pop() Pong {
	s.lock.Lock()
	item := s.pop()
	s.lock.Unlock()
	return item
}
func (s *Stack) pop() Pong {
	a := len(s.items)
	if a == 0 {
		return nil
	}
	item := s.items[a-1]
	s.items = s.items[0 : a-1]
	return item
}
func (s *Stack) Add() error {
	s.lock.Lock()
	a := s.pop()
	b := s.pop()
	s.lock.Unlock()
	if a == nil || b == nil {
		return ErrorStackEmpty
	}
	if a.otype != TypeNumber || b.otype != TypeNumber {
		return ErrorInvalidOp
	}
	ac := *a.obj
	bc := *b.obj
	var result *PingObj
	switch ac.(type) {
	case big.Int:
		switch bc.(type) {
		case big.Int:
			res := ac.(big.Int)
			c := bc.(big.Int)
			b := res.Add(&res, &c)
			a.SetNumber()
			a.obj = PingPtr(b)
			break
		default:
			break
		}
	default:
		break
	}

	return nil
}

package main

import "fmt"

type stack []interface{}

func (s stack) Push(v interface{}) (stack) {
	return append(s,v)
}

func (s stack) Pop() (stack, interface{}) {
	length := len(s);
	if length == 0 {
		return s, nil
	}
	return s[:length - 1], s[length - 1]
}

func countUp(stackVar stack, op Operation) (stack, error) {
	stackVar, j := stackVar.Pop()
	stackVar, k := stackVar.Pop()
	if j == nil || k == nil {
		return nil, fmt.Errorf("Don't have vars to do op %v", op)
	}
	newVar, err := op.Do(k, j)
	if err != nil {
		return nil, fmt.Errorf("Can't do op %v: %v", op, err)
	}
	stackVar = stackVar.Push(newVar)
	return stackVar, nil
}

func culc(tokens []interface{}) (interface{}, error) {
	var (
		err			error
		stackOp		stack
		stackVar	stack
		op			interface{}
	)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i].(type) {
		case Operation:
			stackOp, op  = stackOp.Pop()
			for op != nil && tokens[i].(Operation).CheckW(op.(Operation)) {
				stackVar, err = countUp(stackVar, op.(Operation))
				if err != nil {
					return nil, fmt.Errorf("Can't Countup() op %v: %v", op, err)
				}
				stackOp, op  = stackOp.Pop()
			}
			if op != nil {
				stackOp = stackOp.Push(op)
			}
			stackOp = stackOp.Push(tokens[i])
		default:
			stackVar = stackVar.Push(tokens[i])
		}
	}

	stackOp, op  = stackOp.Pop()
	if op != nil {
		stackVar, err = countUp(stackVar, op.(Operation))
		if err != nil {
			return nil, fmt.Errorf("Can't Countup() op %v: %v", op, err)
		}
		stackOp, op  = stackOp.Pop()

	}
	stackVar, newVar := stackVar.Pop()
	stackVar, secVar := stackVar.Pop()
	if secVar != nil {
		stackVar = stackVar.Push(secVar)
		return nil, fmt.Errorf("In stackVar left vars %v", stackVar)
	}
	return newVar, nil
}

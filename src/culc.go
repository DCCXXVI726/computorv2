package main

import (
	"fmt"
)

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

func culc(tokens []interface{}) (interface{}, error, bool, int) {
	var (
		err			error
		stackOp		stack
		stackVar	stack
		op			interface{}
		i			int
	)
	flag := false
	lastToken := "first"
	Loop:
	for i = 0; i < len(tokens); i++ {
		switch tokens[i].(type) {
		case Operation:
			switch lastToken {
			case "first":
				if tokens[i].(Operation).str == "-" {
					stackVar = stackVar.Push(0.0)
				}
			case "Operation":
				return nil, fmt.Errorf("there is 2 operation side by side in %v", tokens), flag, i
			}
			stackOp, op  = stackOp.Pop()
			for op != nil && tokens[i].(Operation).CheckW(op.(Operation)) {
				stackVar, err = countUp(stackVar, op.(Operation))
				if err != nil {
					return nil, fmt.Errorf("Can't Countup() op %v: %v", op, err), flag, i
				}
				stackOp, op  = stackOp.Pop()
			}
			if op != nil {
				stackOp = stackOp.Push(op)
			}
			stackOp = stackOp.Push(tokens[i])
			lastToken = "Operation"
		case Braket:
			if tokens[i].(Braket).isOpen() {
				v, err, flag1, k := culc(tokens[i+1:])
				if err != nil || !flag1{
					return nil, fmt.Errorf("Problem with open bracket %v", tokens), flag, i
				}
				stackVar = stackVar.Push(v)
				lastToken = "Var"
				i = i + k + 1
			} else if tokens[i].(Braket).isClose() {
				flag = true
				break Loop
			}
		default:
			if lastToken == "Var" {
				return nil, fmt.Errorf("there is 2 var side by side in %v", tokens), flag, i
			}
			stackVar = stackVar.Push(tokens[i])
			lastToken = "Var"
		}
	}

	stackOp, op  = stackOp.Pop()
	for op != nil {
		stackVar, err = countUp(stackVar, op.(Operation))
		if err != nil {
			return nil, fmt.Errorf("Can't Countup() op %v: %v", op, err), flag, i
		}
		stackOp, op  = stackOp.Pop()
	}
	stackVar, newVar := stackVar.Pop()
	stackVar, secVar := stackVar.Pop()
	if secVar != nil {
		stackVar = stackVar.Push(secVar)
		return nil, fmt.Errorf("In stackVar left vars %v", stackVar), flag, i
	}
	return newVar, nil, flag, i
}

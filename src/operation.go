package main

import (
	"fmt"
	"math"
)
//Operation is struct for operation
type Operation struct {
	str		string
	weight	int
}

func (op Operation) CheckW(op2 Operation) bool {
	if op.weight <= op2.weight {
		return true
	}
	return false
}
func (op Operation) Do(i interface{}, k interface{}) (interface{}, error) {
	switch op.str {
	case "+":
		return Sum(i, k)
	case "-":
		return Sub(i, k)
	case "*":
		return Multy(i, k)
	case "/":
		return Dev(i, k)
	case "%":
		return Mod(i, k)

	}
	return nil, fmt.Errorf("can't find operation %v", op.str)
}

func Mod(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return math.Mod(i.(float64),  k.(float64)), nil

		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}
func GetOperation(str string) (interface{}, int, error) {
	switch str[0] {
	case '+':
		return Operation{"+", 1}, 0, nil
	case '-':
		return Operation{"-", 1}, 0, nil
	case '*':
		return Operation{"*", 2}, 0, nil
	case '/':
		return Operation{"/", 2}, 0, nil
	case '%':
		return Operation{"%", 2}, 0, nil
	}
	return nil, 0, fmt.Errorf("Can't get operation from %v", str)
}

func Multy(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return i.(float64) * k.(float64), nil
		case Complex:
			return Complex{i.(float64) * k.(Complex).Re, i.(float64) * k.(Complex).Im}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	case Complex:
		switch stype := k.(type) {
		case float64:
			return Complex{k.(float64) * i.(Complex).Re, k.(float64) * i.(Complex).Im}, nil
		case Complex:
			a := i.(Complex)
			b := k.(Complex)
			return Complex{a.Re*b.Re-a.Im*b.Im, a.Re * b.Im + a.Im * b.Re}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

func Dev(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return i.(float64) / k.(float64), nil
		case Complex:
			a := i.(float64)
			b := k.(Complex)
			return Complex{a * b.Re/(b.Re * b.Re + b.Im * b.Im), -a * b.Im/(b.Re * b.Re + b.Im * b.Im)},nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	case Complex:
		switch stype := k.(type) {
		case float64:
			a := i.(Complex)
			b := k.(float64)
			return Complex{a.Re/b, a.Im/b}, nil
		case Complex:
			a := i.(Complex)
			b := k.(Complex)
			return Complex{(a.Re * b.Re + a.Im * b.Im) / (b.Re * b.Re + b.Im * b.Im), (b.Re * a.Im - a.Re * b.Im) / (b.Re * b.Re + b.Im * b.Im)}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	default:
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

func Sub(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return i.(float64) - k.(float64), nil
		case Complex:
			return Complex{i.(float64) - k.(Complex).Re, -k.(Complex).Im}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	case Complex:
		switch stype := k.(type) {
		case float64:
			return Complex{ i.(Complex).Re - k.(float64), i.(Complex).Im}, nil
		case Complex:
			return Complex{i.(Complex).Re - k.(Complex).Re, i.(Complex).Im - k.(Complex).Im}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	default:
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

//Sum two var
func Sum(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return i.(float64) + k.(float64), nil
		case Complex:
			return Complex{i.(float64) + k.(Complex).Re, k.(Complex).Im}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	case Complex:
		switch stype := k.(type) {
		case float64:
			return Complex{k.(float64) + i.(Complex).Re, i.(Complex).Im}, nil
		case Complex:
			return Complex{k.(Complex).Re + i.(Complex).Re, k.(Complex).Im + i.(Complex).Im}, nil
		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	default:
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

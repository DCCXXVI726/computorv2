package main

import "fmt"
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
	}
	return nil, fmt.Errorf("can't find operation %v", op.str)
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
	}
	return nil, 0, fmt.Errorf("Can't get operation from %v", str)
}

func Multy(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
		case float64:
			return i.(float64) * k.(float64), nil

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

		default:
			return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
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
			default:
				return nil, fmt.Errorf("Can't find type for second var: %v", stype)
		}
	default:
		return nil, fmt.Errorf("Can't find type for first var: %v", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

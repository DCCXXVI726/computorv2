package main

import "fmt"

type Operation struct {
	str		string
	weight	int
}

func Sum(i interface{}, k interface{}) (interface{}, error) {
	switch ftype := i.(type) {
	case float64:
		switch stype := k.(type) {
			case float64:
				return i.(float64) + k.(float64), nil
			default:
				return nil, fmt.Errorf("Can't find type for second var: %T", stype)
		}
	default:
		return nil, fmt.Errorf("Can't find type for first var: %T", ftype)
	}
	return nil, fmt.Errorf("End of method")
}

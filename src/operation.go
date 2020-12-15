package main

// Operation is type for "+" "-" and etc
type Operation struct {
	text string
}

var myMap map[string]OperFunc = map[string]OperFunc{
	"+": Add,
	"*": Mult,
	"-": Minus,
	"/": Div,
}

var priorityMap map[string]int = map[string]int{
	"+": 1,
	"*": 2,
	"-": 1,
	"/": 2,
}

// OperFunc for reduce some code
type OperFunc func(interface{}, interface{}) (interface{}, error)

// return func
func (op Operation) getFunc() OperFunc {
	oper, _ := myMap[op.text]
	return oper
}

//IsOper check string for operation
func IsOper(s string) bool {
	_, ok := myMap[s]
	return ok
}

// Priority check priority
func Priority(f Operation, s Operation) bool {
	firstOper, _ := priorityMap[f.text]
	secondOper, _ := priorityMap[s.text]
	if firstOper <= secondOper {
		return true
	}
	return false
}

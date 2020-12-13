package main

type Operation struct {
	text string
}

type OperFunc func(interface{}, interface{}) (interface{}, error)

func (op Operation) getFunc() OperFunc {
	/*myMap := map[string]operFunc{
		"+": Add,
	}*/
	if op.text == "+" {
		return Add
	}
	return nil
}

func Priority(f Operation, s Operation) bool {
	return true
}

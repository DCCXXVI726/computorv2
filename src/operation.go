package main

type Operation struct {
	text string
}

type operFunc func(interface{}, interface{}) (interface{}, error)

func (op Operation) getFunc() operFunc {
	/*myMap := map[string]operFunc{
		"+": Add,
	}*/
	if op.text == "+" {
		return Add
	}
	return nil
}

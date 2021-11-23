package main

type Complex struct {
	Re float64
	Im float64
}
func checkIm(i interface{}) interface{} {
	switch i.(type) {
	case Complex:
		if i.(Complex).Im == 0 {
			return i.(Complex).Re
		}
	}
	return i
}

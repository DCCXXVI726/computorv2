package main

// Mult multiplies two types
func Mult(fir interface{},
	sec interface{}) (
	interface{},
	error) {
	switch fir.(type) {
	case Complex:
		first := fir.(Complex)
		switch sec.(type) {
		case Complex:
			second := sec.(Complex)
			re := first.Re*second.Re - first.Im*second.Im
			im := first.Re*second.Im + first.Im*second.Re
			return Complex{re, im}, nil
		}
	}
	return nil, nil
}

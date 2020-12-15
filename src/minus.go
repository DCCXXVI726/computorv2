package main

//Minus minus two types
func Minus(fir interface{}, sec interface{}) (interface{}, error) {
	switch fir.(type) {
	case Complex:
		first := fir.(Complex)
		switch sec.(type) {
		case Complex:
			second := sec.(Complex)
			return Complex{first.Re - second.Re, first.Im - second.Im}, nil
		}
	}
	return nil, nil
}
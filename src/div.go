package main

import (
	"fmt"
)

// Div divides two types
func Div(fir interface{},
	sec interface{}) (
	interface{},
	error) {
	switch fir.(type) {
	case Complex:
		first := fir.(Complex)
		switch sec.(type) {
		case Complex:
			second := sec.(Complex)
			abs := second.Re*second.Re + second.Im*second.Im
			if abs == 0 {
				err := fmt.Errorf("Can't div by zero")
				return nil, err
			}
			re := (first.Re*second.Re + first.Im*second.Im) / abs
			im := (-first.Re*second.Im + first.Im*second.Re) / abs
			return Complex{re, im}, nil
		}
	}
	return nil, nil
}

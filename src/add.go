package main

import ()

func Add(fir interface{}, sec interface{}) (interface{}, error) {
	switch fir.(type) {
	case Complex:
		switch sec.(type) {
		case Complex:
			first := fir.(Complex)
			second := sec.(Complex)
			return Complex{first.Re + second.Re, first.Im + second.Im}, nil
		}
	}
	return nil, nil
}

package main

import ()

func Add(fir interface{}, sec interface{}) (interface{}, error) {
	switch fir.(type) {
	case Complex:
		switch sec.(type) {
		case Complex:
			return Complex{0, 0}, nil
		}
	}
	return nil, nil
}

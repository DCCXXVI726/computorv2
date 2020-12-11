package main

import (
	"fmt"
)

func Add(fir interface{}, sec interface{}) {
	switch fir.(type) {
	case Complex:
		switch sec.(type) {
		case Complex:
			fmt.Print("два комплексных")
		}
	}
}

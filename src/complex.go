package main

import (
	"fmt"
	"strconv"
)

//Complex is complex number
type Complex struct {
	Re float64
	Im float64
}

// to string
func (number Complex) toString() string {
	if number.Im == 0 {
		re := strconv.FormatFloat(number.Re, 'f', -1, 64)
		return re
	}
	if number.Re == 0 {
		im := strconv.FormatFloat(number.Im, 'f', -1, 64)
		return im + "i"
	}
	re := strconv.FormatFloat(number.Re, 'f', -1, 64)
	im := strconv.FormatFloat(number.Im, 'f', -1, 64)
	return fmt.Sprintf("%s + %si", re, im)
}

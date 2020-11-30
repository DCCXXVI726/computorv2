package main

type Complex struct {
	Re float64
	Im float64
}

func Add(a Complex, b Complex) Complex {
	c := Complex{}
	c.Re = a.Re + b.Re
	c.Im = a.Im + b.Im
	return c
}

func Multy(a Complex, b Complex) Complex {
	c := Complex{}
	c.Re = a.Re*b.Re - a.Im*b.Im
	c.Im = a.Re*b.Im + b.Re*a.Im
	return c
}

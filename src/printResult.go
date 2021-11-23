package main

import "fmt"

func printResult(i interface{}) {
	switch i.(type) {
	case float64:
		fmt.Print(i.(float64))
	case Complex:
		fmt.Printf("%f + %fi", i.(Complex).Re, i.(Complex).Im)
	case Matrix:
		a := i.(Matrix)
		for i := 0; i < a.H; i++ {
			fmt.Print("[")
			for k := 0; k < a.W; k++ {
				printResult(a.Value[i][k])
				if k != a.W - 1 {
					fmt.Print(", ")
				} else if i != a.H - 1 {
					fmt.Println("]")
				} else {
					fmt.Print("]")
				}
			}
		}
	}
}

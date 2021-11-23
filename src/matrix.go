package main

import "fmt"

type Matrix struct {
	H, W int
	Value [][]interface{}
}

func SumMatrix(a Matrix, b Matrix) (interface{}, error) {
	if a.W != b.W || a.H != b.H {
		return nil, fmt.Errorf("matrix not equal by width, height to sum")
	}
	c := Matrix{}
	c.W = a.W
	c.H = a.H
	c.Value = make([][]interface{}, c.H)
	for i := 0; i < c.H; i++ {
		row := make([]interface{}, c.W)
		for k := 0; k < c.W; k++ {
			tmp, err := Sum(a.Value[i][k], b.Value[i][k])
			if err != nil {
				return nil, fmt.Errorf("can't multy elements from first matrix: %v and second matrix: %v with h - %d, w - %d,: %s",
					a.Value[i][k],b.Value[i][k],i,k, err)
			}
			row[k] = tmp
		}
		c.Value[i] = row
	}
	return c, nil
}

func MultyMatrix(i interface{}, k interface{}) (interface{}, error) {
	var a, b Matrix
	switch i.(type) {
	case Matrix:
		a = i.(Matrix)
	default:
		return nil, fmt.Errorf("first multyplier is not matrix: %v", i)
	}
	switch k.(type) {
	case Matrix:
		b = k.(Matrix)
	default:
		return nil, fmt.Errorf("second multyplier is not matrix: %v", k)
	}
	if a.W != b.H {
		return nil, fmt.Errorf("can't multyply matrix: width of first matrix doesn't equal height of second matrix")
	}
	c := Matrix{}
	c.W = b.W
	c.H = a.H
	c.Value = make([][]interface{}, c.H)
	for i := 0; i < c.H ; i++ {
		row := make([]interface{}, c.W)
		for m := 0; m < c.W; m++ {
			row[m] = 0.0
		}
		for k := 0; k < c.W; k++ {
			for j := 0; j < a.W; j++ {
				tmp, err := Multy(a.Value[i][j],  b.Value[j][k])
				if err != nil {
					return nil, fmt.Errorf("can't multy elements from first matrix: %v with w - %d, h - %d, " +
						"and second matrix: %v with w - %d, h -%d: %s", a.Value[k][j],k,j,b.Value[j][k],j,k, err)
				}
				tmp2, err := Sum(row[k], tmp)
				if err != nil {
					return nil, fmt.Errorf("can't sum elements %v and %v: %s", row[k], tmp, err)
				}
				row[k] = tmp2
			}
		}
		c.Value[i] = row
	}
	return c, nil
}

func getRow(str string) ([]interface{}, int, error) {
	i := 0
	row := make([]interface{}, 0)
	for {
		for i < len(str) && str[i] == ' ' {
			i++
		}
		k := i
		for i < len(str) && str[i] != ']' && str[i] != ',' {
			i++
		}
		if i == len(str) {
			return nil, 0, fmt.Errorf("can't create maxrix: end of string")
		}
		tokens, err := createTokens(str[k:i])
		if err != nil {
			return nil, 0, fmt.Errorf("problem with create token: %s", err)
		}
		token, err, _, _ := culc(tokens)
		if err != nil {
			return nil, 0, fmt.Errorf("problem with culc: %s", err)
		}
		row = append(row, token)
		if str[i] == ']' {
			return row, i, err
		}
		i++
	}
}

func GetMatrix (str string) (interface{}, int, error) {
	i := 1
	matrix := Matrix{}
	matrix.Value = make([][]interface{},0)
	for {
		for i < len(str) && str[i] == ' ' {
			i++
		}
		if i == len(str) || str[i] != '[' {
			return nil, 0, fmt.Errorf("don't have second [ in matrix")
		}
		i++
		if i >= len(str) {
			return nil, 0, fmt.Errorf("can't create maxrix: end of string")
		}
		row, k, err := getRow(str[i:])
		if err != nil {
			return nil, 0, fmt.Errorf("problem with getRow: %s", err)
		}
		if len(row) == 0   {
			return nil, 0, fmt.Errorf("get empty row from %s", str[i:])
		}
		if matrix.W == 0 {
			matrix.W = len(row)
		}
		if len(row) != matrix.W {
			return nil, 0, fmt.Errorf("length of row not equal")
		}
		i = i + k + 1
		matrix.Value = append(matrix.Value, row)
		matrix.H++
		for i < len(str) && str[i] == ' ' {
			i++
		}
		if i >= len(str) {
			return nil, 0, fmt.Errorf("can't create maxrix: end of string")
		}
		if str[i] == ';' {
			i++
		} else if str[i] == ']' {
			return matrix, i, nil
		}
	}
	return nil, 0, fmt.Errorf("end of cicle")
}

package main

import (
	"fmt"
	"strconv"
)

// ConvToken convert token
func ConvToken(tokens []string, vars map[string]interface{}) (interface{}, error) {
	if tokens[0][0] >= '0' && tokens[0][0] <= '9' {
		number, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			err = fmt.Errorf("can't parse in ConvToken: %s", err)
			return nil, err
		}
		return Complex{number, 0}, nil
	}
	if IsOper(tokens[0]) {
		return Operation{tokens[0]}, nil
	}
	if tokens[0] == "i" {
		return Complex{0, 1}, nil
	}
	if IsAlpha(tokens[0][0]) {
		value, ok := vars[tokens[0]]
		if ok {
			return value, nil
		} else {
			return tokens[0], nil
		}
	}
	if tokens[0] == "=" || tokens[0] == "?" {
		return tokens[0], nil
	}
	err := fmt.Errorf("can't find token for %s", tokens[0])
	return nil, err
}

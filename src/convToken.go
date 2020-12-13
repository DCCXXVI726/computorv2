package main

import (
	"fmt"
	"strconv"
)

func ConvToken(tokens []string, vars map[string]interface{}) (interface{}, error) {
	if tokens[0][0] >= '0' && tokens[0][0] <= '9' {
		number, err := strconv.ParseFloat(tokens[0], 64)
		if err != nil {
			err = fmt.Errorf("can't parse in ConvToken: %s", err)
			return nil, err
		}
		comp := Complex{number, 0}
		return comp, nil
	}
	if tokens[0][0] == '+' {
		return Operation{"+"}, nil
	}
	err := fmt.Errorf("can't find token for %s", tokens[0])
	return nil, err
}

package main

import (
	"strconv"
)

func CheckTokens(tokens []string, vars map[string]interface{}) int {
	fi, _ := strconv.ParseFloat(tokens[0], 64)
	se, _ := strconv.ParseFloat(tokens[0], 64)
	f := Complex{fi, 0}
	s := Complex{se, 0}
	Add(f, s)
	return 0
}

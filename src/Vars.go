package main

import "fmt"

var db map[string]interface{}

func init() {
	db = make(map[string]interface{})
}

func getVar(str string) (interface{},error) {
	if v, ok := db[str]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("can't find var %s", str)
}

func setVar(str string, v interface{})  {
	db[str] = v
}

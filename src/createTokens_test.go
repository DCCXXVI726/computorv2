package main

import (
	"testing"
	"reflect"
)

type testCase struct {
	name	string
	str		string
	result	[]interface{}
	err		bool
}

func TestCreateTokens(t *testing.T) {
	tests := []testCase{
		testCase{
			name:	"simple plus",
			str:	"2 + 3",
			result:	append(make([]interface{},0), 2.0, Operation{"+", 1,}, 3.0),
			err:	false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			results, err := createTokens(tc.str)
			if !reflect.DeepEqual(results, tc.result) {
				t.Errorf("createTokens() result = %v, want %v", results, tc.result)
			}
			if (err == nil) == tc.err {
				t.Errorf("createTokens() error = %v, but it should be %v", err, tc.err)
			}
		})
	}
}

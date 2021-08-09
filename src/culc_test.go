package main

import (
	"testing"
	"reflect"
)

type testCaseCulc struct {
	name	string
	tokens	[]interface{}
	result	interface{}
	err		bool
}

func	TestCulc(t *testing.T) {
	tests := []testCaseCulc{
		testCaseCulc{
			name:	"simple plus",
			tokens:	append(make([]interface{},0), 2.0, Operation{"+", 1,}, 3.0),
			result:	5.0,
			err:	false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := culc(tc.tokens)
			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("culc() result = %v, want %v", result, tc.result)
			}
			if (err == nil) == tc.err {
				t.Errorf("culc() error = %v, but it should be %v", err, tc.err)
			}
		})
	}
}

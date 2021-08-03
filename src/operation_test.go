package main

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name	string
		i		interface{}
		k		interface{}
		result	interface{}
		err	error
	}{
		{
			name:	"Simple test",
			i:		1.0,
			k:		2.0,
			result:	3.0,
			err:	nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Sum(tc.i, tc.k)
			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("Sum() result = %v, want %v", result, tc.result)
			}
			if err != tc.err {
				t.Errorf("Sum() error = %v, want %v", err, tc.err)
			}
		})
	}
}

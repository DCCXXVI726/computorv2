package main

import (
	"testing"
	"reflect"
)

type testCaseAtof struct {
	name	string
	str		string
	pos		int
	result	float64
	err		bool
}

func TestAtof(t *testing.T) {
	tests := []testCaseAtof {
		testCaseAtof{
			name:	"zero",
			str:	"0",
			pos:	1,
			result:	0.0,
			err:	false,
		},
		testCaseAtof{
			name:	"simple dot",
			str:	"0.1",
			pos:	3,
			result:	0.1,
			err:	false,
		},
		testCaseAtof{
			name:	"positive dot",
			str:	"1.111",
			pos:	5,
			result:	1.111,
			err:	false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, pos, err := atof(tc.str)
			if !reflect.DeepEqual(result, tc.result) {
				t.Errorf("Atof() result = %f, want %f", result, tc.result)
			}
			if !reflect.DeepEqual(pos, tc.pos) {
				t.Errorf("Atof() position = %d, want %d", pos, tc.pos)
			}
			if (err == nil) == tc.err {
				t.Errorf("Atof() error = %v, but it should be %v", err, tc.err)
			}
		})
	}
}

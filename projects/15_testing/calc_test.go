package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type caseType struct {
		num1 int
		num2 int
		expected int
	}

	cases := []caseType{
		{
			num1: 1,
			num2: 3,
			expected: 4,
		},
		{
			num1: 4,
			num2: 4,
			expected: 8,
		},
	}

	for _, cas := range cases {
		result := Add(cas.num1, cas.num2)

		if result != cas.expected {
			t.Errorf("%v doesn't equal expected %v", result, cas.expected)
		}
	}

}
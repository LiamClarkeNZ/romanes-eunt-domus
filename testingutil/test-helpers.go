package testingutil

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase[T any, U any] struct {
	Input    T
	Expected U
}

func ExecuteTests[T any, U any](t *testing.T, testName string, funcUnderTest func(T) U, testCases []TestCase[T, U]) {
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%s %v", testName, tt.Input), func(t *testing.T) {
			if got := funcUnderTest(tt.Input); !reflect.DeepEqual(got, tt.Expected) {
				t.Errorf("Expected %v, actual was %v", tt.Expected, got)
			}
		})
	}
}
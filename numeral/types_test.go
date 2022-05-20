package numeral

import (
	"fmt"
	"testing"
)

type compareToTestCase struct {
	left    RomanNumeral
	right 	RomanNumeral
	expected Comparison
}


func TestCompareTo(t *testing.T) {
	one := NewRomanNumeral(1)
	five := NewRomanNumeral(5)

	tests := []compareToTestCase{
		{left: one, right: one, expected: EqualTo},
		{left: one, right: five, expected: LessThan},
		{left: five, right: one, expected: GreaterThan},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Compare %d to %d", tt.left.Value(), tt.right.Value()), func(t *testing.T) {
			if compared := tt.left.CompareTo(&tt.right); compared != tt.expected {
				t.Errorf("Expected %v but got %v", tt.expected, compared)
			}
		})
	}
}

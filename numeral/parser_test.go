package numeral

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	one         = NewRomanNumeral(1)
	five        = NewRomanNumeral(5)
	ten         = NewRomanNumeral(10)
	fifty       = NewRomanNumeral(50)
	oneHundred  = NewRomanNumeral(100)
	fiveHundred = NewRomanNumeral(500)
	oneThousand = NewRomanNumeral(1000)

	subOne        = subtractiveNumeral(1)
	subFive       = subtractiveNumeral(5)
	subTen        = subtractiveNumeral(10)
	subOneHundred = subtractiveNumeral(100)
)

func TestSingleNumerals(t *testing.T) {
	tests := []parserTestCase{
		{input: "I", expected: []RomanNumeral{one}},
		{input: "V", expected: []RomanNumeral{five}},
		{input: "X", expected: []RomanNumeral{ten}},
		{input: "L", expected: []RomanNumeral{fifty}},
		{input: "C", expected: []RomanNumeral{oneHundred}},
		{input: "D", expected: []RomanNumeral{fiveHundred}},
		{input: "M", expected: []RomanNumeral{oneThousand}},
	}
	executeTests(t, tests)
}

func TestAdditiveNumerals(t *testing.T) {
	tests := []parserTestCase{
		{input: "II", expected: []RomanNumeral{one, one}},
		{input: "III", expected: []RomanNumeral{one, one, one}},
		{input: "VII", expected: []RomanNumeral{five, one, one}},
		{input: "XII", expected: []RomanNumeral{ten, one, one}},
		{input: "XV", expected: []RomanNumeral{ten, five}},
		{input: "XXVI", expected: []RomanNumeral{ten, ten, five, one}},
		{input: "MMDCL", expected: []RomanNumeral{oneThousand, oneThousand, fiveHundred, oneHundred, fifty}},
	}
	executeTests(t, tests)
}

func TestSubtractiveNumerals(t *testing.T) {
	tests := []parserTestCase{
		{input: "IV", expected: []RomanNumeral{subOne, five}},
		{input: "IIV", expected: []RomanNumeral{subOne, subOne, five}},
		{input: "IX", expected: []RomanNumeral{subOne, ten}},
		{input: "VL", expected: []RomanNumeral{subFive, fifty}},
		{input: "CM", expected: []RomanNumeral{subOneHundred, oneThousand}},
	}
	executeTests(t, tests)
}

func TestCaseInsensitivity(t *testing.T) {
	tests := []parserTestCase{
		{input: "i", expected: []RomanNumeral{one}},
		{input: "iv", expected: []RomanNumeral{subOne, five}},
	}
	executeTests(t, tests)
}

func TestScenariosFromSpec(t *testing.T) {
	tests := []parserTestCase{
		{input: "ix", expected: []RomanNumeral{subOne, ten}},
		{input: "XIIII", expected: []RomanNumeral{ten, one, one, one, one}},
		{input: "MCMXCIX", expected: []RomanNumeral{oneThousand, subOneHundred, oneThousand, subTen, oneHundred, subOne, ten}},
	}
	executeTests(t, tests)
}

// testing utils below

func subtractiveNumeral(value int) RomanNumeral {
	n := NewRomanNumeral(value)
	n.SetContribution(Subtractive)
	return n
}

type parserTestCase struct {
	input    string
	expected []RomanNumeral
}

func executeTests(t *testing.T, testCases []parserTestCase) {
	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Parse %s", tt.input), func(t *testing.T) {
			if got := Parse(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Parse() = %v, want %v", got, tt.expected)
			}
		})
	}
}

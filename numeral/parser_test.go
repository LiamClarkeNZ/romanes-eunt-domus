package numeral

import (
	"reflect"
	tu "roman/testingutil"
	"testing"
)

type parserTestCase = tu.UnaryFuncTestCase[string, []RomanNumeral]

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
	yeah          = true
	nah       = false
)

// Wrap to ignore error and still meet single arg generic requirements.
func liberalParse(input string) []RomanNumeral {
	result, _ := Parse(input, &nah)
	return result
}


func TestSingleNumerals(t *testing.T) {
	tests := []parserTestCase{
		{Input: "I", Expected: []RomanNumeral{one}},
		{Input: "V", Expected: []RomanNumeral{five}},
		{Input: "X", Expected: []RomanNumeral{ten}},
		{Input: "L", Expected: []RomanNumeral{fifty}},
		{Input: "C", Expected: []RomanNumeral{oneHundred}},
		{Input: "D", Expected: []RomanNumeral{fiveHundred}},
		{Input: "M", Expected: []RomanNumeral{oneThousand}},
	}

	tu.RunUnaryFuncTests(t, "Parse", liberalParse, tests)
}

func TestAdditiveNumerals(t *testing.T) {
	tests := []parserTestCase{
		{Input: "II", Expected: []RomanNumeral{one, one}},
		{Input: "III", Expected: []RomanNumeral{one, one, one}},
		{Input: "VII", Expected: []RomanNumeral{five, one, one}},
		{Input: "XII", Expected: []RomanNumeral{ten, one, one}},
		{Input: "XV", Expected: []RomanNumeral{ten, five}},
		{Input: "XXVI", Expected: []RomanNumeral{ten, ten, five, one}},
		{Input: "MMDCL", Expected: []RomanNumeral{oneThousand, oneThousand, fiveHundred, oneHundred, fifty}},
	}
	tu.RunUnaryFuncTests(t, "Parse", liberalParse, tests)
}

func TestSubtractiveNumerals(t *testing.T) {
	tests := []parserTestCase{
		{Input: "IV", Expected: []RomanNumeral{subOne, five}},
		{Input: "IIV", Expected: []RomanNumeral{subOne, subOne, five}},
		{Input: "IX", Expected: []RomanNumeral{subOne, ten}},
		{Input: "VL", Expected: []RomanNumeral{subFive, fifty}},
		{Input: "CM", Expected: []RomanNumeral{subOneHundred, oneThousand}},
	}
	tu.RunUnaryFuncTests(t, "Parse", liberalParse, tests)
}

func TestCaseInsensitivity(t *testing.T) {
	tests := []parserTestCase{
		{Input: "i", Expected: []RomanNumeral{one}},
		{Input: "iv", Expected: []RomanNumeral{subOne, five}},
	}
	tu.RunUnaryFuncTests(t, "Parse", liberalParse, tests)
}

func TestScenariosFromSpec(t *testing.T) {
	tests := []parserTestCase{
		{Input: "ix", Expected: []RomanNumeral{subOne, ten}},
		{Input: "XIIII", Expected: []RomanNumeral{ten, one, one, one, one}},
		{Input: "MCMXCIX", Expected: []RomanNumeral{oneThousand, subOneHundred, oneThousand, subTen, oneHundred, subOne, ten}},
	}
	tu.RunUnaryFuncTests(t, "Parse", liberalParse, tests)
}

func TestLessLiberalParse(t *testing.T) {

	if result, _ := Parse("iiiiv", &yeah); !reflect.DeepEqual(result, []RomanNumeral{subOne, subOne, subOne, subOne, five}) {
		t.Errorf("Failed to parse iiiiv...")
	}

	if _, err := Parse("iiiiiv", &yeah); err == nil {
		t.Errorf("Expected Parse to throw an error for iiiiiv")
	}
}

// testing utils below

func subtractiveNumeral(value int) RomanNumeral {
	n := NewRomanNumeral(value)
	n.SetContribution(Subtractive)
	return n
}

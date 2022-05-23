package numeral

import (
	"errors"
	"fmt"
	"strings"
)

// Parse numerals into a slice of structs representing:
//	* the value of the individual numeral
//  * what role the numeral plays - is it additive, or subtractive
// The algorithm of determining additive/subtractive is simple - any smaller number preceding a larger
// number is considered subtractive
func Parse(romanNumerals string, postel *bool) ([]RomanNumeral, error) {
	fullyParsed := make([]RomanNumeral, 0, len(romanNumerals))
	partiallyParsed := make([]RomanNumeral, 0)

	var last RomanNumeral
	for _, char := range romanNumerals {
		converted := NewRomanNumeral(RomanNumeralValues[strings.ToUpper(string(char))])

		if last == nil {
			partiallyParsed = append(partiallyParsed, converted)
			last = converted
			continue
		}

		switch last.CompareTo(&converted) {
		case EqualTo:
			// The numbers are the same, so whether the buffered ones are additive or subtractive isn't known just yet
			partiallyParsed = append(partiallyParsed, converted)
		case GreaterThan:
			// the last digit in the buffer was bigger, so we can complete the buffer now, then buffer the current char
			fullyParsed = completePartiallyParsed(fullyParsed, partiallyParsed, Additive)
			partiallyParsed = nil
			partiallyParsed = append(partiallyParsed, converted)
		case LessThan:
			// This digit is bigger, so everything in the buffer is subtractive

			// If sanity checking is on, ensure that subtractive numerals aren't turning their root number to 0 or negative
			if *postel {
				err := sanityCheckSubtractive(&partiallyParsed, &converted)
				if err != nil {
					return nil, err
				}
			}

			fullyParsed = completePartiallyParsed(fullyParsed, partiallyParsed, Subtractive)
			partiallyParsed = nil
			partiallyParsed = append(partiallyParsed, converted)
		}
		last = converted
	}

	//If there's still items in the buffer, it's repeated characters that are additive
	if len(partiallyParsed) > 0 {
		fullyParsed = completePartiallyParsed(fullyParsed, partiallyParsed, Additive)
	}

	return fullyParsed, nil
}

// Check that subtractively prefixed numerals don't reduce current value to 0 or below e.g.,
// IIIIV is 1, that's okay
// IIIIIV is 0, not okay
// IIIIIIIV is -2, super not okay
func sanityCheckSubtractive(partiallyParsedBuffer *[]RomanNumeral, current *RomanNumeral) error {
	totalSubtractive := 0
	asRomanAgain := make([]string, len(*partiallyParsedBuffer))
	for _, pp := range *partiallyParsedBuffer {
		totalSubtractive += pp.Value()
		asRomanAgain = append(asRomanAgain, ArabicToRoman[pp.Value()])
	}

	currentVal := (*current).Value()
	if totalSubtractive >= currentVal {
		errMsg := fmt.Sprintf("%s%s converts to a number <= 0, which is most likely not intended", strings.Join(asRomanAgain, ""), ArabicToRoman[currentVal])
		return errors.New(errMsg)
	}
	return nil
}

func completePartiallyParsed(completed, partially []RomanNumeral, contribution Contribution) []RomanNumeral {
	for _, num := range partially {
		num.SetContribution(contribution)
		completed = append(completed, num)
	}
	return completed
}

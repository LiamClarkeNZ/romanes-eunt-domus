package numeral

import "strings"

// Parse numerals into a slice of structs representing:
//	* the value of the individual numeral
//  * what role the numeral plays - is it additive, or subtractive
// The algorithm of determining additive/subtractive is simple - any smaller number preceding a larger
// number is considered subtractive
func Parse(romanNumerals string) []RomanNumeral {
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

	return fullyParsed
}

func completePartiallyParsed(completed, partially []RomanNumeral, contribution Contribution) []RomanNumeral {
	for _, num := range partially {
		num.SetContribution(contribution)
		completed = append(completed, num)
	}
	return completed
}

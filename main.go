package main

import (
	"fmt"
	"os"
	"regexp"
	"roman/numeral"
	"strings"
)

func main() {
	inputNumeral := validateInput()
	parsedNumerals := numeral.Parse(inputNumeral)
	summed := sumNumerals(parsedNumerals)
	println(summed)
	println(convertToRomanNumerals(summed))
}

func validateInput() string {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s <roman number>\n", os.Args[0])
		os.Exit(1)
	}

	inputNumeral := strings.TrimSpace(os.Args[1])

	//Build (case-insensitive) regex from map keys so only one source of truth for what's a roman numeral exists
	keys := extractKeys(numeral.RomanNumeralValues)
	onlyRomanNumerals := fmt.Sprintf("(?i)^[%s]+$", strings.Join(keys, ""))

	match, _ := regexp.Match(onlyRomanNumerals, []byte(inputNumeral))

	if match == false {
		_, _ = fmt.Fprintf(os.Stderr, "Input '%s' has characters that aren't valid Roman numerals\n", inputNumeral)
		os.Exit(1)
	}
	return inputNumeral
}

func sumNumerals(numerals []numeral.RomanNumeral) int {
	sum := 0
	for _, n := range numerals {
		sum += n.Value()
	}
	return sum
}

type conversionToRoman struct {
	quantity       int
	romanSingle    string
	romanFiveTimes string
	romanNextUnit  string
}

func (p conversionToRoman) AsString() string {
	var result string

	//Only 1000 doesn't have these (and I'm studiously ignoring M̄ and friends with the multiplier line)
	if p.romanFiveTimes != "" && p.romanNextUnit != "" {
		if p.quantity == 4 {
			result = p.romanSingle + p.romanFiveTimes
		} else if p.quantity == 9 {
			result = p.romanSingle + p.romanNextUnit
		} else if p.quantity > 5 {
			afterFive := p.quantity - 5
			result = p.romanFiveTimes + strings.Repeat(p.romanSingle, afterFive)
		} else {
			result = strings.Repeat(p.romanSingle, p.quantity)
		}
	} else {
		result = strings.Repeat(p.romanSingle, p.quantity)
	}
	return result
}

func convertToRomanNumerals(number int) string {
	inverted := make(map[int]string)
	for roman, arabic := range numeral.RomanNumeralValues {
		inverted[arabic] = roman
	}

	places := []int{1000, 100, 10, 1}

	parts := separateNumberIntoParts(number)
	quantities := make(map[int]int)
	quantities[1000] = parts[0]
	quantities[100] = parts[1]
	quantities[10] = parts[2]
	quantities[1] = parts[3]

	romanStrings := make([]string, 0, 4)

	for _, place := range places {
		var romanFive string
		var romanNextUnit string

		if place != 1000 {
			romanFive = inverted[place*5]
			romanNextUnit = inverted[place*10]
		}

		conversion := conversionToRoman{
			quantity:       quantities[place],
			romanSingle:    inverted[place],
			romanFiveTimes: romanFive,
			romanNextUnit:  romanNextUnit,
		}
		romanStrings = append(romanStrings, conversion.AsString())
	}

	return strings.Join(romanStrings, "")
}

func separateNumberIntoParts(number int) []int {
	// We'll always get 4 elements back, but converting from a slice to a [4]int is surprisingly painful, but apparently
	// Go 1.19 will fix this
	return recurseThroughNumberParts(number, 1000, make([]int, 0, 4))
}

func recurseThroughNumberParts(remaining int, place int, parts []int) []int {
	//When we hit the ones, all that remains is ones
	if place == 1 {
		return append(parts, remaining)
	}

	// Taking advantage of integer division, for a number like 950, 950 / 1000 = 0,
	// so remaining = 950 - 0 * 1000 will leave it untouched
	placeAmount := remaining / place
	remaining = remaining - placeAmount*place

	//Drop to the next place
	return recurseThroughNumberParts(remaining, place/10, append(parts, placeAmount))

}

// Go 1.19 will also ship with maps.Keys() which is nice.
func extractKeys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

package main

import (
	"roman/numeral"
	tu "roman/testingutil"
	"testing"
)

func TestSumNumerals(t *testing.T) {
	testCases := []tu.UnaryFuncTestCase[[]numeral.RomanNumeral, int]{
		{Input: numeral.Parse("ix"), Expected: 9},
		{Input: numeral.Parse("XIII"), Expected: 13},
		{Input: numeral.Parse("XIIII"), Expected: 14},
		{Input: numeral.Parse("XIV"), Expected: 14},
		{Input: numeral.Parse("MCMXCIX"), Expected: 1999},
	}
	tu.RunUnaryFuncTests(t, "sumNumerals", sumNumerals, testCases)

}

func TestSeparateNumbersIntoParts(t *testing.T) {

	testCases := []tu.UnaryFuncTestCase[int, []int]{
		{Input: 8, Expected: []int{0, 0, 0, 8}},
		{Input: 12, Expected: []int{0, 0, 1, 2}},
		{Input: 125, Expected: []int{0, 1, 2, 5}},
		{Input: 9079, Expected: []int{9, 0, 7, 9}},
		{Input: 2_123_456, Expected: []int{2123, 4, 5, 6}},
	}

	tu.RunUnaryFuncTests(t, "separateNumberIntoParts", separateNumberIntoParts, testCases)
}

func TestConvertToRomanNumerals(t *testing.T) {
	testCases := []tu.UnaryFuncTestCase[int, string]{
		{Input: 9, Expected: "IX"},
		{Input: 13, Expected: "XIII"},
		{Input: 14, Expected: "XIV"},
		{Input: 1999, Expected: "MCMXCIX"},
	}

	tu.RunUnaryFuncTests(t, "convertToRomanNumerals", convertToRomanNumerals, testCases)
}

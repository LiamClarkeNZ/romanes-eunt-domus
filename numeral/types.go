package numeral

import "strconv"

var RomanNumeralValues map[string]int

func init() {
	m := make(map[string]int)
	m["M"] = 1000
	m["D"] = 500
	m["C"] = 100
	m["L"] = 50
	m["X"] = 10
	m["V"] = 5
	m["I"] = 1
	RomanNumeralValues = m
}


type Contribution int
type Comparison int

const (
	Additive Contribution = iota
	Subtractive
)

const (
	LessThan Comparison = -1
	EqualTo = 0
	GreaterThan = 3
)

type RomanNumeral interface {
	SetContribution(contribution Contribution)
	Value() int
	CompareTo(other *RomanNumeral) Comparison
	String() string
}

func NewRomanNumeral(value int) RomanNumeral {
	return &parsedNumber{value: value, contribution: Additive}
}

type parsedNumber struct {
	value int
	contribution Contribution
}

func (p *parsedNumber) SetContribution(contribution Contribution) {
	p.contribution = contribution
}

func (p parsedNumber) Value() int {
	if p.contribution == Subtractive {
		return -p.value
	} else {
		return p.value
	}
}

func (p parsedNumber) CompareTo(other *RomanNumeral) Comparison {
	ownValue := p.value
	otherValue := (*other).Value()

	var result Comparison
	if ownValue == otherValue {
		result = EqualTo
	} else if ownValue < otherValue {
		result = LessThan
	} else {
		result = GreaterThan
	}

	return result
}

func (p parsedNumber) String() string {
	return strconv.Itoa(p.Value())
}




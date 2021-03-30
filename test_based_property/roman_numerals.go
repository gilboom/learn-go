package test_based_property

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r *RomanNumerals) ValueOf(symbols ...byte) uint16 {
	roman := string(symbols)
	for _, c := range *r {
		if c.Symbol == roman {
			return c.Value
		}
	}
	return 0
}

var allRomanNumrals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	//if arabic == 3 {
	//	return "III"
	//}
	//if arabic == 2 {
	//	return "II"
	//}
	//return "I"

	//if arabic == 4 {
	//	return "IV"
	//}

	var result strings.Builder

	//for arabic > 0 {
	//	switch {
	//	case arabic > 9:
	//		result.WriteString("X")
	//		arabic = arabic - 10
	//	case arabic > 8:
	//		result.WriteString("IX")
	//		arabic = arabic - 9
	//	case arabic > 4:
	//		result.WriteString("V")
	//		arabic = arabic - 5
	//	case arabic > 3:
	//		result.WriteString("IV")
	//		arabic = arabic - 4
	//	default:
	//		result.WriteString("I")
	//		arabic = arabic - 1
	//	}
	//}

	for _, numeral := range allRomanNumrals {
		for arabic >= numeral.Value {
			arabic = arabic - numeral.Value
			result.WriteString(numeral.Symbol)
		}
	}

	return result.String()

}

func ConvertToArabic(roman string) uint16 {
	//if roman == "III" {
	//	return 3
	//}
	//if roman == "II" {
	//	return 2
	//}

	//total := 0
	//
	//for range roman {
	//	total++
	//}

	total := uint16(0)

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, roman, symbol) {

			nextSymbol := roman[i+1]

			if v := allRomanNumrals.ValueOf(symbol, nextSymbol); v != 0 {
				total += v
				i++
			} else {
				total += allRomanNumrals.ValueOf(symbol)
			}

		} else {
			total += allRomanNumrals.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubtractive(i int, roman string, currentSymbol uint8) bool {
	isSubtractiveSymbol := currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'L' || currentSymbol == 'C'
	return i+1 < len(roman) && isSubtractiveSymbol
}

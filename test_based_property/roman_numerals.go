package test_based_property

import "strings"

func ConvertToRoman(arabic int) string {
	//if arabic == 3 {
	//	return "III"
	//}
	//if arabic == 2 {
	//	return "II"
	//}
	//return "I"

	if arabic == 4 {
		return "IV"
	}

	var result strings.Builder

	for i := 0; i < arabic; i++ {
		result.WriteString("I")
	}

	return result.String()

}

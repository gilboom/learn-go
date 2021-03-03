package iteration

import "strings"

func Repeat(character string, times uint) string {
	//var repeated string
	//for i := uint(0); i < times; i++ {
	//	repeated += character
	//}
	//
	//return repeated

	return strings.Repeat(character, int(times))
}

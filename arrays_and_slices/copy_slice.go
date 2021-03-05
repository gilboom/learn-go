package arrays_and_slices

func CopyStrings(s []string) []string {
	return append(make([]string, 0, len(s)), s...)
	//return []string{}
}

//func CopyStrings(s []string) []string {
//	t := make([]string, len(s))
//	copy(t, s)
//	return t
//}

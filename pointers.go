package util

// ConvertStringSliceToPointerSlice makes a copy of each element, zeroes the existing one, and finally sets the given
// slice to nil in order to explicitly state it is no longer needed
func StringSliceToPointers(str []string) []*string {
	var idsptrs []*string
	if len(str) > 0 {
		for _, s := range str {
			func(st string) {
				idsptrs = append(idsptrs, &st)
			}(s)
			s = ""
		}
		str = nil
	}
	return idsptrs
}

// ConvertPointerSliceToStrings converts a slice of pointers to strings into a slice of strings
func PointerSliceToStrings(strPtrs []*string) []string {
	var str []string
	if len(strPtrs) > 0 {
		for _, ptr := range strPtrs {
			str = append(str, *ptr)
			ptr = nil
		}
		strPtrs = nil
	}
	return str
}

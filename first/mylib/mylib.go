package mylib

func Invert(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-i-1]
		slice[len(slice)-i-1] = temp
	}

	return slice
}

func IsAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	result := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[len(b)-1-i] {
			result = false
			break
		}
	}
	return result
}

func AppendString(stringsArray []string, valueToAppend string) []string {
	if cap(stringsArray) == 0 {
		stringsArray = make([]string, 1)
		stringsArray[0] = valueToAppend
		return stringsArray
	}

	tempStrings := make([]string, len(stringsArray)+1)
	copy(tempStrings[0:], stringsArray[0:])

	if cap(stringsArray) == len(stringsArray) {
		stringsArray = make([]string, len(stringsArray)+1, cap(stringsArray)*2)
	} else {
		stringsArray = make([]string, len(stringsArray)+1, cap(stringsArray))
	}

	copy(stringsArray[0:], tempStrings[0:])

	stringsArray[len(stringsArray)-1] = valueToAppend

	return stringsArray
}

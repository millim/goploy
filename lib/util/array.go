package util

//ArrayExists exists string in strings
func ArrayExists(sa *[]string, s string) bool {
	for _, v := range *sa {
		if v == s {
			return true
		}
	}
	return false
}

package strings

// StringInSlice returns true if the given string appears in the given slice.
func StringInSlice(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

// SubStringInSlice returns true if the given string appears in any of the strings in the given slice.
func SubStringInSlice(str string, list []string) bool {
	for _, s := range list {
		for i := 0; i < len(s)-len(str)+1; i++ {
			if s[i:i+len(str)] == str {
				return true
			}
		}
	}
	return false
}

// ExclusiveAppend appends the given string to the given slice if it does not already exist in the slice.
func ExclusiveAppend(slice []string, s string) []string {
	for _, element := range slice {
		if element == s {
			return slice
		}
	}
	return append(slice, s)
}

// ExclusiveCombine combines the given slices, excluding any strings that already exist in the first slice.
func ExclusiveCombine(slice1 []string, slice2 []string) []string {
	for _, newElement := range slice2 {
		exists := false
		for _, element := range slice1 {
			if element == newElement {
				exists = true
				break
			}
		}
		if !exists {
			slice1 = append(slice1, newElement)
		}
	}
	return slice1
}

// Exclude returns a new slice that contains all the elements of the first slice except those that appear in the second slice.
func Exclude(slice []string, remove []string) (out []string) {
	toRemove := make(map[string]bool)
	for _, str := range remove {
		toRemove[str] = true
	}
	for _, str := range slice {
		if !toRemove[str] {
			out = append(out, str)
		}
	}
	return
}

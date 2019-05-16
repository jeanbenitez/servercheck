package utils

// IndexOf return index of element into data array
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// Find returns the smallest index i at fn(n) is true, or -1
func Find(a []string, fn func(x string) bool) int {
	for i, n := range a {
		if fn(n) {
			return i
		}
	}
	return -1
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

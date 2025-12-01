package strutil

import "strconv"

// ToInt transforms a string to int. Panics on error.
func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

package hamming

import (
	"errors"
)
// Distance take two strings
// return Hamming distance between these two strings
func Distance(a, b string) (int, error) {
	var count int
	if len(a) != len(b) {
		return 0, errors.New("can't match the length of strings")
	}
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}

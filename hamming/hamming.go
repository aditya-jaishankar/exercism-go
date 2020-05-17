package hamming

import "errors"

// import "errors"

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("strings are of unequal length")
	}

	distance := 0
	for idx, char := range a {
		if string(char) != string(b[idx]) {
			distance++
		}
	}
	return distance, nil
}

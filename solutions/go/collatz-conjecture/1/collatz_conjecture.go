package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var count int

	if n == 1 {
		return count, nil
	}

	if n <= 0 {
		return count, errors.New("")
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		count++
	}

	return count, nil
}

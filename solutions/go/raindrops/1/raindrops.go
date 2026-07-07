package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	if isDivisibleBy3(number) {
		result += "Pling"
	}

	if isDivisibleBy5(number) {
		result += "Plang"
	}

	if isDivisibleBy7(number) {
		result += "Plong"
	}

	if !isDivisibleBy3(number) && !isDivisibleBy5(number) && !isDivisibleBy7(number) {
		return strconv.Itoa(number)
	}

	return result
}

func isDivisibleBy3(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}

func isDivisibleBy5(number int) bool {
	if number%5 == 0 {
		return true
	}
	return false
}

func isDivisibleBy7(number int) bool {
	if number%7 == 0 {
		return true
	}
	return false
}

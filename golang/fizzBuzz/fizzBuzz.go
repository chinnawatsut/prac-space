package fizzBuzz

import (
	"strconv"
)

func fizzBuzz(a int) string {
	var shout string = ""

	if isFizz(a) {
		shout += "Fizz"
	}

	if isBuzz(a) {
		shout += "Buzz"
	}

	if len(shout) > 0 {
		return shout
	}

	return strconv.Itoa(a)
}

func isFizz(a int) bool {
	if a%3 == 0 {
		return true
	}
	return false
}

func isBuzz(a int) bool {
	if a%5 == 0 {
		return true
	}
	return false
}

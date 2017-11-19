package fizzBuzz

import (
	"strconv"
)

func fizzBuzz(a int) string {

	if isFizz(a) && isBuzz(a) {
		return "FizzBuzz"
	}

	if isFizz(a) {
		return "Fizz"
	}

	if isBuzz(a) {
		return "Buzz"
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

package fizzBuzz

import (
	"strconv"
)

func fizzBuzz(a int) string {

	if a%3 == 0 && a%5 == 0 {
		return "FizzBuzz"
	}

	if a%3 == 0 {
		return "Fizz"
	}

	if a%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(a)
}

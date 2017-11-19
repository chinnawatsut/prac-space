package fizzBuzz

import (
	"strconv"
)

func fizzBuzz(a int) string {
	if a == 3 {
		return "fizz"
	}
	return strconv.Itoa(a)
}

package fizzBuzz

import "testing"

func TestSayFizzIfDividableBy5(t *testing.T) {
	compare(fizzBuzz(3), "Fizz", t)
	compare(fizzBuzz(6), "Fizz", t)
	compare(fizzBuzz(9), "Fizz", t)

}

func TestSayBuzzIfDividableBy5(t *testing.T) {
	compare(fizzBuzz(5), "Buzz", t)
	compare(fizzBuzz(10), "Buzz", t)
}

func TestSayFizzBuzzIfDeviddableBy3or5(t *testing.T) {
	compare(fizzBuzz(15), "FizzBuzz", t)
	compare(fizzBuzz(30), "FizzBuzz", t)

}

func TestSayAsInput(t *testing.T) {
	compare(fizzBuzz(1), "1", t)
	compare(fizzBuzz(2), "2", t)
	compare(fizzBuzz(4), "4", t)
}

func compare(result, expecting string, t *testing.T) {
	if result != expecting {
		t.Error("Expecting ", expecting, "but got", result)
	}
}

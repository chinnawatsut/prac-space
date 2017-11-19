package fizzBuzz

import "testing"

func TestSayFizz(t *testing.T) {
	compare(fizzBuzz(3), "fizz", t)
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

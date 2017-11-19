package fizzBuzz

import "testing"

func TestSay1(t *testing.T) {
	compare(fizzBuzz(1), "1", t)
}

func TestSay2(t *testing.T) {
	compare(fizzBuzz(2), "2", t)
}

func compare(result, expecting string, t *testing.T) {
	if result != expecting {
		t.Error("Expecting ", expecting, "but got", result)
	}
}

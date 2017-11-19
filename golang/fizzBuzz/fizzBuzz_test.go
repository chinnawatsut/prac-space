package fizzBuzz

import "testing"

func TestSayOne(t *testing.T) {
	if fizzBuzz(1) != "1" {
		t.Error("Expecting ", "1", "but got", fizzBuzz(1))
	}
}

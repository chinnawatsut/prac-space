package fizzBuzz

import "testing"

func TestSay1(t *testing.T) {
	if fizzBuzz(1) != "1" {
		t.Error("Expecting ", "1", "but got", fizzBuzz(1))
	}
}

func TestSay2(t *testing.T) {
	if fizzBuzz(2) != "2" {
		t.Error("Expecting ", "2", "but got", fizzBuzz(2))
	}
}

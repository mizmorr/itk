package task1

import "testing"

func TestGetThreeInts(t *testing.T) {
	ints := getThreeInts()
	if len(ints) != 3 {
		t.Error("Expected 3 integers, got", len(ints))
	}
	if ints[0] >= 1000 || ints[0] < 0 {
		t.Error("Expected integer between 0 and 999, got", ints[0])
	}
	if ints[1] >= 0o7777 || ints[1] < 0 {
		t.Error("Expected octal integer between 0 and 7777, got", ints[1])
	}
	if ints[2] >= 0xFFFF || ints[2] < 0 {
		t.Error("Expected hexadecimal integer between 0 and FFFF, got", ints[2])
	}
}

func TestTask(t *testing.T) {
	Task1()
}

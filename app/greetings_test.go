package main

import "testing"

func TestAdd(t *testing.T) {
	a := 5
	b := 3
	result := Add(3, 5)
	if result != 8 {
		t.Errorf("Add(%d, %d) = %d; want 8", a, b, result)
	}
}

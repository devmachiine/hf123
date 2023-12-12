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

func TestValidPayment(t *testing.T) {
	customerPayments := []int{5, 10, 5, 20}

	accepted := AcceptPayments(customerPayments)
	if !accepted {
		t.Errorf("expected valid customer payments %v not accepted", customerPayments)
	}
}

func TestWrongPaymentNotEnoughChange(t *testing.T) {
	customerPayments := []int{5, 20, 5, 20}

	accepted := AcceptPayments(customerPayments)
	if accepted {
		t.Errorf("expected invalid customer payments %v not accepted", customerPayments)
	}
}

func TestWrongPaymentNoCorrectChange(t *testing.T) {
	customerPayments := []int{5, 10, 10}

	accepted := AcceptPayments(customerPayments)
	if accepted {
		t.Errorf("expected invalid customer payments %v not accepted", customerPayments)
	}
}

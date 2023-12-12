package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	Hello()
}

func AcceptPayments(customerPayments []int) bool {

	avaliableChange := make(map[int]int)

	purchasePrice := 5

	for _, payment := range customerPayments {

		expectedChange := payment - purchasePrice

		avaliableChange[payment] += 1

		change, valid := composeChange(expectedChange, avaliableChange)

		if !valid {
			return false
		}

		//remove the bill(s) from avaliable change

		for bill := range change {
			avaliableChange[bill] -= 1
		}
	}

	return true
}

func composeChange(amount int, avaliableBills map[int]int) (map[int]int, bool) {

	for bill, n := range avaliableBills {
		if (bill == amount) && n > 0 {

			avaliableBills[bill] -= 1

			return avaliableBills, true
		}

		if bill > amount {
			continue
		}

		// ignore wrong denomination early problem

		avaliableBills[bill] -= 1
		amount -= bill

		return composeChange(amount, avaliableBills)
	}
}

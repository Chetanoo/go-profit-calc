package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	acquireValue(&revenue, "Enter the revenue")

	acquireValue(&expenses, "Enter the expenses")

	acquireValue(&taxRate, "Enter the tax rate")

	ebt, profit, ratio := calculateResults(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax is %.2f\n", ebt)
	fmt.Printf("Profit is %.2f\n", profit)
	fmt.Println("Ratio is", ratio)
}

func acquireValue(val *float64, text string) {
	fmt.Println(text)
	_, err := fmt.Scan(val)
	if err != nil {
		return
	}
}

func calculateResults(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

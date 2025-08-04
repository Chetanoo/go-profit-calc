package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("Enter the revenue")
	_, err := fmt.Scan(&revenue)
	if err != nil {
		return
	}

	fmt.Println("Enter the expenses")
	_, err = fmt.Scan(&expenses)
	if err != nil {
		return
	}

	fmt.Println("Enter the tax rate")
	_, err = fmt.Scan(&taxRate)
	if err != nil {
		return
	}

	ebt := revenue - expenses
	profit := ebt - (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Earning before tax is %.2f\n", ebt)
	fmt.Printf("Profit is %.2f\n", profit)
	fmt.Println("Ratio is", ratio)
}

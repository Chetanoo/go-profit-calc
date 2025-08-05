package main

import (
	"errors"
	"fmt"
	"os"
)

func writeResultsToAFile(ebt, profit, ratio float64) {
	resultString := fmt.Sprintf("Profit: %.2f\nEBT: %2f\nRatio: %2f\n", ebt, profit, ratio)
	err := os.WriteFile("result.txt", []byte(resultString), 0644)

	if err != nil {
		fmt.Println(err)
		panic("Something went wrong while writing to file")
	}
}

func main() {
	values, err := getUserInputs([3]string{"Enter the revenue",
		"Enter the expense",
		"Enter the tax rate"})

	if err != nil {
		fmt.Println(err)
		panic("Something went wrong while getting data from user")
		return
	}

	revenue, expenses, taxRate := values[0], values[1], values[2]

	ebt, profit, ratio := calculateResults(revenue, expenses, taxRate)
	writeResultsToAFile(ebt, profit, ratio)

	fmt.Printf("Earning before tax is %.2f\n", ebt)
	fmt.Printf("Profit is %.2f\n", profit)
	fmt.Printf("Ratio is %.2f\n", ratio)
}

func getUserInputs(messages [3]string) ([]float64, error) {
	var values []float64
	index := 0
	for index < len(messages) {
		var v float64
		text := messages[index]
		fmt.Println(text)
		_, err := fmt.Scan(&v)
		if err != nil {
			return []float64{}, errors.New("error while reading input")
		}
		if v <= 0 {
			return []float64{}, errors.New("value cannot be less than 1")
		}
		values = append(values, v)
		index++
	}

	return values, nil
}

func calculateResults(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

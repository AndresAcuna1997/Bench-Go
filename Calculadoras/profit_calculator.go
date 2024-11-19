package main

import (
	"errors"
	"fmt"
	"os"
)

func ProfitCalculator() {
	var err error
	var revenue float64
	var expense float64
	var taxRate float64

	revenue, err = userInput("Revenue: ")

	if err != nil {
		fmt.Print(err)
		return
	}

	expense, err = userInput("Expenses: ")

	if err != nil {
		fmt.Print(err)
		return
	}

	taxRate, err = userInput("Tax rate: ")

	if err != nil {
		fmt.Print(err)
		return
	}

	ebt, profit, ratio := calculations(revenue, expense, taxRate)

	saveValuesToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func saveValuesToFile(ebt, profit, ratio float64) {
	textFile := fmt.Sprintf("EBT: %.1f, Profit: %.1f, Ratio: %.3f", ebt, profit, ratio)
	os.WriteFile("profit.txt", []byte(textFile), 0644)
}

func calculations(revenue, expense, taxRate float64) (ebt, profit, ratio float64) {
	ebt = (revenue - expense)
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func userInput(text string) (float64, error) {
	var floatInput float64

	fmt.Print(text)
	fmt.Scan(&floatInput)

	if floatInput <= 0.0 {
		return 0, errors.New("invalid value")
	}

	return floatInput, nil
}

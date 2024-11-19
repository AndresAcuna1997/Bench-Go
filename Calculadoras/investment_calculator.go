package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	const inflationRate = 2.5

	fmt.Print("Ingrese la cantidad a invertir: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Ingrese los años de la inversion: ")
	fmt.Scan(&years)

	fmt.Print("Ingrese la tasa de retorno: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue, futureRealValue)
}

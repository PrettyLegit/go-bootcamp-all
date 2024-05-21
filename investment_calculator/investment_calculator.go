package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	fmt.Println("Investment Calculator")
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("future value: ", futureValue)
	fmt.Println("futureRealValue: ", futureRealValue)
}

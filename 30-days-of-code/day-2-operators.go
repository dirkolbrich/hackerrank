package main

import "fmt"

func main() {
	var mealCost float32
	var tipPercent, taxPercent uint32

	fmt.Scanf("%f\n%v\n%d\n", &mealCost, &tipPercent, &taxPercent)
	tip := mealCost * float32(tipPercent) / 100
	tax := mealCost * float32(taxPercent) / 100
	totalCost := mealCost + tip + tax

	fmt.Printf("The total meal cost is %.f dollars.", totalCost)
}

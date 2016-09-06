package main

import "fmt"

func main() {
	var n, e, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&e)
		sum = sum + e
	}
	fmt.Println(sum)
}

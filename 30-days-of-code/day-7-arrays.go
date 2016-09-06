package main

import "fmt"

func main() {
	var n int
	var a []string
	var s string
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		a = append(a, s)
	}

	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%v ", a[i])
	}
}

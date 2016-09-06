package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		a[i] = v
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%d ", a[i])
	}
}

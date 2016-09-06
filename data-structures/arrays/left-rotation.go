package main

import "fmt"

func main() {
	// setup
	var n, d int
	fmt.Scanf("%d %d\n", &n, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		a[i] = v
	}
	// shifting
	for i := 0; i < d; i++ {
		a = append(a[1:], a[0])
	}
	// print
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
}

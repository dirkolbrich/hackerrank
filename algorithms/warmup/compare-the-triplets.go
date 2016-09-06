package main

import "fmt"

func main() {
	// setup
	var a, b [3]int
	var alice, bob int
	fmt.Scanf("%d %d %d\n", &a[0], &a[1], &a[2])
	fmt.Scanf("%d %d %d\n", &b[0], &b[1], &b[2])

	// compare
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			alice++
		} else if a[i] < b[i] {
			bob++
		}
	}
	fmt.Printf("%d %d", alice, bob)
}

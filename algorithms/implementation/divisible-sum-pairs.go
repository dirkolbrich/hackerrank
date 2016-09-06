package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		arr[i] = v
	}

	// build pairs
	var max int
	for i := 0; i < len(arr); i++ {
		x := arr[i]
		for j := i + 1; j < len(arr); j++ {
			y := arr[j]
			if (x+y)%k == 0 {
				max++
			}
		}
	}
	fmt.Println(max)
}

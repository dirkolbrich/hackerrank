package main

import (
	"fmt"
	"math"
)

func main() {
	// setup
	var n, d1, d2 int
	var m [][]int
	fmt.Scan(&n)

	//build matrix
	for i := 0; i < n; i++ {
		var line []int
		for j := 0; j < n; j++ {
			var k int
			fmt.Scan(&k)
			line = append(line, k)
		}
		m = append(m, line)
	}

	// calculate diagonals
	for i := 0; i < n; i++ {
		d1 = d1 + m[i][i]
	}
	for i, j := 0, n-1; i < n; i, j = i+1, j-1 {
		d2 = d2 + m[i][j]
	}

	fmt.Println(math.Abs(float64(d1 - d2)))
}

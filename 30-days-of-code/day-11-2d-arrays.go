package main

import "fmt"

func main() {
	// read Stdin into slice of slices
	var m [][]int
	var n int
	for i := 0; i < 6; i++ {
		var line []int
		for j := 0; j < 6; j++ {
			fmt.Scan(&n)
			line = append(line, n)
		}
		m = append(m, line)
	}

	// counting sums in hourglasses and store in slice
	var hg []int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := m[i][j] + m[i][j+1] + m[i][j+2] + m[i+1][j+1] + m[i+2][j] + m[i+2][j+1] + m[i+2][j+2]
			hg = append(hg, sum)
		}
	}

	// get highest sum from slice
	var max int
	for i, v := range hg {
		if i == 0 {
			max = v
		}
		if v > max {
			max = v
		}
	}
	fmt.Println(max)

}

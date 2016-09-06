package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// convert decimal to binary
	var bin []int
	for n >= 1 {
		rem := n % 2
		n = n / 2
		bin = append(bin, rem)
	}
	// reverse bin
	for i, j := 0, len(bin)-1; i < j; i, j = i+1, j-1 {
		bin[i], bin[j] = bin[j], bin[i]
	}

	// test for consecutive 1's
	var sub, max int
	for _, v := range bin {
		if v == 1 {
			sub++
		}
		if sub > max {
			max = sub
		}
		if v == 0 {
			sub = 0
		}
	}

	fmt.Println(max)
}

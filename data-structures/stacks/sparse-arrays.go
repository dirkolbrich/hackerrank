package main

import "fmt"

func main() {
	// setup strings
	var n int
	fmt.Scan(&n)
	var strings = make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		strings[i] = s
	}

	// fetch queries
	var q int
	fmt.Scan(&q)
	var queries = make([]string, q)
	var hits = make([]int, q)
	for i := 0; i < q; i++ {
		var s string
		fmt.Scan(&s)
		queries[i] = s
	}

	// search
	for i, v := range queries {
		for j := 0; j < len(strings); j++ {
			if v == strings[j] {
				hits[i] = hits[i] + 1
			}
		}
	}

	// print
	for _, v := range hits {
		fmt.Println(v)
	}
}

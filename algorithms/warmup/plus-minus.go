package main

import "fmt"

func main() {
	var n, pos, neg, zero int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var k int
		fmt.Scan(&k)
		if k > 0 {
			pos++
		} else if k < 0 {
			neg++
		} else {
			zero++
		}
	}

	fmt.Printf("%.6f\n", float64(pos)/float64(n))
	fmt.Printf("%.6f\n", float64(neg)/float64(n))
	fmt.Printf("%.6f\n", float64(zero)/float64(n))
}

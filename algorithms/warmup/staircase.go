package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var s []string
		for j := n - i; j > 0; j-- {
			s = append(s, " ")
		}
		for j := 0; j < i; j++ {
			s = append(s, "#")
		}

		for _, v := range s {
			fmt.Print(v)
		}
		fmt.Print("\n")
	}
}

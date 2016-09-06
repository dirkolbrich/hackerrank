package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var S string
		var even, odd []byte
		fmt.Scan(&S)
		for i := 0; i < len(S); i++ {
			if i%2 == 0 {
				even = append(even, S[i])
			} else {
				odd = append(odd, S[i])
			}
		}
		fmt.Printf("%s %s\n", even, odd)
	}
}

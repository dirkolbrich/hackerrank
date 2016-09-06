package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scanf("%d\n", &n)

	if n%2 != 0 {
		s = "Weird"
	} else if n%2 == 0 && n >= 2 && n <= 5 {
		s = "Not Weird"
	} else if n%2 == 0 && n >= 6 && n <= 20 {
		s = "Weird"
	} else if n%2 == 0 && n > 20 {
		s = "Not Weird"
	}
	fmt.Println(s)
}

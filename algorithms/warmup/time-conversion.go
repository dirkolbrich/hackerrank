package main

import "fmt"

func main() {
	var h, m, s int
	var t string
	fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &t)

	if t == "PM" && h != 12 {
		h = h + 12
	}
	if t == "AM" && h == 12 {
		h = 0
	}
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}

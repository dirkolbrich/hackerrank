package main

import "fmt"

func main() {
	var t, pair int
	fmt.Scanf("%d", &t)

	for i := t; i > 0; i-- {
		fmt.Scanf("%d", &pair)
		draw := pair + 1
		fmt.Println(draw)
	}
}

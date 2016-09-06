package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		var stack []rune
		for _, c := range s {
			x := string(c)
			if x == "(" || x == "{" || x == "[" {
				stack = append(stack, c)
			} else if len(stack) == 0 {
				stack = append(stack, c)
				break
			} else {
				last := string(stack[len(stack)-1]) // last of stack
				y := last + x                       // combine
				if y == "()" || y == "{}" || y == "[]" {
					stack = stack[:len(stack)-1]
				}
			}
		}
		if len(stack) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

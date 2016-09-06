package main

import "fmt"

func main() {
	// setup
	var x1, x2, v1, v2 int
	fmt.Scanf("%d %d %d %d\n", &x1, &v1, &x2, &v2)

	if (v2 - v1) == 0 {
		fmt.Println("NO")
	} else if (x1-x2)%(v2-v1) == 0 && (v2-v1) < 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

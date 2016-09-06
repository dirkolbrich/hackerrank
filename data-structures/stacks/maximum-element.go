package main

import "fmt"

func main() {
	// setup
	var n int
	var s []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var t, v int
		fmt.Scanf("%d %d\n", &t, &v)
		switch t {
		case 1:
			s = append(s, v)
		case 2:
			s = s[:len(s)-1]
		case 3:
			var max int
			for _, k := range s {
				if k > max {
					max = k
				}
			}
			fmt.Println(max)
		}
	}
}

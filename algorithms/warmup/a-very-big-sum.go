package main

import "fmt"

func main() {
	var n, k, sum int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&k)
		sum = sum + k
	}
	fmt.Println(sum)
}

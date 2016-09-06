package main

import "fmt"

func main() {
	// setup
	var n, k, q int
	fmt.Scanf("%d %d %d\n", &n, &k, &q)

	// get array
	arr, pos := make([]int, n), k
	for i := 0; i < n; i++ {
		if pos >= n {
			pos = pos % n
		}
		var v int
		fmt.Scan(&v)
		arr[pos] = v
		pos++
	}

	// perform query
	for i := 0; i < q; i++ {
		var qry int
		fmt.Scan(&qry)
		fmt.Println(arr[qry])
	}
}

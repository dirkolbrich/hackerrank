package main

import "fmt"

func main() {
	// setup
	var n, q int
	fmt.Scanf("%d %d\n", &n, &q)
	seqList := make([][]int, n)
	for i := 0; i < n; i++ {
		var seq []int
		seqList[i] = seq
	}
	var lastAns int = 0

	// perform queries
	for i := 0; i < q; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d\n", &t, &x, &y)
		index := (x ^ lastAns) % n
		if t == 1 {
			seqList[index] = append(seqList[index], y)
		} else if t == 2 {
			elem := y % len(seqList[index])
			lastAns = seqList[index][elem]
			fmt.Println(lastAns)
		}
	}
}

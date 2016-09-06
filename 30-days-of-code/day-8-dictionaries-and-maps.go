package main

import "fmt"

func main() {
	var n int
	pb := make(map[string]string)

	// get number of entries and put them in a map
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var name, number string
		fmt.Scanf("%s %s\n", &name, &number)
		pb[name] = number
	}

	var query string
	for {
		_, err := fmt.Scanf("%s\n", &query)
		if err != nil {
			break
		}
		num, ok := pb[query]
		if ok {
			fmt.Printf("%s=%s\n", query, num)
		} else {
			fmt.Println("Not found")
		}
	}
}

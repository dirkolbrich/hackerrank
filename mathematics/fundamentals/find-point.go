package main

import "fmt"

func main() {
	var t, Px, Py, Qx, Qy int
	fmt.Scanf("%d", &t)

	for i := t; i > 0; i-- {
		fmt.Scanf("%d %d %d %d", &Px, &Py, &Qx, &Qy)
		fmt.Println(Px+2*(Qx-Px), Py+2*(Qy-Py))
	}
}

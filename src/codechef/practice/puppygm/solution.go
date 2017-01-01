package main

import "fmt"

func main() {
	var t, a, b int
	fmt.Scanf("%d", &t)
	for t > 0 {
		fmt.Scanf("%d %d", &a, &b)

		if a%2 == 0 || b%2 == 0 {
			fmt.Println("Tuzik")
		} else {
			fmt.Println("Vanka")
		}

		t--
	}
}

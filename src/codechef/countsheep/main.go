package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)
		r, ok := countSheep(n)
		if !ok {
			fmt.Println("INSOMNIA")
		} else {
			fmt.Println(r)
		}
	}
}

func countSheep(n int) (int, bool) {
	seen := make(map[int]bool)
}

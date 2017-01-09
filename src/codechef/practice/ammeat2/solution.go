package main

import "fmt"

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)

		if n == 1 && k == 1 {
			fmt.Println(1)
		} else {
			if 2*k > n {
				fmt.Println(-1)
			} else {
				fmt.Print(2)
				x := 4
				for k > 1 {
					fmt.Printf(" %d", x)
					x += 2
					k--
				}
				fmt.Println()
			}
		}

		t--
	}
}

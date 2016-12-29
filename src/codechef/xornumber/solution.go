package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		t--

		var n int
		fmt.Scanf("%d", &n)

		if n == 1 {
			fmt.Println(2)
			continue
		}

		i := -1
		for n > 0 {
			if n&1 == 0 {
				break
			}

			n = n >> 1
			i++
		}

		if n > 0 {
			fmt.Println(-1)
			continue
		}
		m := 1<<uint(i) - 1
		fmt.Println(m)
	}
}

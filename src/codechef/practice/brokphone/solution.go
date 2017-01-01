package main

import "fmt"

func main() {
	var t, n, x, y int
	fmt.Scanf("%d", &t)
	for t > 0 {
		t--
		fmt.Scanf("%d", &n)
		var ans int
		var prevWrong bool
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &x)

			if i > 0 && x != y {
				ans += 2
				if prevWrong {
					ans--
				}
				prevWrong = true
			} else {
				prevWrong = false
			}

			y = x
		}

		fmt.Println(ans)
	}
}

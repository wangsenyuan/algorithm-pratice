package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for t > 0 {
		var n int
		fmt.Scanf("%d", &n)

		nums := make([]int, 3)
		flag := true
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &nums[i%3])
			if i >= 2 && flag {
				a, b, c := nums[(i-2)%3], nums[(i-1)%3], nums[i%3]
				flag = !((c < a && c < b) || (a > b && a > c))
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}

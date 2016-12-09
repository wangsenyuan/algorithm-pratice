package main

import "fmt"

func main() {
	t := 0
	fmt.Scanf("%d\n", &t)

	for t > 0 {
		n := 0
		fmt.Scanf("%d\n", &n)
		num := 0
		res := 0
		found := make(map[int]bool)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &num)
			if !found[num] {
				res++
			}
			found[num] = true
		}

		fmt.Println(res)

		t--
	}
}

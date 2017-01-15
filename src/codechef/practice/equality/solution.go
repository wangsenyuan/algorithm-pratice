package main

import "fmt"

func main() {

	var t, n int

	fmt.Scanf("%d", &t)

	for t > 0 {
		fmt.Scanf("%d", &n)
		as := make([]int64, n)
		s := int64(0)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &as[i])
			s += as[i]
		}

		sx := s / int64(n - 1)

		for i := 0; i < n; i++ {
			fmt.Printf("%d", sx-as[i])
			if i == n-1 {
				fmt.Println()
			} else {
				fmt.Print(" ")
			}
		}

		t--
	}
}

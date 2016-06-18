package main

import "fmt"

func main() {
	n := 0
	k := 0
	fmt.Scanf("%d %d", &n, &k)

	xs := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &xs[i])
	}

	for i := 0; i < k; i++ {
		for j := 0; j < n-i-1; j++ {
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d", xs[i])

		if i < n-1 {
			fmt.Printf(" ")
		}
	}
}

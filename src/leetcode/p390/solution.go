package main

func lastRemaining(n int) int {
	array := make([]int, n)

	for i := 1; i <= n; i++ {
		array[i-1] = i
	}
}

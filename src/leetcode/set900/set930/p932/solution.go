package main

import "fmt"

func beautifulArray(N int) []int {
	if N == 1 {
		return []int{1}
	}
	left := beautifulArray(N / 2)
	right := beautifulArray(N - N/2)

	res := make([]int, N)
	for i := 0; i < len(left); i++ {
		res[i] = left[i] * 2
	}
	for i := 0; i < len(right); i++ {
		res[len(left)+i] = right[i]*2 - 1
	}
	return res
}

func isBeautiful(A []int) bool {
	n := len(A)
	cnt := make([]int, n+1)
	for i := 0; i < n; i++ {
		if A[i] < 1 || A[i] > n {
			return false
		}
		cnt[A[i]]++
	}

	for i := 1; i <= n; i++ {
		if cnt[i] == 0 {
			return false
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 2; j < n; j++ {
			for k := i + 1; k < j; k++ {
				if A[k]*2 == A[i]+A[j] {
					return false
				}
			}
		}
	}
	return true
}

func generate(N int) {
	A := make([]int, N)

	var loop func(i int, mask int)

	loop = func(i int, mask int) {
		if i == N {
			if isBeautiful(A) {
				fmt.Println(A)
			}
			return
		}
		for j := 0; j < N; j++ {
			if mask&(1<<uint(j)) == 0 {
				A[i] = j + 1
				loop(i+1, mask|(1<<uint(j)))
			}
		}
	}

	loop(0, 0)

}

func main() {
	generate(11)
}

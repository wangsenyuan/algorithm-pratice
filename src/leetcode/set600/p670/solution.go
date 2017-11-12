package main

import "fmt"

func maximumSwap(num int) int {
	xs := make([]int, 31)
	j := 0
	for num > 0 {
		xs[j] = num % 10
		num /= 10
		j++
	}

	st := make([]int, j)
	st[0] = 0

	for k := 1; k < j; k++ {
		if xs[k] > xs[st[k-1]] {
			st[k] = k
		} else {
			st[k] = st[k-1]
		}
	}

	for k := j - 1; k > 0; k-- {
		if st[k] != k && xs[st[k]] != xs[k] {
			xs[st[k]], xs[k] = xs[k], xs[st[k]]
			break
		}
	}

	res := 0
	for k := j - 1; k >= 0; k-- {
		res = res*10 + xs[k]
	}

	return res
}

func main() {
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(2727))
	fmt.Println(maximumSwap(9973))
	fmt.Println(maximumSwap(89796))
}

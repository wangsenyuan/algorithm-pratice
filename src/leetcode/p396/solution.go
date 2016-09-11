package main

import "fmt"

func main() {
	A := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(A))
}

func maxRotateFunction(A []int) int {
	sum := int64(0)
	f := int64(0)
	for i, a := range A {
		sum += int64(a)
		f += int64(i) * int64(a)
	}
	n := len(A)
	res := f
	for k := 1; k < n; k++ {
		f = f + sum - int64(n)*int64(A[n-k])
		if f > res {
			res = f
		}
	}

	return int(res)
}

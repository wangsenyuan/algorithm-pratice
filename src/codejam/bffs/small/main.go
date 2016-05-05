package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func main() {
	var t int
	fmt.Scanf("%d", &t)

	for i := 1; i <= t; i++ {
		var n int
		fmt.Scanf("%d", &n)

		f := make([]int, 0, n)

		for i := 0; i < n; i++ {
			var x = 0
			fmt.Scanf("%d", &x)
			f = append(f, x-1)
		}

		ps := permutations(sliceRange(0, n))
		//fmt.Printf("%s\n", concat(f, ","))
		r := play(f, ps)

		fmt.Printf("Case #%d: %d\n", i, r)
	}
}

func sliceRange(start, end int) []int {
	slice := make([]int, 0, end-start)

	for i := start; i < end; i++ {
		slice = append(slice, i)
	}

	return slice
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func play(f []int, ps [][]int) int {
	r := 0
	n := len(f)
	rc := make(chan int, n)
	defer close(rc)

	for _, p := range ps {
		//fmt.Printf("%s\n", concat(p, ","))
		go func(p []int) {
			rc <- checkPermutation(f, p, n)
		}(p)
	}

	for i := 0; i < len(ps); i++ {
		x := <-rc
		r = max(r, x)
	}

	return r
}

func checkPermutation(f []int, p []int, n int) int {
	start := time.Now()
	first := p[0]
	second := p[1]
	var r = 0
	for i := 1; i < n; i++ {
		prev := p[i-1]
		cur := p[i]

		if (f[cur] == first || f[cur] == prev) &&
			(f[first] == cur || f[first] == second) {
			r = max(r, i+1)
		}

		if f[cur] != prev && (i == n-1 || f[cur] != p[i+1]) {
			break
		}
	}
	//fmt.Printf("take %s to process %s\n", time.Since(start), concat(p, ", "))
	return r
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func concat(ints []int, sep string) string {
	var buffer bytes.Buffer
	_sep := ""
	for _, x := range ints {
		buffer.WriteString(_sep)
		buffer.WriteString(strconv.Itoa(x))
		_sep = sep
	}
	return buffer.String()
}

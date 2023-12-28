package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)

	a := make([]int, 20)

	for tc > 0 {
		tc--
		var n int
		fmt.Fscan(in, &n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}
		res := solve(a[:n])
		fmt.Fprintln(out, len(res))
		for i := 0; i < len(res); i++ {
			fmt.Fprintln(out, res[i][0]+1, res[i][1]+1)
		}
	}

}

func solve(a []int) [][]int {
	n := len(a)
	neg, pos, maxI, minI := 0, 0, 0, 0
	for i, v := range a {
		if v < 0 {
			neg++
		} else if v > 0 {
			pos++
		}
		if v > a[maxI] {
			maxI = i
		} else if v < a[minI] {
			minI = i
		}
	}

	var ans [][]int
	makePos := func() {
		for i, v := range a {
			if v < 0 {
				ans = append(ans, []int{i, maxI})
			}
		}
		for i := 1; i < n; i++ {
			ans = append(ans, []int{i, i - 1})
		}
	}
	makeNeg := func() {
		for i, v := range a {
			if v > 0 {
				ans = append(ans, []int{i, minI})
			}
		}
		for i := n - 2; i >= 0; i-- {
			ans = append(ans, []int{i, i + 1})
		}
	}

	if a[maxI] >= -a[minI] {
		if neg <= 12 {
			makePos()
		} else {
			for i := 0; i < 5; i++ {
				ans = append(ans, []int{minI, minI})
			}
			makeNeg()
		}
	} else {
		if pos <= 12 {
			makeNeg()
		} else {
			for i := 0; i < 5; i++ {
				ans = append(ans, []int{maxI, maxI})
			}
			makePos()
		}
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

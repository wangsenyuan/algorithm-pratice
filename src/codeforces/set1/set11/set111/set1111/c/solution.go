package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	nums := readNNums(reader, 4)
	n, A, B := nums[0], nums[2], nums[3]
	a := readNNums(reader, nums[1])
	return solve(n, A, B, a)
}

func solve(n int, A int, B int, a []int) int {
	// A burn the current base (单数)
	// 这个题目的说法绝对是有毒的
	sort.Ints(a)
	for i := 0; i < len(a); i++ {
		a[i]--
	}
	count := func(l int, r int) int {
		// 在范围l...r中有多少
		i := search(len(a), func(i int) bool {
			return a[i] >= l
		})

		j := search(len(a), func(j int) bool {
			return a[j] >= r
		})

		return j - i
	}

	var dfs func(i int, k int) int

	dfs = func(i int, k int) int {
		m := count(i, i+k)

		if m == 0 {
			return A
		}

		if k == 1 {
			return B * m
		}

		ans := B * m * k
		ansl := dfs(i, k/2)
		ansr := dfs(i+k/2, k/2)
		return min(ans, ansl+ansr)
	}

	return dfs(0, 1<<n)
}

func search(n int, f func(int) bool) int {
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

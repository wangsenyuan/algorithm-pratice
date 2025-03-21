package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

const inf = 1 << 30

func solve(a []int) int {
	const k = 1
	f := [k + 1][]int{}
	for i := range f {
		f[i] = []int{-inf}
	}
	for _, v := range a {
		for i := k; i >= 0; i-- {
			j := sort.SearchInts(f[i], v)
			if j < len(f[i]) {
				f[i][j] = v
			} else {
				f[i] = append(f[i], v)
			}
			if i > 0 {
				g := f[i-1]
				j = len(g)
				w := g[len(g)-1] + 1
				if j < len(f[i]) {
					f[i][j] = min(f[i][j], w)
				} else {
					f[i] = append(f[i], w)
				}
			}
		}
	}
	ans := 0
	for _, g := range f {
		ans = max(ans, len(g)-1)
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(a []int) int {
	n := len(a)
	// pref[i] = lis until i without changing
	pref := make([]int, n)
	dp := make([]int, n)
	arr := make([]int, n)
	var pos int

	var ans int

	for i := 0; i < n; i++ {
		j := sort.Search(pos, func(j int) bool {
			return a[arr[j]] >= a[i]
		})
		dp[i] = j + 1
		if j == 0 {
			pref[i] = -1
		} else {
			pref[i] = arr[j-1]
		}
		ans = max(ans, dp[i])

		if i+1 < n {
			ans = max(ans, dp[i]+1)
		}

		if j == pos || a[arr[j]] != a[i] {
			arr[j] = i
		}
		if j == pos {
			pos++
		}
	}

	pos = 0

	for i := n - 1; i >= 0; i-- {
		j := sort.Search(pos, func(j int) bool {
			return a[arr[j]] <= a[i]
		})
		if i > 0 {
			ans = max(ans, j+2)
		}
		// j = 后面的lis
		if pref[i] >= 0 && pref[i]+1 < i && a[pref[i]]+1 < a[i] {
			// 中间有空间把数字变成它们中间的数
			ans = max(ans, dp[i]+j+1)
		}
		arr[j] = i
		if j == pos {
			pos++
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

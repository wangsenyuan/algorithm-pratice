package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) (craters [][]int, res []int) {
	n := readNum(reader)
	craters = make([][]int, n)
	for i := range n {
		craters[i] = readNNums(reader, 2)
	}
	res = solve(craters)
	return
}

func solve(craters [][]int) []int {

	var arr []int
	for _, cur := range craters {
		l, r := cur[0]-cur[1], cur[0]+cur[1]
		arr = append(arr, l)
		arr = append(arr, r)
	}

	sort.Ints(arr)
	var m int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i] > arr[i-1] {
			arr[m] = arr[i-1]
			m++
		}
	}
	arr = arr[:m]

	fp := make([][]int, m)
	dp := make([][]int, m)
	from := make([][]int, m)
	for i := 0; i < m; i++ {
		fp[i] = make([]int, m)
		dp[i] = make([]int, m)
		from[i] = make([]int, m)
		for j := range m {
			from[i][j] = -1
			dp[i][j] = -1
		}
	}

	at := make([][]int, m)

	for i, cur := range craters {
		l, r := cur[0]-cur[1], cur[0]+cur[1]
		l = sort.SearchInts(arr, l)
		r = sort.SearchInts(arr, r)
		fp[l][r] = i + 1
		at[l] = append(at[l], r)
	}

	for i := 0; i < m; i++ {
		sort.Ints(at[i])
	}
	var dfs func(l int, r int) int

	dfs = func(l int, r int) int {
		if l >= r {
			return 0
		}
		if dp[l][r] >= 0 {
			return dp[l][r]
		}

		dp[l][r] = dfs(l+1, r)
		from[l][r] = l + 1
		for _, i := range at[l] {
			if i >= r {
				break
			}
			if dp[l][r] < dfs(l, i)+dfs(i, r) {
				dp[l][r] = dp[l][i] + dp[i][r]
				from[l][r] = i
			}
		}
		if fp[l][r] > 0 {
			dp[l][r]++
		}
		return dp[l][r]
	}

	dfs(0, m-1)

	var res []int
	var dfs2 func(l int, r int)
	dfs2 = func(l int, r int) {
		if l >= r || dp[l][r] <= 0 {
			return
		}
		if fp[l][r] > 0 {
			res = append(res, fp[l][r])
		}

		i := from[l][r]

		if i > l && i < r {
			dfs2(l, i)
			dfs2(i, r)
		}
	}

	dfs2(0, m-1)
	sort.Ints(res)
	return res
}

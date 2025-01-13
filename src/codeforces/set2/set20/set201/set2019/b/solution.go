package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d ", x))
		}
		buf.WriteByte('\n')
	}

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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	x := readNNums(reader, n)
	queries := readNNums(reader, m)
	return solve(x, queries)
}

func solve(x []int, queries []int) []int {
	// k
	n := len(x)
	// 有多少个点，正好被包含在k个区间中
	// 考虑任意一个x[i] (0 开始)
	// 它和前面有(i+1)点有连接
	// 后面有 (n - i)个点有连线 (i + 1) * (n - i)个区间覆盖
	// 然后考虑i左边的点，这些点被i * (n - i)个区间覆盖
	type pair struct {
		first  int
		second int
	}

	var arr []pair

	for i := range n {
		arr = append(arr, pair{(i+1)*(n-i) - 1, 1})
		if i > 0 {
			arr = append(arr, pair{i * (n - i), x[i] - x[i-1] - 1})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	m := len(arr)
	sum := make([]int, m+1)
	for i := range m {
		sum[i+1] = sum[i] + arr[i].second
	}

	ans := make([]int, len(queries))

	for i, k := range queries {
		l := search(m, func(j int) bool {
			return arr[j].first >= k
		})
		if l == m || arr[l].first > k {
			ans[i] = 0
			continue
		}
		// arr[l].first = k
		r := search(m, func(j int) bool {
			return arr[j].first > k
		})
		ans[i] = sum[r] - sum[l]
	}

	return ans
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
	return r
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%.10f\n", x))
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

func process(reader *bufio.Reader) []float64 {
	n := readNum(reader)
	p := make([]int, n)
	a := make([]int, n)
	for i := range n {
		p[i], a[i] = readTwoNums(reader)
	}
	k := readNum(reader)
	queries := make([]int, k)
	for i := range k {
		queries[i] = readNum(reader)
	}
	return solve(p, a, queries)
}

func solve(p []int, a []int, queries []int) []float64 {
	n := len(p)
	left := make([]int, n)
	right := make([]int, n)
	for i := range n {
		left[i] = -1
		right[i] = -1
	}
	root := 0
	for i := range n {
		if p[i] > 0 {
			j := p[i] - 1
			if left[j] < 0 {
				left[j] = i
			} else if a[j] < a[left[j]] {
				right[j] = left[j]
				left[j] = i
			} else {
				right[j] = i
			}
		} else {
			root = i
		}
	}
	type pair struct {
		first  int
		second int
	}
	m := len(queries)
	arr := make([]pair, m)
	for i, k := range queries {
		arr[i] = pair{k, i}
	}
	slices.SortFunc(arr, func(x, y pair) int {
		return x.first - y.first
	})

	depth := make([]int, n)
	val := make([]pair, n)

	var dfs func(u int)
	dfs = func(u int) {
		if left[u] < 0 && right[u] < 0 {
			// leaf node
			val[u] = pair{a[u], a[u]}
			return
		}
		depth[left[u]] = depth[u] + 1
		depth[right[u]] = depth[u] + 1
		dfs(left[u])
		dfs(right[u])
		val[u].first = val[left[u]].first
		val[u].second = val[right[u]].second
	}
	dfs(root)
	sum := make([]int, m+1)
	div := make([]int, m+1)
	var dfs2 func(u int, l int, r int)
	dfs2 = func(u int, l int, r int) {
		if left[u] < 0 && right[u] < 0 {
			for i := l; i <= r; i++ {
				div[i] = depth[u]
			}
			return
		}

		mid := sort.Search(m, func(i int) bool {
			return arr[i].first >= a[u]
		})
		// arr[mid] > a[u]
		// l...mid-1 全部 右子树的最小值
		if l < mid {
			sum[l] += val[right[u]].first
			sum[mid] -= val[right[u]].first
			dfs2(left[u], l, mid-1)
		}
		if mid <= r {
			sum[mid] += val[left[u]].second
			sum[r+1] -= val[left[u]].second
			dfs2(right[u], mid, r)
		}
	}

	dfs2(root, 0, m-1)
	for i := 1; i < m; i++ {
		sum[i] += sum[i-1]
	}
	ans := make([]float64, m)
	for i := 0; i < m; i++ {
		ans[arr[i].second] = float64(sum[i]) / float64(div[i])
	}
	return ans
}

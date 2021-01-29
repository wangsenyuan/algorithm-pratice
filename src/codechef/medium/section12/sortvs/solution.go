package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		P := readNNums(reader, n)
		X := make([][]int, m)
		for i := 0; i < m; i++ {
			X[i] = readNNums(reader, 2)
		}
		res := solve(n, m, P, X)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n, m int, P []int, X [][]int) int {
	set := NewDSU(n)

	for _, x := range X {
		u, v := x[0], x[1]
		u--
		v--
		set.Union(u, v)
	}

	E := make([][]int, 0, n)

	for i := 0; i < n; i++ {
		x := set.Find(i)
		y := set.Find(P[i] - 1)
		if x != y {
			E = append(E, []int{x, y})
		}
	}
	k := len(E)

	vis := make([]int, n)
	nxt := make([]int, n)
	var dfs func(u int) int
	dfs = func(u int) int {
		vis[u]++
		if nxt[u] < 0 {
			return 0
		}
		if vis[nxt[u]] == 0 {
			return 1 + dfs(nxt[u])
		}
		return 1
	}

	degree := make([]int, n)
	isCycle := func(mask int) bool {
		for i := 0; i < n; i++ {
			vis[i] = 0
			degree[i] = 0
			nxt[i] = -1
		}

		for i := 0; i < k; i++ {
			if mask&(1<<uint(i)) > 0 {
				u, v := E[i][0], E[i][1]
				degree[u]++
				degree[v]++
				nxt[u] = v
			}
		}
		var cnt int
		// all degree need to be 2 if set
		for i := 0; i < n; i++ {
			if degree[i] != 0 && degree[i] != 2 {
				return false
			}
			if degree[i] == 2 {
				cnt++
			}
		}
		// check if connected
		var cnt2 int

		for i := 0; i < n; i++ {
			if degree[i] == 2 {
				cnt2 = dfs(i)
				break
			}
		}
		return cnt == cnt2
	}

	// need to cover all k
	K := 1 << uint(k)
	dp := make([]int, K)

	for i := 1; i < K; i++ {
		if isCycle(i) {
			dp[i] = 1
		}
	}

	for i := 1; i < K; i++ {
		for j := i; j > 0; j = (j - 1) & i {
			if dp[j] > 0 && dp[i^j] > 0 {
				dp[i] = max(dp[i], dp[j]+dp[i^j])
			}
		}
	}
	return k - dp[K-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type DSU struct {
	arr  []int
	cnt  []int
	find func(int) int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}

	var find func(x int) int
	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	return &DSU{arr, cnt, find}
}

func (dsu *DSU) Find(x int) int {
	return dsu.find(x)
}

func (dsu *DSU) Union(x, y int) bool {
	x = dsu.find(x)
	y = dsu.find(y)
	if x == y {
		return false
	}
	if dsu.cnt[x] < dsu.cnt[y] {
		x, y = y, x
	}
	dsu.cnt[x] += dsu.cnt[y]
	dsu.arr[y] = x
	return true
}

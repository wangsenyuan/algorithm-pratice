package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	P := make([][]int, n)

	for i := 0; i < n; i++ {
		P[i] = readNNums(reader, 2)
	}

	k, res := solve(n, P)

	fmt.Println(k)

	var buf bytes.Buffer

	for i := 0; i < len(res); i++ {
		buf.WriteString(strconv.Itoa(res[i]))
		buf.WriteByte(' ')
	}

	fmt.Println(buf.String())
}

const MAX_MASK = 1 << 20

type Pair struct {
	first, second int
}

var G [][]Pair
var vis []bool

func init() {
	G = make([][]Pair, MAX_MASK)
	for i := 0; i < MAX_MASK; i++ {
		G[i] = make([]Pair, 0, 10)
	}
	vis = make([]bool, MAX_MASK)
}

func solve(n int, P [][]int) (int, []int) {
	for i := 20; i >= 0; i-- {
		mask := 1<<i - 1
		ok, res := find(n, P, mask)
		if ok {
			return i, res
		}
	}
	return 0, nil
}

func find(n int, P [][]int, mask int) (bool, []int) {
	for i := 0; i <= mask; i++ {
		G[i] = G[i][:0]
		vis[i] = false
	}

	for i := 1; i <= n; i++ {
		u := P[i-1][0] & mask
		v := P[i-1][1] & mask
		G[u] = append(G[u], Pair{v, 2*i - 1})
		G[v] = append(G[v], Pair{u, 2*i - 2})
	}

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for _, v := range G[u] {
			if !vis[v.first] {
				dfs(v.first)
			}
		}
	}

	var cnt int

	for i := 0; i <= mask && cnt < 2; i++ {
		if len(G[i])&1 == 1 {
			return false, nil
		}
		if !vis[i] && len(G[i]) > 0 {
			dfs(i)
			cnt++
		}
	}

	if cnt != 1 {
		return false, nil
	}

	var dfs2 func(u int, prev int)
	ans := make([]int, 0, 2*n)

	dfs2 = func(u int, prev int) {
		for i := 0; i < len(G[u]); i++ {
			if vis[G[u][i].second/2] {
				continue
			}
			vis[G[u][i].second/2] = true
			dfs2(G[u][i].first, G[u][i].second)
		}

		if prev != -1 {
			ans = append(ans, prev)
			ans = append(ans, prev^1)
		}
	}

	for i := 0; i <= n; i++ {
		vis[i] = false
	}

	for i := 0; i <= mask; i++ {
		if len(G[i]) > 0 {
			dfs2(i, -1)
			break
		}
	}

	for i := 0; i < len(ans); i++ {
		ans[i]++
	}

	return true, ans
}

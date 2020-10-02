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
		n := readNum(reader)
		E := readNNums(reader, n)
		res := solve(n, E)
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

const MOD = 1000000007

const N = 100010

var vis [N]int
var visit [N]int64
var notVisit [N]int64
var inCycle [N]bool

func solve(n int, enemies []int) int {
	for i := 0; i < n; i++ {
		vis[i] = 0
		inCycle[i] = false
		visit[i] = 1
		notVisit[i] = 1
	}
	//vis := make([]int, n)
	//inCycle := make([]bool, n)
	var dfs1 func(u int, cycle *[]int) int

	dfs1 = func(u int, cycle *[]int) int {
		if vis[u] == 2 {
			return -1
		}

		if vis[u] == 1 {
			return u
		}
		vis[u]++

		res := dfs1(enemies[u]-1, cycle)

		vis[u]++
		if res >= 0 {
			inCycle[u] = true
			*cycle = append(*cycle, u)
		}
		if res == u {
			return -1
		}
		return res
	}
	cycles := make([][]int, 0, 3)
	for i := 0; i < n; i++ {
		if vis[i] == 0 {
			tmp := make([]int, 0, n)
			dfs1(i, &tmp)
			if len(tmp) > 0 {
				cycles = append(cycles, tmp)
			}
		}
	}

	//visit := make([]int64, n)
	//notVisit := make([]int64, n)

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, 3)
	}

	for i := 0; i < n; i++ {
		adj[i] = append(adj[i], enemies[i]-1)
		adj[enemies[i]-1] = append(adj[enemies[i]-1], i)
	}

	var dfs2 func(p, u int)

	dfs2 = func(p, u int) {
		visit[u] = 1
		notVisit[u] = 1

		for _, v := range adj[u] {
			if p == v || inCycle[v] {
				continue
			}
			dfs2(u, v)
			visit[u] = (visit[u] * notVisit[v]) % MOD
			notVisit[u] = (notVisit[u] * (visit[v] + notVisit[v]) % MOD) % MOD
		}
	}

	for i := 0; i < n; i++ {
		if inCycle[i] {
			dfs2(-1, i)
		}
	}

	var ans int64 = 1

	for _, cyc := range cycles {
		ans = (ans * calc(cyc)) % MOD
	}

	return int(ans)
}

func calc(cycle []int) int64 {

	var a, b int64 = visit[cycle[0]], 0

	for i := 1; i < len(cycle); i++ {
		var na, nb int64
		if i == len(cycle)-1 {
			na = 0
			nb = (notVisit[cycle[len(cycle)-1]] * (a + b) % MOD) % MOD
		} else {
			na = (visit[cycle[i]] * b) % MOD
			nb = (notVisit[cycle[i]] * (a + b)) % MOD
		}
		a, b = na, nb
	}
	var ans1 = (a + b) % MOD

	a, b = 0, notVisit[cycle[0]]

	for i := 1; i < len(cycle); i++ {
		na := (visit[cycle[i]] * b) % MOD
		nb := (notVisit[cycle[i]] * (a + b)) % MOD
		a, b = na, nb
	}
	ans2 := (a + b) % MOD

	return (ans1 + ans2) % MOD
}

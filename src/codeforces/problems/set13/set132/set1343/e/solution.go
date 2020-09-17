package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		firstLine := readNNums(reader, 5)
		n, m, a, b, c := firstLine[0], firstLine[1], firstLine[2], firstLine[3], firstLine[4]
		P := readNNums(reader, m)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(reader, 2)
		}
		fmt.Println(solve(n, m, a, b, c, P, edges))
	}
}

func solve(n, m, a, b, c int, P []int, edges [][]int) int64 {
	sort.Ints(P)

	pref := make([]int64, m+1)

	for i := 0; i < m; i++ {
		pref[i+1] = pref[i] + int64(P[i])
	}

	E := make([][]int, n)

	for i := 0; i < n; i++ {
		E[i] = make([]int, 0, 10)
	}

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		E[u] = append(E[u], v)
		E[v] = append(E[v], u)
	}

	da := make([]int, n)
	db := make([]int, n)
	dc := make([]int, n)

	for i := 0; i < n; i++ {
		da[i] = -1
		db[i] = -1
		dc[i] = -1
	}

	que := make([]int, n)

	bfs := func(x int, dist []int) {
		var front, end int
		que[end] = x
		dist[x] = 0
		end++
		for front < end {
			u := que[front]
			front++
			for _, v := range E[u] {
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					que[end] = v
					end++
				}
			}
		}
	}

	bfs(a-1, da)
	bfs(b-1, db)
	bfs(c-1, dc)

	var best = int64(-1)

	for i := 0; i < n; i++ {
		if da[i]+db[i]+dc[i] > m {
			continue
		}
		tmp := pref[db[i]] + pref[da[i]+db[i]+dc[i]]
		if best < 0 || best > tmp {
			best = tmp
		}
	}

	return best
}

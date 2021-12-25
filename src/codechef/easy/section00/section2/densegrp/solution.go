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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		first := readNNums(reader, 4)
		m := first[1]
		x, y := first[2], first[3]
		A := make([]int, m)
		B := make([]int, m)
		C := make([]int, m)
		D := make([]int, m)
		for i := 0; i < m; i++ {
			line := readNNums(reader, 4)
			A[i] = line[0]
			B[i] = line[1]
			C[i] = line[2]
			D[i] = line[3]
		}
		res := solve(x, y, A, B, C, D)
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

func solve(x, y int, A, B, C, D []int) int {
	m := len(A)
	vecs := make([]int, 4*m)
	mp := make(map[int]int)
	mp[x]++
	mp[y]++

	for i := 0; i < m; i++ {
		vecs[4*i] = A[i]
		vecs[4*i+1] = B[i]
		vecs[4*i+2] = C[i]
		vecs[4*i+3] = D[i]
		mp[A[i]]++
		mp[B[i]]++
		mp[C[i]]++
		mp[D[i]]++
	}
	mp = compressAndSort(mp)
	for i := 0; i < len(vecs); i++ {
		vecs[i] = mp[vecs[i]]
	}

	x = mp[x]
	y = mp[y]

	n := len(mp)

	adj := make([][]Pair, 6*n+m+1)
	for i := 0; i < len(adj); i++ {
		adj[i] = make([]Pair, 0, 2)
	}

	var build func(v int, start, end int)
	var xvec, yvec int
	build = func(v int, start, end int) {
		adj[v/2] = append(adj[v/2], Pair{v, 0})
		if start == end {
			adj[v] = append(adj[v], Pair{4*n + v/2, 0})

			if start == x {
				xvec = v
			}
			if start == y {
				yvec = v
			}
			return
		}
		adj[4*n+v] = append(adj[4*n+v], Pair{4*n + v/2, 0})
		mid := (start + end) / 2
		build(v*2, start, mid)
		build(v*2+1, mid+1, end)
	}
	build(1, 1, n)

	var search func(v int, start, end int, l, r int, i int, state bool)

	search = func(v int, start, end int, l, r int, i int, state bool) {
		if end < l || r < start {
			return
		}
		if l <= start && end <= r {
			if state {
				adj[6*n+i] = append(adj[6*n+i], Pair{v, 1})
			} else if start == end {
				adj[v] = append(adj[v], Pair{6*n + i, 1})
			} else {
				adj[4*n+v] = append(adj[4*n+v], Pair{6*n + i, 1})
			}
			return
		}
		mid := (start + end) / 2
		search(2*v, start, mid, l, r, i, state)
		search(2*v+1, mid+1, end, l, r, i, state)
	}
	for i, cnt := 0, 0; i < len(vecs); i += 4 {
		cnt++
		search(1, 1, n, vecs[i], vecs[i+1], cnt, false)
		search(1, 1, n, vecs[i+2], vecs[i+3], cnt, true)
	}

	que := make([]int, (6*n+m+1)*2+10)
	dist := make([]int, 6*n+m+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = -2
	}

	front, end := len(que)/2, len(que)/2
	que[end] = xvec
	end++

	dist[xvec] = 0
	for front < end {
		u := que[front]
		front++
		for _, next := range adj[u] {
			v := next.first
			if dist[v] < 0 {
				dist[v] = dist[u] + next.second
				if next.second == 0 {
					front--
					que[front] = v
				} else {
					que[end] = v
					end++
				}
			}
		}
	}
	return dist[yvec] / 2
}

func compressAndSort(mp map[int]int) map[int]int {
	arr := make([]int, 0, len(mp))
	for k := range mp {
		arr = append(arr, k)
	}
	sort.Ints(arr)
	res := make(map[int]int)
	for j := 0; j < len(arr); j++ {
		res[arr[j]] = j + 1
	}
	return res
}

type Pair struct {
	first, second int
}

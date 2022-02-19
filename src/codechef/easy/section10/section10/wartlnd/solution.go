package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func main() {
	scanner := bufio.NewReader(os.Stdin)

	tc := readNum(scanner)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(scanner)

		roads := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			roads[i] = readNNums(scanner, 3)
		}

		buf.WriteString(fmt.Sprintf("%d\n", solve(n, roads)))
	}
	fmt.Print(buf.String())
}

const MAX = 100005

var mobius []int
var D [][]int

func init() {
	mobius = make([]int, MAX)
	mobius[1] = 1

	D = make([][]int, MAX)

	for i := 0; i < MAX; i++ {
		D[i] = make([]int, 0, 10)
	}

	for i := 1; i < MAX; i++ {
		D[i] = append(D[i], i)
		for j := i + i; j < MAX; j += i {
			mobius[j] -= mobius[i]
			D[j] = append(D[j], i)
		}
	}
	// fmt.Fprintf(os.Stderr, "%v\n", mobius)
}

type Pair struct {
	first  int
	second int
}

func solve(n int, roads [][]int) int64 {
	edge := make([][]Pair, MAX)

	for i := 0; i < MAX; i++ {
		edge[i] = make([]Pair, 0, 2)
	}

	for _, road := range roads {
		u, v, z := road[0], road[1], road[2]

		for _, d := range D[z] {
			edge[d] = append(edge[d], Pair{u, v})
		}
	}

	arr := make([]int, n+1)
	cnt := make([]int, n+1)

	for i := 0; i <= n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	var find func(x int) int

	find = func(x int) int {
		if arr[x] != x {
			arr[x] = find(arr[x])
		}
		return arr[x]
	}

	g := make([]int64, MAX)

	rollback := make([]int, 2*n+1)

	for i := 1; i < MAX; i++ {
		g[i] = 0
		var p int
		for _, ed := range edge[i] {
			u, v := ed.first, ed.second
			rollback[p] = u
			p++
			rollback[p] = v
			p++

			u = find(u)
			v = find(v)

			if u == v {
				continue
			}

			g[i] += int64(cnt[u]) * int64(cnt[v])

			if cnt[u] < cnt[v] {
				u, v = v, u
			}
			arr[u] = v
			cnt[v] += cnt[u]
		}

		for j := 0; j < p; j++ {
			x := rollback[j]
			arr[x] = x
			cnt[x] = 1
		}
	}

	f := make([]int64, MAX)

	for i := 1; i < MAX; i++ {
		for y := i; y < MAX; y += i {
			f[i] += int64(mobius[y/i]) * g[y]
		}
	}

	var res int64

	for i := 1; i < MAX; i++ {
		res += f[i] * int64(i)
	}

	return res
}

func solve1(n int, roads [][]int) int64 {
	outs := make([][]int, n)

	for i := 0; i < n; i++ {
		outs[i] = make([]int, 0, 10)
	}

	update := func(x, y, z int) {
		outs[x] = append(outs[x], y)
		outs[x] = append(outs[x], z)
	}

	for _, road := range roads {
		x, y, z := road[0]-1, road[1]-1, road[2]
		update(x, y, z)
		update(y, x, z)
	}

	var dfs func(p, u int) map[int]int

	var ans int64

	dfs = func(p, x int) map[int]int {
		aa := make(map[int]int)

		for i := 0; i < len(outs[x]); i += 2 {
			y := outs[x][i]
			z := outs[x][i+1]

			if y == p {
				continue
			}

			cc := dfs(x, y)

			bb := make(map[int]int)

			for b, d := range cc {
				x := gcd(b, z)
				bb[x] += d
			}

			bb[z]++

			for a, c := range aa {
				for b, d := range bb {
					e := gcd(a, b)
					ans += int64(e) * int64(c) * int64(d)
				}
			}

			for b, d := range bb {
				aa[b] += d
			}
		}

		for a, c := range aa {
			ans += int64(a) * int64(c)
		}

		return aa
	}

	dfs(-1, 0)

	return ans
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

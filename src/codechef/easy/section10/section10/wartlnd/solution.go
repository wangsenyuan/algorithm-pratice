package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n := readNum(scanner)

		roads := make([][]int, n-1)

		for i := 0; i < n-1; i++ {
			roads[i] = readNNums(scanner, 3)
		}

		fmt.Println(solve(n, roads))
	}
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
		edge[i] = make([]Pair, 0, 10)
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

	for i := 1; i < MAX; i++ {
		g[i] = 0
		rollback := make([]int, 0, 10)

		for _, ed := range edge[i] {
			u, v := ed.first, ed.second
			rollback = append(rollback, u)
			rollback = append(rollback, v)

			u = find(u)
			v = find(v)

			g[i] += int64(cnt[u]) * int64(cnt[v])
			arr[u] = v
			cnt[v] += cnt[u]
		}

		for _, x := range rollback {
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

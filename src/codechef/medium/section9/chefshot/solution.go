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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readAtMostNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, 0, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		if x == len(bs) {
			break
		}
		var y int

		x = readInt(bs, x, &y)

		res = append(res, y)
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		m := readNum(reader)
		C := make([][]int, m)
		for i := 0; i < m; i++ {
			C[i] = readAtMostNNums(reader, 17)
		}
		fmt.Println(solve(n, A, m, C))
	}
}

const MAX_N = 16
const MAX_S = 1 << MAX_N

var dist [][][]int

func init() {
	dist = make([][][]int, MAX_N)

	v := make([]int, 16)

	for i := 1; i < MAX_N; i++ {
		dist[i] = make([][]int, 1<<uint(i))

		for j := 0; j < 1<<uint(i); j++ {
			v = v[:0]
			for k := 0; k < i; k++ {
				if j&(1<<uint(k)) == 0 {
					v = append(v, k)
				}
			}
			s := len(v)

			dist[i][j] = make([]int, 1<<uint(s))

			for k := 0; k < s; k++ {
				for l := 0; l < 1<<uint(k); l++ {
					dist[i][j][l|1<<uint(k)] = dist[i][j][l] | (1 << uint(v[k]))
				}
			}
		}
	}
}

func solve(n int, A []int, M int, C [][]int) int {
	sort.Ints(A)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}

	subset := make([]Pair, M)
	for i := 0; i < M; i++ {
		var mask int
		for j := 1; j < len(C[i]); j++ {
			mask |= 1 << uint(C[i][j]-1)
		}
		subset[i] = Pair{C[i][0], mask}
	}
	sort.Sort(Pairs(subset))

	N := 1 << uint(n)

	ans := make([]int, N)
	sz := make([]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < n; j++ {
			if i&(1<<uint(j)) > 0 {
				sz[i] = sz[i-(1<<uint(j))] + 1
			}
		}
	}
	for i := 0; i < M; i++ {
		if i == 0 || subset[i] != subset[i-1] {
			for _, it := range dist[n][subset[i].second] {
				ans[it|subset[i].second] = max(ans[it|subset[i].second], ans[it]+A[sz[it]+subset[i].first-1])
			}
		}
	}
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	var save int
	for i := 0; i < N; i++ {
		save = max(save, ans[i])
	}
	return sum - save
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

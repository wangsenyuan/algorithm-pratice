package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 3)
	}
	fmt.Println(solve(n, m, k, E))
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

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(99))

}

var MOD = [...]uint64{1000000007, 1000000008, 1000000009}

type HashNum struct {
	val [3]uint64
}

func (this HashNum) Add(that HashNum) HashNum {
	var res [3]uint64
	res[0] = (this.val[0] + that.val[0]) % MOD[0]
	res[1] = (this.val[1] * that.val[1]) % MOD[1]
	res[2] = (this.val[2] + that.val[2]) % MOD[2]
	return HashNum{res}
}

func NewHashNum(v uint64) HashNum {
	var arr [3]uint64
	for i := 0; i < 3; i++ {
		arr[i] = v % MOD[i]
	}
	return HashNum{arr}
}

func solve(n, m, k int, E [][]int) int {
	hashValues := make([]HashNum, n+1)

	mem := make(map[uint64]bool)
	for i := 0; i <= n; i++ {
		tmp := r.Uint64()
		for mem[tmp] {
			tmp = r.Uint64()
		}
		hashValues[i] = NewHashNum(tmp)
		mem[tmp] = true
	}

	var tot HashNum

	for i := 1; i <= n; i++ {
		tot = tot.Add(hashValues[i])
	}

	degree := make([]int, n+1)
	for i := 0; i < m; i++ {
		e := E[i]
		degree[e[0]]++
	}
	G := make([][]Pair, n+1)
	for i := 0; i <= n; i++ {
		G[i] = make([]Pair, 0, degree[i])
	}
	for i := 0; i < m; i++ {
		e := E[i]
		u, v, w := e[0], e[1], e[2]
		G[u] = append(G[u], Pair{w, v})
	}

	C := make([][]HashNum, k+1)

	for i := 0; i <= k; i++ {
		C[i] = make([]HashNum, k+1)
	}

	for i := 1; i <= n; i++ {
		sort.Sort(Pairs(G[i]))
		d := degree[i]
		for j := 1; j <= d; j++ {
			v := G[i][j-1].second
			C[d][j] = C[d][j].Add(hashValues[v])
		}
	}

	var res int
	var dfs func(x int, cur HashNum)

	dfs = func(x int, cur HashNum) {
		if x == k {
			if cur == tot {
				res++
			}
			return
		}
		for i := 1; i <= x+1; i++ {
			dfs(x+1, cur.Add(C[x+1][i]))
		}
	}
	dfs(0, NewHashNum(0))

	return res
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

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
		n, k := readTwoNums(reader)
		C := readNNums(reader, n)
		E := make([][]int, n-1)
		for i := 0; i < n-1; i++ {
			E[i] = readNNums(reader, 2)
		}
		res := solve(n, k, C, E)
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

const MOD = 998244353

func pow(a, b int) int {
	A := int64(a)
	var R int64 = 1
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

func inv(x int) int {
	return pow(x, MOD-2)
}

func modMul(a int, b int) int {
	A := int64(a)
	B := int64(b)
	return int(A * B % MOD)
}

func modAdd(a int, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func solve(n, k int, C []int, E [][]int) int {
	set := CreateUFSet(n)

	for _, e := range E {
		u, v := e[0], e[1]
		u--
		v--
		if C[u] == 0 || C[v] == 0 {
			continue
		}
		set.Union(u, v)
	}

	cnts := make(map[int][]int)

	for i := 0; i < n; i++ {
		if C[i] == 0 {
			continue
		}
		p := set.Find(i)

		if _, ok := cnts[p]; !ok {
			cnts[p] = make([]int, k)
		}
		cnts[p][C[i]-1]++
	}

	K := 1 << uint(k)

	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, K)
	}

	dp[0][0] = 1
	var p int
	for _, v := range cnts {
		for i := 0; i < K; i++ {
			dp[p^1][i] = dp[p][i]
		}
		for i := 0; i < k; i++ {
			if v[i] == 0 {
				continue
			}
			mask := (K - 1) ^ (1 << uint(i))
			for j := mask; j >= 0; {
				dp[p^1][j|(1<<uint(i))] = modAdd(dp[p^1][j|(1<<uint(i))], modMul(v[i], dp[p][j]))
				if j == 0 {
					break
				}
				j = (j - 1) & mask
			}
		}
		p ^= 1
	}

	return dp[p][K-1]
}

type UFSet struct {
	set  []int
	cnt  []int
	size int
}

func CreateUFSet(n int) UFSet {
	set := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = 1
		set[i] = i
	}
	return UFSet{set, cnt, n}
}

func (uf *UFSet) Find(a int) int {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}
	return find(a)
}

func (uf *UFSet) Union(a, b int) {
	var find func(u int) int
	find = func(u int) int {
		if uf.set[u] != u {
			uf.set[u] = find(uf.set[u])
		}
		return uf.set[u]
	}

	pa := find(a)
	pb := find(b)
	if pa == pb {
		return
	}
	if uf.cnt[pa] < uf.cnt[pb] {
		uf.set[pa] = pb
		uf.cnt[pb] += uf.cnt[pa]
	} else {
		uf.set[pb] = pa
		uf.cnt[pa] += uf.cnt[pb]
	}
	uf.size--
}

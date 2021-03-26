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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(k, A, B)
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

const MAX_N = 100010
const MOD = 1000000007

var fact [MAX_N]int64
var bit [MAX_N]int

func init() {
	fact[0] = 1
	for i := 1; i < MAX_N; i++ {
		fact[i] = int64(i) * fact[i-1]
		fact[i] %= MOD
	}
}

func solve(k int, P []int, Q []int) int {
	n := len(P)

	if k == n {
		for i := 0; i < n; i++ {
			if Q[i] == P[0] {

				for j := 0; j < n; j++ {
					if P[j] != Q[(j+i)%n] {
						return -1
					}
				}

				return Q[0]
			}
		}
	}
	// k < n
	if k%2 == 0 {
		id := index(Q, MOD)
		return (id + 1) % MOD
	}
	if parity(P) != parity(Q) {
		return -1
	}

	id := index(Q, MOD)

	if index(Q, 2) == 1 {
		id = (id + MOD - 1) % MOD
	}

	id = int(int64(id) * int64(pow(2, MOD-2)) % MOD)

	return (id + 1) % MOD
}

func reset(n int) {
	for i := 0; i < MAX_N; i++ {
		bit[i] = 0
	}
}

func add(v int, d int) {
	for v < MAX_N {
		bit[v] += d
		v += v & -v
	}
}

func get(v int) int {
	var res int
	for v > 0 {
		res += bit[v]
		v -= v & -v
	}
	return res
}

func index(p []int, mod int64) int {
	n := len(p)

	reset(n)

	for i := 1; i <= n; i++ {
		add(i, 1)
	}

	var ans int64

	for i := 0; i < n; i++ {
		id := get(p[i] - 1)
		ans = (ans + int64(id)*fact[n-1-i]) % mod
		add(p[i], -1)
	}

	return int(ans)
}

func parity(p []int) int {
	n := len(p)
	reset(n)

	var ans int

	for i := n - 1; i >= 0; i-- {
		lt := get(p[i] - 1)
		ans = (ans + lt) % 2
		add(p[i], 1)
	}
	return ans
}

func pow(a, b int) int {
	A := int64(a)
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return int(R)
}

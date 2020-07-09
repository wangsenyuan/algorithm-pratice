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
		n, p := readTwoNums(reader)
		K := readNNums(reader, n)
		fmt.Println(solve(n, p, K))
	}
}

const MOD = 1000000007
const INF = 1000010

func solve(n, p int, K []int) int {
	sort.Ints(K)
	E := make([]int, n)
	copy(E, K)

	var cur int64
	var res int64
	var infy bool
	var prevExp = 1000010

	var i, j = n - 1, n - 1

	for i >= 0 {
		k := E[i]
		i--

		delta := prevExp - k
		prevExp = k

		res = res * pow(p, delta) % MOD

		for x := 0; x < delta && x < 20; x++ {
			cur *= int64(p)
			if cur > INF {
				infy = true
			}
		}

		for j >= 0 && K[j] == k {
			j--
			if infy || cur > 0 {
				cur--
				res += MOD - 1
			} else {
				cur++
				res++
			}
			res %= MOD
		}
	}

	res = res * pow(p, prevExp) % MOD

	return int(res)
}

func pow(a, b int) int64 {
	A := int64(a)
	var R int64 = 1

	for b > 0 {
		if b&1 == 1 {
			R *= A
			R %= MOD
		}
		A = A * A % MOD
		b >>= 1
	}
	return R
}

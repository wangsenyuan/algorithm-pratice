package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	fmt.Println(solve(n, A))
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

const MAX_N = 100010
const MOD = 1000000007

var M [MAX_N]int
var m [MAX_N]int

const INF = 1 << 30
const N_INF = -INF

var prod [MAX_N]int64

func solve2(n int, A []int) int {

	prod[0] = 1
	var res int64 = 1

	for j := 0; j < n; j++ {
		M[j] = N_INF
		m[j] = INF
		cur := A[j]
		k1 := j
		for k1 >= 0 && M[k1] < cur {
			M[k1] = cur
			k1--
		}
		k2 := j
		for k2 >= 0 && m[k2] > cur {
			m[k2] = cur
			k2--
		}
		k := min(k1, k2)
		if k < 0 {
			k = 0
		}
		for i := k; i < j; i++ {
			prod[i+1] = (prod[i] * int64(M[i]-m[i])) % MOD
		}

		res = (res * prod[j]) % MOD
	}
	return int(res)
}

func solve(n int, A []int) int {

	prod[0] = 1
	var res int64 = 1
	M[0] = INF
	m[0] = N_INF
	M[1] = A[0]
	m[1] = A[0]
	for j := 2; j <= n; j++ {
		cur := A[j-1]
		M[j] = cur
		m[j] = cur
		k := j - 1
		for M[k] < cur {
			M[k] = cur
			k--
		}
		for m[k] > cur {
			m[k] = cur
			k--
		}
		k++
		for i := k; i < j; i++ {
			prod[i] = (prod[i-1] * int64(M[i]-m[i])) % MOD
		}

		res = (res * prod[j-1]) % MOD
	}

	return int(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

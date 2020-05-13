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
		n := readNum(reader)
		A := readNNums(reader, n)

		fmt.Println(solve(n, A))
	}
}

const MOD = 998244353
const MAX_N = 100001

var F []int64

func init() {
	F = make([]int64, MAX_N)

	F[0] = 1

	for i := 1; i < MAX_N; i++ {
		F[i] = (F[i-1] * 2) % MOD
	}
}

func solve(n int, A []int) int {
	sort.Ints(A)

	var res int64

	var i int

	x := int64(1)
	cur := 1
	for {
		j := i

		for j < n && A[j] == cur {
			j++
		}

		tmp := int64(cur) * x
		tmp %= MOD
		tmp *= F[n-j]
		tmp %= MOD

		res += tmp
		res %= MOD

		if j == i {
			break
		}
		x *= (F[j-i] - 1)
		x %= MOD
		i = j

		cur++
	}

	return int(res)
}

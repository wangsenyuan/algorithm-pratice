package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		var N, M uint64
		s, _ := reader.ReadBytes('\n')
		pos := readUint64(s, 0, &N)
		readUint64(s, pos+1, &M)
		cnt, ways := solve(int64(N), int64(M))

		fmt.Printf("%d %d\n", cnt, ways)
	}
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

func solve(N, M int64) (int64, int64) {
	var pw int64 = 1

	var last = N
	var max int64
	var ways int64 = 1
	for size := 1; pw <= N; size++ {
		nxt := pw
		if pw <= N/M {
			nxt *= M
		} else {
			nxt = N + 1
		}

		hi := N / nxt

		cnt := (last - hi - (last/M - hi/M))

		max += cnt * int64((size+1)/2)

		if size%2 == 0 {
			ways = (ways * pow(int64(size/2+1), cnt)) % MOD
		}
		last = hi
		pw = nxt
	}

	return max, ways
}

func pow(A, b int64) int64 {
	R := int64(1)

	for b > 0 {
		if b&1 == 1 {
			R *= A
			R %= MOD
		}
		A = (A * A) % MOD
		b >>= 1
	}
	return R
}

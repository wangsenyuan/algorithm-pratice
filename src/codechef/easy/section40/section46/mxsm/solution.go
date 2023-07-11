package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		L := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var m int
			pos := readInt(s, 0, &m)
			L[i] = make([]int, m)
			for j := 0; j < m; j++ {
				pos = readInt(s, pos+1, &L[i][j])
			}
		}
		sum, cnt := solve(L)
		buf.WriteString(fmt.Sprintf("%d %d\n", sum, cnt))
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	if len(bs) == 0 || bs[0] == '\n' {
		return readNum(reader)
	}
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

const MOD = 1_000_000_007
const inf = 1 << 30

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}

func mul(a, b int) int {
	return int(int64(a) * int64(b) % MOD)
}

func solve(L [][]int) (int64, int) {
	n := len(L)
	lh := make([][]int, n)
	sum := make([][]int, n)
	znz := make([][]int, n)
	for i := 0; i < n; i++ {
		lh[i], sum[i], znz[i] = calcLowHigh(L[i])
	}

	var max_sum int64 = math.MinInt64
	var max_cnt int

	N := 1 << n

	for state := 0; state < N; state++ {
		var tmp int64 = 1
		var cnt = 1
		for i := 0; i < n; i++ {
			bit := (state >> i) & 1
			tmp *= int64(lh[i][bit])
			cnt = mul(cnt, sum[i][bit])
		}
		if tmp > max_sum {
			max_sum = tmp
			max_cnt = cnt
		} else if tmp == max_sum {
			max_cnt = add(max_cnt, cnt)
		}
	}

	if max_sum != 0 {
		return max_sum, max_cnt
	}

	max_cnt = 0

	for state := 1; state < N; state++ {
		var cnt int = 1
		for i := 0; i < n; i++ {
			bit := ((state >> i) & 1) ^ 1
			cnt = mul(cnt, znz[i][bit])
		}
		max_cnt = add(max_cnt, cnt)
	}

	return 0, max_cnt
}

func calcLowHigh(v []int) (lh []int, sum []int, znz []int) {
	n := len(v)
	lh = make([]int, 2)
	lh[0] = inf
	lh[1] = -inf
	sum = make([]int, 2)
	znz = make([]int, 2)

	for i := 0; i < n; i++ {
		var cur int
		for j := i; j < n; j++ {
			cur += v[j]
			if cur == 0 {
				znz[0]++
			}
			if cur > lh[1] {
				lh[1] = cur
				sum[1] = 1
			} else if cur == lh[1] {
				sum[1]++
			}

			if cur < lh[0] {
				lh[0] = cur
				sum[0] = 1
			} else if cur == lh[0] {
				sum[0]++
			}
		}
	}

	if lh[0] == lh[1] {
		sum[0] = 0
	}

	znz[1] = n*(n+1)/2 - znz[0]
	return
}

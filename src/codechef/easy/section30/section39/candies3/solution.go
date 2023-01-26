package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readNNums(reader, n)
		C := readNNums(reader, m)
		res := solve(A, C)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

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

func solve(A []int, C []int) int64 {
	//n := len(A)
	M := len(C)
	// n <= 1e5, m < 1e5
	// C[i] < 1e6
	// A[i] < m
	// for a given j sum(A[i] / j) * C[j]
	sort.Ints(A)

	freq := make([]int, M+1)
	var sum int64
	for _, a := range A {
		freq[a]++
		sum += int64(a)
	}

	for i := 1; i <= M; i++ {
		freq[i] += freq[i-1]
	}

	best := sum * int64(C[0])

	for price := 2; price <= M; price++ {
		var cnt int64
		for cur := price; cur <= M; cur += price {
			next := cur + price
			cnt += int64(freq[min(next-1, M)]-freq[cur-1]) * int64(cur/price)
		}
		best = max(best, cnt*int64(C[price-1]))
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(A []int, C []int) int64 {
	n := len(A)
	M := len(C)
	// n <= 1e5, m < 1e5
	// C[i] < 1e6
	// A[i] < m
	// for a given j sum(A[i] / j) * C[j]
	sort.Ints(A)

	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(A[i])
	}

	best := sum[n] * int64(C[0])

	for price := 2; price <= M; price++ {
		var cnt int64
		i := search(n, func(i int) bool {
			return A[i] >= price
		})
		for cur := price; cur <= M && i < n; cur += price {
			next := cur + price
			j := search(n, func(j int) bool {
				return A[j] >= next
			})
			cnt += int64(j-i) * int64(cur/price)
			i = j
		}
		best = max(best, cnt*int64(C[price-1]))
	}

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

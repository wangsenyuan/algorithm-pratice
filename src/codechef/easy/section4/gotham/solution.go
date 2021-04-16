package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	P := readNNums(reader, n)
	m := readNum(reader)
	Q := make([][]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNNums(reader, 2)
	}
	res := solve(n, P, Q)
	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d\n", res[i]))
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

const INF = 1000000000

func solve(n int, P []int, Q [][]int) []int64 {
	X := make([]int, n)
	arr := make([]int, 2*n)
	for i := n; i < 2*n; i++ {
		arr[i] = i - n
	}
	for i := n - 1; i > 0; i-- {
		arr[i] = min(arr[2*i], arr[2*i+1])
	}
	update := func(p, v int) {
		p += n
		X[p-n] += v
		if X[p-n] < P[p-n] {
			return
		}
		arr[p] = INF
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l, r int) int {
		l += n
		r += n

		res := INF

		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	res := make([]int64, len(Q))

	for i, cur := range Q {
		x, k := cur[0], cur[1]
		x--
		var tmp int64
		for k > 0 {
			p := get(x, n)
			// p is the first place at right side of x, that not full yet
			if p == INF {
				// out of reach
				break
			}
			if k >= P[p]-X[p] {
				tmp += int64(p-x) * int64(P[p]-X[p])
				k -= P[p] - X[p]
				update(p, P[p]-X[p])
			} else {
				// k < P[p] - X[p]
				tmp += int64(p-x) * int64(k)
				update(p, k)
				k = 0
			}
		}
		res[i] = tmp
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

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
	n, h := readTwoNums(reader)
	A := readNNums(reader, n)
	best, res := solve(A, h)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", best))

	for i := 0; i < n; i++ {
		if i+1 < n {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		} else {
			buf.WriteString(fmt.Sprintf("%d", res[i]))
		}
	}
	fmt.Println(buf.String())
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

func solve(A []int, h int) (int, []int) {
	sort.Ints(A)
	// 在同一个partition中
	// max(f(p)) = a[n] + a[n-1]
	// min(f(p)) = a[0] + a[1]
	// 在不同partition中间
	// max(f) = a[n-1] + b[m-1] + h
	// min(f) = a[0] + b[0] + h
	// 需要拿到最小值 max(f) - min(f)
	n := len(A)

	if n == 2 {
		return 0, []int{1, 2}
	}

	type Pair struct {
		first  int
		second int
	}

	B := make([]Pair, n)
	for i := 0; i < n; i++ {
		B[i] = Pair{A[i], i}
	}

	sort.Slice(B, func(i, j int) bool {
		return B[i].first < B[j].first
	})

	// n > 2
	best := (B[n-1].first + B[n-2].first) - (B[0].first + B[1].first)

	pos := -1

	for i := 0; i+1 < n; i++ {
		mn := B[0].first + B[i+1].first + h
		if i >= 1 {
			mn = min(mn, B[0].first+B[1].first)
		}
		if i+1 < n-1 {
			mn = min(mn, B[i+1].first+B[i+2].first)
		}
		mx := B[n-1].first + B[i].first + h
		if i+1 < n-1 {
			mx = max(mx, B[n-1].first+B[n-2].first)
		}
		if mx-mn < best {
			best = mx - mn
			pos = i
		}
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		if i <= pos {
			res[B[i].second] = 1
		} else {
			res[B[i].second] = 2
		}
	}
	return best, res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

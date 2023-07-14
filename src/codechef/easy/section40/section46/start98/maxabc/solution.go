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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(A, B []int) int64 {
	// a[i], b[i] 占据为止 (2 * i, 2 * i - 1)
	// 那在 2 * i - 1， 前面的部分，也就是 i-1, 让bob来安排的话，是固定的
	// 这个可以用dp计算出来
	// 后面的也可以计算出来
	// 然后剩下的是中间包含i的部分
	// 所以还需要一个前缀和
	n := len(A)
	till := make([]int64, n+2)
	from := make([]int64, n+2)
	pref := make([]int64, n+2)
	suf := make([]int64, n+2)
	for i := 1; i <= n; i++ {
		a, b := int64(A[i-1]), int64(B[i-1])
		pref[i] = max(pref[i-1], till[i-1]+max(a, b, a+b))
		till[i] = max(0, till[i-1]+a+b, a, b)
	}
	for i := n; i > 0; i-- {
		a, b := int64(A[i-1]), int64(B[i-1])
		suf[i] = max(suf[i+1], from[i+1]+max(a, b, a+b))
		from[i] = max(0, from[i+1]+a+b, a, b)
	}

	var best int64 = math.MaxInt64

	for i := 1; i <= n; i++ {
		a, b := int64(A[i-1]), int64(B[i-1])
		cur := max(pref[i-1], suf[i+1])
		ab := max(till[i-1]+a, from[i+1]+b, till[i-1]+from[i+1]+a+b)
		ba := max(till[i-1]+b, from[i+1]+a, till[i-1]+from[i+1]+a+b)
		cur = max(cur, min(ab, ba))
		best = min(best, cur)
	}

	return best
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}

func min(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res > arr[i] {
			res = arr[i]
		}
	}
	return res
}

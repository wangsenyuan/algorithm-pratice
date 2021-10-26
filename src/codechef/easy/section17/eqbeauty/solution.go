package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

const INF = math.MaxInt64

func solve(A []int) int64 {
	n := len(A)

	sort.Ints(A)

	check := func(l, r int, x int64) int64 {
		j := search(r-l, func(i int) bool {
			return int64(A[i+l]) > x
		})
		j += l
		var res int64 = INF
		if j < r {
			res = abs(int64(A[j]) - x)
		}
		j--
		// A[j] <= x
		if j >= l {
			res = min(res, abs(int64(A[j])-x))
		}
		return res
	}

	var ans int64 = INF
	// 0 < i < j < n z
	// (A[0], A[j]), (A[i], A[n-1])
	for i := 1; i+1 < n; i++ {
		// (A[n-1] - A[j]) - (A[i] - A[0])
		cost := int64(A[n-1]) + int64(A[0]) - int64(A[i])
		ans = min(ans, check(i+1, n-1, cost))
	}

	// (A[0], A[n-1]), (A[i], A[j])
	for i := 1; i+1 < n; i++ {
		// A[n-1] - A[0] - (A[j] - A[i])
		cost := int64(A[n-1]) - int64(A[0]) + int64(A[i])
		ans = min(ans, check(i+1, n-1, cost))
	}
	// (A[0], A[i]), (A[i+1], A[n])
	for i := 1; i+2 < n; i++ {
		// A[n] - A[i+1] - (A[i])
		ans = min(ans, abs((int64(A[i])-int64(A[0]))-(int64(A[n-1])-int64(A[i+1]))))
	}

	// if pick A[i] as B1 only
	var jp1, jp2 int64
	for i := 0; i < n; i++ {
		jp1 += abs(int64(A[(n-1)/2]) - int64(A[i]))
		jp2 += abs(int64(A[(n-1)/2+1]) - int64(A[i]))
	}

	for i := 0; i < n; i++ {
		var tmp int64

		if n&1 == 1 {
			//odd
			if i == (n-1)/2 {
				tmp = jp2 - abs(int64(A[(n-1)/2+1])-int64(A[i]))
			} else {
				tmp = jp1 - abs(int64(A[(n-1)/2])-int64(A[i]))
			}

		} else {
			if i <= (n-1)/2 {
				tmp = jp2 - abs(int64(A[(n-1)/2+1])-int64(A[i]))
			} else {
				tmp = jp1 - abs(int64(A[(n-1)/2])-int64(A[i]))
			}
		}
		if tmp < ans {
			ans = tmp
		}
	}

	return ans
}

func search(n int, fn func(int) bool) int {
	var l, r = 0, n
	for l < r {
		mid := (l + r) / 2
		if fn(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

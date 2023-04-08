package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		P := readNNums(reader, n)
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 4)
		}
		res := solve(P, Q)
		for _, ans := range res {
			buf.WriteString(fmt.Sprintf("%d\n", ans))
		}
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

func solve(P []int, Q [][]int) []int64 {
	n := len(P)
	mn := make([]int, n)
	mx := make([]int, n)
	for i, x := range P {
		mn[x] = i
		mx[x] = i
	}
	for i := 1; i < n; i++ {
		mn[i] = min(mn[i], mn[i-1])
		mx[i] = max(mx[i], mx[i-1])
	}

	calc := func(M int, L int) int64 {
		//number of subarrays with mex >= M, length >= L
		if M > n || L > n {
			return 0
		}
		if M == 0 {
			return int64(n-L+1) * int64(n-L+2) / 2
		}
		lo, hi := mn[M-1], mx[M-1]
		ret := int64(max(0, min(lo+1, hi-L+2))) * int64(n-hi)
		if hi+1 < n && hi-L+1 < lo {
			x := min(n-L+1, n-hi-1)
			y := 0
			if lo+L < n {
				y = n - lo - L + 1
			}
			// y + y+1 + ... + x
			ret += int64(x)*int64(x+1)/2 - int64(y)*int64(y-1)/2
		}
		return ret
	}

	ans := make([]int64, len(Q))

	for i, cur := range Q {
		l, r, x, y := cur[0], cur[1], cur[2], cur[3]
		ans[i] = calc(x, l) - calc(x, r+1) - calc(y+1, l) + calc(y+1, r+1)
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	return a + b - min2(a, b)
}

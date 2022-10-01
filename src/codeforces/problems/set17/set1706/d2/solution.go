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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

const N = 100010

var closest [N]int

const INF = math.MaxInt64

func solve(A []int, k int) int {
	ps := make([]int, N)

	n := len(A)

	for i := 0; i < n; i++ {
		pv := 1 << 30
		for j := 1; j <= min2(A[i], k); j = (A[i] / (A[i] / j)) + 1 {
			nv := A[i] / j
			ps[nv+1] = max(ps[nv+1], pv)
			pv = nv
		}
		ps[0] = max(ps[0], pv)
	}

	ans := 1 << 30
	var cm int
	for i := 0; i <= A[0]; i++ {
		cm = max(cm, ps[i])
		ans = min2(ans, cm-i)
	}

	return ans
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve1(A []int, k int) int {
	n := len(A)
	if n == 1 {
		return 0
	}

	arr := make([]int, n+1)
	copy(arr[1:], A)

	arr[0] = -1

	fill := func(s int, e int, v int) {
		for i := s; i < e; i++ {
			closest[i] = v
		}
	}

	for i := 1; i <= n; i++ {
		fill(arr[i-1]+1, arr[i]+1, arr[i])
	}

	fill(arr[n]+1, N, INF)

	v := make([]Pair, 2*k)

	var end int

	push := func(x Pair) {
		v[end] = x
		end++
	}

	back := func() Pair {
		return v[end-1]
	}

	var ans int64 = INF

	for mx := 1; mx <= 100000; mx++ {
		end = 0
		push(Pair{1, int64(mx)})
		for i := 2; i <= k; i++ {
			push(Pair{back().second + 1, int64(mx+1)*int64(i) - 1})
			if back().second > int64(arr[n]) {
				break
			}
		}
		if back().second >= int64(arr[n]) {
			var mn int64 = INF
			for i := 0; i < end; i++ {
				nxt := int64(closest[v[i].first])
				if nxt <= v[i].second {
					mn = min(mn, nxt/int64(i+1))
				}
			}
			ans = min(ans, int64(mx)-mn)
		}
	}

	return int(ans)
}

type Pair struct {
	first  int64
	second int64
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int64) int64 {
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

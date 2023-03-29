package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		B := readNNums(reader, n)
		C := readNNums(reader, n)
		res := solve(k, B, C)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const N = 1010

var dp [N]int

func init() {
	dp[1] = 0

	for i := 2; i < N; i++ {
		dp[i] = N
	}

	for i := 1; i < N; i++ {
		for x := 1; x <= i; x++ {
			j := i + i/x
			if j >= N {
				continue
			}
			if j == i {
				break
			}
			// j > i
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func solve(k int, B, C []int) int {
	// at most k operations
	// A[i] + A[i]/ x = B[i]
	// 先计算处B[i] 需要的步骤数
	n := len(B)
	var res int
	p := make([]Pair, 0, n)
	for i := 0; i < n; i++ {
		if B[i] == 1 {
			res += C[i]
			continue
		}
		if dp[B[i]] > k {
			continue
		}
		p = append(p, Pair{dp[B[i]], C[i]})
	}
	fp := make([]int, k+1)

	if k > 12*len(p) {
		k = 12 * len(p)
	}

	for i := 0; i < len(p); i++ {
		for j := k; j >= p[i].first; j-- {
			fp[j] = max(fp[j], fp[j-p[i].first]+p[i].second)
		}
		for j := 1; j <= k; j++ {
			fp[j] = max(fp[j], fp[j-1])
		}
	}

	return fp[k] + res
}

type Pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

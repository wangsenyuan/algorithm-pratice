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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

const N = 2*1e5 + 10

var L [N]int

func init() {
	L[1] = 1
	for i := 2; i < N; i++ {
		L[i] = L[i-1]
		if i-1 == L[i-1] {
			L[i] *= 2
		}
	}
}

func solve(A []int) int {
	n := len(A)
	if n == 1 {
		return 2
	}
	sort.Ints(A)
	// pick (x, y), count( < x) = pow(2, i)
	// count(< y) - count( < x) = pow(2, j)
	// N - count( < y) = pow(2, k)
	// 一个n * n的算法是可以正确给出答案的，但是太慢了
	// 某个i后面的都算做第三段 (A[i] > A[i-1]) 这个需要邀请的人数是确定的
	// 某个i前面的都算作第一段 (A[i] < A[i+1]) 这个也很容易计算,
	// dp[i] = a
	// fp[j] = min(dp[i] + h - (j - i)) h 是大于等于(j-i)的最小的pow(2, ?)
	//       =   = h1 - i + h2 - (j - i)
	//       =   = h1 + h2 - j (其实和i是没关系的)
	//   h1 + h2 最小
	// h2 依赖于i的选择
	// 假设 i从left到right移动,dp[i] = h1, dp[i+1] = h1 (偶然 dp[i+1] = 2 * h1 when i = h1)
	// 对于左边的i，有一个规律是 1, 2, 4, 4, 8, 8, 8，8 ..
	// 假设选定了h1, 那么i最好选择 dp[i] = h1, but dp[i+1] = 2 * h1的那个位置
	best := L[n] + 2
	pos := make([]int, 2*n)

	for y := 1; y < n; y++ {
		if A[y-1] == A[y] {
			// not able to split at position y
			continue
		}
		tmp := 1 + L[y]

		for h1 := 1; h1 < L[y]; h1 *= 2 {
			i := pos[h1]
			if i == 0 {
				continue
			}
			tmp = min(tmp, h1+L[y-i])
		}
		pos[L[y]] = y

		tmp += L[n-y]

		if best > tmp {
			best = tmp
		}
	}

	return best - n
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

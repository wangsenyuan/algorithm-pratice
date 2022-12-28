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

	t := readNum(reader)

	for t > 0 {
		t--
		m := readNum(reader)
		A := make([][]int, 2)
		for i := 0; i < 2; i++ {
			A[i] = readNNums(reader, m)
		}
		res := solve(m, A)
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

const INF = 1 << 50

func solve(m int, A [][]int) int64 {
	// 只有一个策略，就是在某一列 i(之前),
	// robot 按照 下 -> 右 -> 上 -> 右 的顺序访问
	// 到达i后，按照 右 -> ... n -> 换行  -> 左 -> i
	suf := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		suf[i] = make([]int64, m+1)
		for j := 0; j <= m; j++ {
			suf[i][j] = -INF
		}
	}

	for i := 0; i < 2; i++ {
		for j := m - 1; j >= 0; j-- {
			suf[i][j] = max3(suf[i][j+1]-1, int64(A[i][j]), int64(A[i^1][j]-2*(m-j)+1))
		}
	}

	pr := int64(A[0][0] - 1)
	var ans int64 = INF

	for j := 0; j < m; j++ {
		jj := j & 1
		ans = min2(ans, max3(pr, suf[jj][j+1]-2*int64(j)-1, int64(A[jj^1][j])-2*int64(m)+1))
		pr = max2(pr, int64(A[jj^1][j])-int64(2*j)-1)

		ans = min2(ans, max2(pr, suf[jj^1][j+1]-int64(2*j)-2))

		if j < m-1 {
			pr = max2(pr, int64(A[jj^1][j+1]-2*j-2))
		}
	}

	return ans + 2*int64(m)
}

func max3(a, b, c int64) int64 {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func max2(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min2(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

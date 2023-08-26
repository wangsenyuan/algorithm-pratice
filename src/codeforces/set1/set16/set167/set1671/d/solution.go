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
		n, x := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, x)

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

func solve(A []int, x int) int64 {
	// sum abs(a[i] - a[i+1])
	// 考虑这些数， 如果a[i] >= a[i+1], 贡献是 a[i], a[i+1] 贡献就是 -a[i+1]
	// 如果把c插入到 a 和 b的中间，并且它 c < max(a, b) > min(a, b), 那么它的贡献就是0
	// 最后组成的数组，是一个z形的
	// 谷贡献负值，峰贡献正值
	// 只需要考虑 1，x即可，
	// 1肯定会贡献负值，最好出现两次，x如果能贡献负值，（处在两个 >= x 的数中间）

	n := len(A)

	if n == 1 {
		if x < A[0] {
			return int64(A[0] - 1)
		}
		return int64(x - 1)
	}
	var cur int64

	for i := 0; i+1 < n; i++ {
		cur += int64(abs(A[i] - A[i+1]))
	}

	var ans int64 = 1 << 60

	for times := 0; times <= 1; times++ {
		best := int64(A[0] - 1)
		ans = min(ans, cur+int64(abs(A[0]-x))+int64(x-1))
		for i := 0; i+1 < n; i++ {
			// put x between i & i + 1
			ans = min(ans, cur+best-int64(abs(A[i]-A[i+1]))+int64(abs(A[i]-x))+int64(abs(A[i+1]-x)))
			// put x & 1, between i & i + 1
			ans = min(ans, cur-int64(abs(A[i]-A[i+1]))+int64(abs(A[i]-x))+int64(A[i+1]-1)+int64(x-1))

			best = min(best, -int64(A[i]-A[i+1])+int64(A[i]-1)+int64(A[i+1]-1))
		}

		ans = min(ans, cur+best+int64(abs(x-A[n-1])))

		reverse(A)
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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

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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int, k int) int64 {
	n := len(A)
	// x >= 1 && x <= 2 * n
	// 对于两个数字 x, y, 一开始不在A中，
	// 假设 x < y, 并且始终最终的结果中也使用了x，y
	// 且是先加入x，后加入y
	// 假设加入x时，最大值是 a
	//    加入y时，最大值是 b
	// 则现在的贡献是 a - x + b - y
	//  第一种情况   a <= x (a < y), 交换 x和y的顺序 max(a, y) - y + b - x = y - y + b - x = b - x > b - x + a - y
	//  第二种情况  a > x and a <= y, 交换x和y  max(a, y) - y + b - x = y - y + b - x = b - x >= b - x + a - y
	//  第三种情况  a > y 交换 x 和 y max(a, y) - y + b - x = a - y + b - x 至少没有变坏
	// 所以结论是，如果  x < y, 那么优先使用 y, 再使用 x
	// 然后考虑 k < n, 也就是有一部分不使用的情况，哪些不使用呢？
	// 考虑只剩一个数，且有x, 和 y可以选, x < y
	// 如果 x 和 y都大于现有的数字， 选x和y没有区别
	// 假设 x < a && y > a, 选x
	// x < a && y < a 选x
	// 那是不是选择最小的k个数字，并且倒序排就是结果呢？
	// 假设是这样子, 剩下的数字中，我们选择某个x， 显然如果x比选中的数都大，交换它不会更好
	// 假设x < a, 那么a只能是原有的最大值， 假设它替换y (y < x), a - x < a - y, 会使得结果更差
	vis := make([]bool, 2*n+1)

	var mx int

	for _, num := range A {
		vis[num] = true
		mx = max(mx, num)
	}

	B := make([]int, 0, n)

	for i := 1; i <= 2*n; i++ {
		if !vis[i] {
			B = append(B, i)
		}
	}

	sort.Ints(B)
	var res int64

	if mx > B[k-1] {
		for i := 0; i < k; i++ {
			res += int64(mx - B[i])
		}
	}
	// 如果使用一个比mx大的数字，那么这个数字最好是 2 * n , 且越早越好
	if mx < 2*n {
		mx = 2 * n
		var tmp int64
		for i := 0; i < k-1; i++ {
			tmp += int64(mx - B[i])
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

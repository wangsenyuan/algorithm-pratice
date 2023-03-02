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
		n := readNum(reader)
		res := solve(n)
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

func solve(k int) int64 {
	// F[n] == x
	// F[n] = perfect squars - perfect cubes
	// x * x <= n
	// y * y * y <= n
	// 对于给定的n，perfect squars = sqrt(n),
	// 对于给定的n, perfect cubes = cube(n)
	// 不是一个简单的关系， 假设x= 3
	// 当 n = 31是，
	// 平方数有5个
	// 但是立方数有3个  它们的差是2
	// 但是 n= 25 时，平方数是5， 立方数是2， 差是3
	// 可以确定的一点是 n 要么是一个平方数，要们是一个立方数
	// 假设它既不是平方数，也不是立方数，通过将它减小到最近的立方数（平方数）不会改变结果
	// 假设它是一个平方数
	//  存在两种情况有a个平方数 (n = a * a), 立方数的数量 b * b * b <= n
	//    如果 a - b = x 或者 a - b == x + 1
	a := search(1<<31, func(a int64) bool {
		// n == x * x
		N := a * a
		b := search(1<<21, func(b int64) bool {
			M := b * b * b
			return M > N
		})
		b--
		return a-b >= int64(k)
	})

	b := search(1<<21, func(b int64) bool {
		M := b * b * b
		a := search(1<<31, func(a int64) bool {
			N := int64(a) * int64(a)
			return N > M
		})
		a--
		return a-b >= int64(k)
	})

	res := min(a*a, b*b*b)

	return res
}

func search(n int64, f func(int64) bool) int64 {
	var l, r int64 = 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
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

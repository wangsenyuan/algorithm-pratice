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
		n := readNum(reader)
		p := readNNums(reader, n)
		res := solve(p)

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

const inf = 1 << 30

func solve(a []int) int {
	// 只能选择奇数i的时候，才会有分数
	// 选择偶数i的时候，删除
	// 考虑最后的结果，正数最大的能被选中越好
	// 也就是希望正数在奇数位置被选走
	// 选择一个负数，是否更有利呢？
	// 比如 -1,10
	// 直接无法取到10，但是先取-1， 在取10，显然是最优解
	// 如果前面取了偶数个数，对当前数是没有影响的
	// 是不是可以dp
	// 应该是不能dp的，因为处理的顺序是不确定的，
	// 从后面往前，
	//	如果i是奇数，且a[i] > 0, res += a[i]
	// 	同时 res += sum of positive 偶数
	n := len(a)
	var s2 int
	var best int
	var res int
	for i := n; i >= 1; i-- {
		if i&1 == 1 {
			// 或者通过使用a[i], 将所有的偶数变成奇数位
			best = max(best, s2+a[i-1]+res)
			// a odd position
			if a[i-1] > 0 {
				// better to take a[i]
				res += a[i-1]
				res += s2
				s2 = 0
			}
			// else skip it
		} else {
			// 先取它
			best = max(best, s2+res)
			// a even position
			s2 += max(0, a[i-1])
		}
	}

	return max(res, best)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

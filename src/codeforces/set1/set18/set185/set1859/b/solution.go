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

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := make([][]int, n)
		for i := 0; i < n; i++ {
			m := readNum(reader)
			a[i] = readNNums(reader, m)
		}
		res := solve(a)
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

func solve(a [][]int) int {
	// max the beauty
	// 尽量把最小值放在同一个队列里
	// 假设队列里只有两个元素 [x, y], [u, v]
	// x < y, and u < v
	// 如果将下移动到 [u, v] 中， y + min(x, u)
	// 如果 x > u, 那这个交换是good的
	// x < u, x + y, 只有当 y > u时，这个交换也是合理的
	// 首先找到拥有最小值的队列，这个无论怎么变化，它都不能更好
	// 然后把所有其他队列中的最小值，移动到这里，
	var it int
	for i, x := range a {
		sort.Ints(x)
		if x[0] < a[it][0] {
			it = i
		}
	}
	var jt int
	res := a[it][0]
	// 找到第2位里面最小的数
	for i := 0; i < len(a); i++ {
		//if i == it {
		//	continue
		//}
		if a[i][1] < a[jt][1] {
			jt = i
		}
		res += a[i][1]
	}
	// 把最小的那个减掉（把it里面的数移动过去)

	res -= a[jt][1]

	return res
}

type Pair struct {
	first  int
	second int
}

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
		a, b := readTwoNums(reader)
		res := solve(a, b)
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

const inf = 1 << 60

func solve(a int, b int) int {
	// 先假设跳到了a行1列
	// 这里不需要限制m（腿的长度），或者说不需要考虑腿的长度
	// 假设腿长为m时，跳过了 a，但因为腿长为+1长上去的
	// 完全可以在合适的时候，跳一次，再继续长，满足剩下的条件即可

	best := a + b
	for m := 1; m <= 100000; m++ {
		tmp := (a+m-1)/m + (b+m-1)/m + m - 1
		best = min(best, tmp)
	}

	return best
}

func solve1(a, b int) int {
	fa := factors(a)
	fb := factors(b)
	res := inf
	for _, x := range fa {
		tmp := x - 1 + a/x
		if x >= b {
			break
		}
		// x < b
		for _, y := range fb {
			if y < x {
				continue
			}
			// y >= x
			tmp2 := y - x + b/y
			res = min(res, tmp+tmp2)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func factors(num int) []int {
	var res []int
	for x := 1; x*x <= num; x++ {
		if num%x == 0 {
			res = append(res, x)
			y := num / x
			if y != x {
				res = append(res, y)
			}
		}
	}
	sort.Ints(res)
	return res
}

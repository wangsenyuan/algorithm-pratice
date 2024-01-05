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
		lambs := make([][]int, n)
		for i := 0; i < n; i++ {
			lambs[i] = readNNums(reader, 2)
		}
		res := solve(lambs)
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

func solve(lambs [][]int) int {
	// 假设最终turn on 了x个灯
	// 那么剩下的 n - x个灯的 a[i] <= x,
	// 按照b降序排
	// 假设我希望把前x都使用掉
	// 对于第i个来说，如果加入它必须淘汰一个，那么就跳过它
	// 但这个判断，并不是判断 a[i] <= len(set)
	// 而是看是否存在 a[j] = a[i] ?，
	// 那是不是a[j] = a[i] 的情况下，直接忽略掉它即可？除非 a[i] >= len(set)
	sort.Slice(lambs, func(i, j int) bool {
		a := lambs[i]
		b := lambs[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	n := len(lambs)

	use := make([]bool, n)
	var sum int

	var l, sz int
	for i := 0; i < n; {
		// lambs[i][0] > sz
		sum += lambs[i][1]
		use[i] = true
		i++
		sz++
		for i < n && lambs[i][0] <= sz {
			i++
		}
		for l <= i && l < n && sz > 0 {
			if use[l] {
				if lambs[l][0] > sz {
					break
				}
				sz--
			}
			l++
		}
	}

	return sum
}

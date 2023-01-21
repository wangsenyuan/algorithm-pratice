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
		n, m := readTwoNums(reader)
		infected := readNNums(reader, m)
		res := solve(n, m, infected)
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

func solve(n int, m int, infected []int) int {
	// 每个已感染的house，每天的贡献为1
	// 所以需要protect越多的house，应该是让传播的天数越低越好，而且应该是成对处理
	if m == 0 {
		return 0
	}

	if m == 1 {
		// use two days to protect
		return 2
	}

	var segs []Seg

	sort.Ints(infected)

	for i := 1; i < m; i++ {
		l := infected[i] - infected[i-1] - 1

		if l > 0 {
			segs = append(segs, Seg{infected[i-1] + 1, infected[i] - 1, l})
		}
	}

	segs = append(segs, Seg{infected[m-1] + 1, infected[0] - 1 + n, infected[0] - 1 + n - (infected[m-1] + 1) + 1})

	sort.Slice(segs, func(i, j int) bool {
		return segs[i].seg_len > segs[j].seg_len
	})

	var res int

	for i, day := 0, 0; i < len(segs); i++ {
		cur := segs[i]
		// 两头感染
		l := cur.left
		r := cur.right
		l += day
		r -= day
		if l < r {
			r--
			res += r - l + 1
			day += 2
		} else if l == r {
			res++
			day++
		} else {
			break
		}
	}

	return n - res
}

type Seg struct {
	left    int
	right   int
	seg_len int
}

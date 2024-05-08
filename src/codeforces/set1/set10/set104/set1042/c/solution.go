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

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	var buf bytes.Buffer
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

func solve(a []int) [][]int {
	var zeros []int
	var neg []int
	var pos []int
	for i, num := range a {
		if num == 0 {
			zeros = append(zeros, i)
		} else if num < 0 {
			neg = append(neg, i)
		} else {
			pos = append(pos, i)
		}
	}
	if len(neg)%2 == 1 {
		sort.Slice(neg, func(i, j int) bool {
			return a[neg[i]] > a[neg[j]]
		})
	}
	var res [][]int
	for i := len(neg) - 1; i > 0; i -= 2 {
		x := neg[i]
		y := neg[i-1]
		res = append(res, []int{1, x + 1, y + 1})
		pos = append(pos, y)
	}
	if len(neg)%2 == 1 {
		neg = neg[:1]
	} else {
		neg = nil
	}
	for i := len(zeros) - 1; i > 0; i-- {
		x := zeros[i]
		y := zeros[i-1]
		res = append(res, []int{1, x + 1, y + 1})
	}
	if len(zeros) > 0 {
		zeros = zeros[:1]
	}
	if len(neg) > 0 && len(zeros) > 0 {
		// 都是1
		x := neg[0]
		y := zeros[0]
		res = append(res, []int{1, x + 1, y + 1})
		if len(pos) > 0 {
			// 只有可能得到正的数时
			res = append(res, []int{2, y + 1})
		}
	} else if len(pos) > 0 {
		if len(neg) > 0 {
			res = append(res, []int{2, neg[0] + 1})
		} else if len(zeros) > 0 {
			res = append(res, []int{2, zeros[0] + 1})
		}
	}

	for i := len(pos) - 1; i > 0; i-- {
		x := pos[i]
		y := pos[i-1]
		res = append(res, []int{1, x + 1, y + 1})
	}

	return res
}

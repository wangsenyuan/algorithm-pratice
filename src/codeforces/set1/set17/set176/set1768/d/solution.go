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
		a := readNNums(reader, n)
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
func solve(p []int) int {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i]--
	}

	// 首先是形成cycle的那部分，它们需要ln - 1次操作， 放回原位
	// 如果 x, x + 1 在同一个环内，怎么样
	//  [3, 1, 2] 组成了一个个环
	// 用2次可以回到原位，但是只需要一次，只需要一次
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		pos[p[i]] = i
	}

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = -1
	}

	var ans int
	for i := 0; i < n; i++ {
		if id[i] >= 0 {
			continue
		}
		var cnt int
		j := i
		for id[j] < 0 {
			id[j] = i
			cnt++
			j = pos[j]
		}
		ans += cnt - 1
	}

	for i := 0; i+1 < n; i++ {
		if id[i] == id[i+1] {
			// in same cycle
			return ans - 1
		}
	}

	return ans + 1
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

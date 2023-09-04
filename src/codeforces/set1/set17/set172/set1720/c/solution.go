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
		n, m := readTwoNums(reader)
		mat := make([]string, n)
		for i := 0; i < n; i++ {
			mat[i] = readString(reader)[:m]
		}
		res := solve(mat)
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

func solve(mat []string) int {
	// 只需要考虑首次操作的最优性，只要出现两个00，后续都可以通过1次操作删除一个1进行
	var cnt int
	for _, row := range mat {
		for i := 0; i < len(row); i++ {
			if row[i] == '1' {
				cnt++
			}
		}
	}

	if cnt <= 1 {
		return cnt
	}

	n := len(mat)
	m := len(mat[0])
	ok := make([]bool, 5)
	for i := 0; i+1 < n; i++ {
		for j := 0; j+1 < m; j++ {
			var cur int
			for r := 0; r < 2; r++ {
				for c := 0; c < 2; c++ {
					if mat[i+r][j+c] == '1' {
						cur++
					}
				}
			}
			if cur > 0 {
				ok[max(cur-1, 1)] = true
			}
		}
	}
	if ok[1] {
		return cnt
	}
	if ok[2] {
		return cnt - 1
	}
	return cnt - 2
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

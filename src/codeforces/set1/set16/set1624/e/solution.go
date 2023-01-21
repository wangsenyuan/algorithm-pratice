package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		reader.ReadString('\n')
		n, m := readTwoNums(reader)
		A := make([]string, n)
		for i := 0; i < n; i++ {
			A[i] = readString(reader)
		}
		S := readString(reader)
		res := solve(n, m, A, S)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(fmt.Sprintf("%d\n", len(res)))
			for _, cur := range res {
				buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
			}
		}
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

func solve(n int, m int, A []string, S string) [][]int {
	// 貌似只有2位或者3位的字符串才需要记录
	marks := make([][]Mark, 4)

	// at most 999
	for i := 0; i < 4; i++ {
		marks[i] = make([]Mark, 1000)
	}

	set := func(key string, l, r, i int) {
		if r == l {
			// at least length 2
			return
		}
		x := r - l + 1

		y, _ := strconv.Atoi(key)

		marks[x][y] = Mark{l, r, i}
	}

	get := func(key string) (bool, Mark) {
		if len(key) < 2 {
			return false, Mark{}
		}
		x := len(key)
		y, _ := strconv.Atoi(key)
		cur := marks[x][y]
		if cur.r-cur.l+1 == len(key) {
			return true, cur
		}
		return false, cur
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for x := 2; x <= 3 && j+x <= m; x++ {
				set(A[i][j:j+x], j, j+x-1, i)
			}
		}
	}

	ans := make([]Mark, m)

	vis := make([]bool, m)

	var build func(i int) bool

	build = func(i int) bool {
		if i == m || ans[i].Len() > 1 {
			return true
		}
		if vis[i] {
			return false
		}
		vis[i] = true

		for x := 2; x <= 3 && i+x <= m; x++ {
			ok, mark := get(S[i : i+x])
			if ok {
				ans[i] = mark
				if build(i + mark.Len()) {
					return true
				}
				ans[i] = Mark{}
			}
		}
		return false
	}

	if build(0) {
		var res [][]int
		pos := 0
		for pos < m {
			cur := ans[pos]
			res = append(res, []int{cur.l + 1, cur.r + 1, cur.i + 1})
			pos += cur.Len()
		}
		return res
	}
	return nil
}

type Mark struct {
	l int
	r int
	i int
}

func (this Mark) Len() int {
	return this.r - this.l + 1
}

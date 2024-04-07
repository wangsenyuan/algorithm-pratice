package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ask := func(i int, j int) int {
		fmt.Printf("? %d %d\n", i, j)
		return readNum(reader)
	}

	n := readNum(reader)

	cnt, ans := solve(n, ask)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("!\n%d\n", cnt))

	for i := 0; i < len(ans); i++ {
		buf.WriteString(fmt.Sprintf("%d ", ans[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, ask func(int, int) int) (int, []int) {
	que := make([][]int32, n)
	for i := 0; i < n; i++ {
		que[i] = make([]int32, n)
		for j := 0; j < n; j++ {
			que[i][j] = -1
		}
	}

	query := func(i, j int) int {
		if que[i][j] >= 0 {
			return int(que[i][j])
		}
		que[i][j] = int32(ask(i, j))
		return int(que[i][j])
	}

	b := make([]int, n)

	test := func(p []int) bool {
		for i := 0; i < n; i++ {
			b[p[i]] = i
		}
		for i := 0; i < n; i++ {
			if query(0, i) != p[0]^b[i] {
				return false
			}
			if query(i, 0) != p[i]^b[0] {
				return false
			}
		}

		return true
	}

	var cnt int
	var ans []int

	vis := make([]bool, n)
	for b0 := 0; b0 < n; b0++ {
		for i := 0; i < n; i++ {
			vis[i] = false
		}
		ok := true
		p := make([]int, n)

		for i := 0; i < n; i++ {
			p[i] = query(i, 0) ^ b0
			if p[i] >= n || vis[p[i]] {
				ok = false
				break
			}
		}
		if ok && test(p) {
			cnt++
			if len(ans) == 0 {
				ans = p
			}
		}
	}
	return cnt, ans
}

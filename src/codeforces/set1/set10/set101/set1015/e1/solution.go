package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ok, res, _ := process(reader)

	if !ok {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
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

func process(reader *bufio.Reader) (ok bool, res [][]int, a []string) {
	n, m := readTwoNums(reader)
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)[:m]
	}
	ok, res = solve(a)
	return
}

func solve(a []string) (bool, [][]int) {

	var res [][]int

	n := len(a)
	m := len(a[0])
	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, m)
	}

	process := func(x int, y int) {
		d := 1
		for x-d >= 0 && x+d < n && y-d >= 0 && y+d < m {
			if a[x-d][y] == '*' && a[x+d][y] == '*' && a[x][y-d] == '*' && a[x][y+d] == '*' {
				cnt[x-d][y]++
				cnt[x+d][y]++
				cnt[x][y-d]++
				cnt[x][y+d]++
			} else {
				break
			}
			d++
		}
		if d > 1 {
			res = append(res, []int{x + 1, y + 1, d - 1})
			cnt[x][y]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '*' {
				process(i, j)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if a[i][j] == '*' && cnt[i][j] == 0 {
				return false, nil
			}
			if a[i][j] == '.' && cnt[i][j] > 0 {
				return false, nil
			}
		}
	}

	return true, res
}

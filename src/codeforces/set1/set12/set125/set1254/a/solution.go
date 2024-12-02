package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		res, _, _ := process(reader)
		for _, row := range res {
			buf.WriteString(row)
			buf.WriteByte('\n')
		}
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

func process(reader *bufio.Reader) (res []string, land []string, k int) {
	n, m, k := readThreeNums(reader)
	land = make([]string, n)
	for i := range n {
		land[i] = readString(reader)[:m]
	}
	res = solve(k, land)
	return
}

var letters string

func init() {
	var buf []byte
	for i := byte('0'); i <= byte('9'); i++ {
		buf = append(buf, i)
	}
	for i := byte('A'); i <= byte('Z'); i++ {
		buf = append(buf, i)
	}
	for i := byte('a'); i <= byte('z'); i++ {
		buf = append(buf, i)
	}
	letters = string(buf)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, land []string) []string {
	n := len(land)
	m := len(land[0])
	buf := make([][]byte, n)
	var cnt int
	for i := 0; i < n; i++ {
		buf[i] = []byte(land[i])
		for j := 0; j < m; j++ {
			if land[i][j] == 'R' {
				cnt++
			}
			buf[i][j] = '-'
		}
	}

	que := make([]int, 0, n*m)
	rect := []int{0, 0, n - 1, m - 1}
	for rect[0] <= rect[2] && rect[1] <= rect[3] {
		for j := rect[1]; j <= rect[3]; j++ {
			que = append(que, rect[0]*m+j)
		}
		rect[0]++
		if rect[0] > rect[2] {
			break
		}
		for i := rect[0]; i <= rect[2]; i++ {
			que = append(que, i*m+rect[3])
		}
		rect[3]--
		if rect[1] > rect[3] {
			break
		}
		for j := rect[3]; j >= rect[1]; j-- {
			que = append(que, rect[2]*m+j)
		}
		rect[2]--

		for i := rect[2]; i >= rect[0]; i-- {
			que = append(que, i*m+rect[1])
		}

		rect[1]++
	}

	x, y := cnt/k, cnt%k
	var j int
	for i := 0; i < k; i++ {
		z := x
		if i < y {
			z++
		}
		for j < len(que) && z > 0 {
			r, c := que[j]/m, que[j]%m
			buf[r][c] = letters[i]
			if land[r][c] == 'R' {
				z--
			}
			j++
		}
	}

	for j < len(que) {
		r, c := que[j]/m, que[j]%m
		buf[r][c] = letters[k-1]
		j++
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = string(buf[i])
	}

	return ans
}

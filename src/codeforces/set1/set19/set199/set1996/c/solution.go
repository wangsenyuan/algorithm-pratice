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
		res := process(reader)
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) []int {
	_, m := readTwoNums(reader)
	s := readString(reader)
	t := readString(reader)
	qs := make([][]int, m)
	for i := 0; i < m; i++ {
		qs[i] = readNNums(reader, 2)
	}
	return solve(s, t, qs)
}

func solve(s string, t string, qs [][]int) []int {
	n := len(s)

	pref := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		pref[i] = make([]int, 26)
	}

	for i := 0; i < n; i++ {
		copy(pref[i+1], pref[i])
		x := int(s[i] - 'a')
		y := int(t[i] - 'a')
		pref[i+1][x]++
		pref[i+1][y]--
	}

	ans := make([]int, len(qs))

	for i, cur := range qs {
		l, r := cur[0], cur[1]
		var sum int
		for j := 0; j < 26; j++ {
			sum += abs(pref[r][j] - pref[l-1][j])
		}
		ans[i] = sum / 2
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}

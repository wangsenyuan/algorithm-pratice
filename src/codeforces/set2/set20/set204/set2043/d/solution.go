package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for range tc {
		l, r, g := readThreeNums(reader)
		res := solve(l, r, g)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

func solve(l int, r int, g int) []int {
	if r < g {
		return []int{-1, -1}
	}
	var que [][]int
	a := (l + g - 1) / g * g
	b := r / g * g

	if a > b {
		return []int{-1, -1}
	}

	que = append(que, []int{a, b})

	var tail int

	for tail < len(que) {
		cur := que[tail]
		tail++
		a, b := cur[0], cur[1]
		if gcd(a, b) == g {
			return cur
		}
		if a+g <= b {
			que = append(que, []int{a, b - g})
			que = append(que, []int{a + g, b})
		}
	}

	return []int{-1, -1}
}
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

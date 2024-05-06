package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	cmds := make([][]int, n)
	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var x int
		pos := readInt(s, 0, &x)
		if x == 1 {
			cmds[i] = make([]int, 2)
			cmds[i][0] = 1
			readInt(s, pos+1, &cmds[i][1])
		} else {
			cmds[i] = []int{2}
		}
	}
	res := solve(cmds)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%.9f\n", x))
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(cmds [][]int) []float64 {
	// 假设在中间取到最大值吧
	// 不知道咋证明
	prev := []int{0}
	a := []int{0}
	var sum int

	var res []float64

	calc := func(c int) float64 {
		an := a[len(a)-1]
		return float64(an) - float64(prev[c]+an)/float64(c+1)
	}

	get := func() float64 {
		if len(a) == 1 {
			return 0
		}
		l, r := 1, len(a)-1
		for r-l > 3 {
			m1 := l + (r-l)/3
			m2 := (r + m1) / 2
			if calc(m1) > calc(m2) {
				r = m2
			} else {
				l = m1
			}
		}
		m1 := l + (r-l)/3
		m2 := (r + m1) / 2
		return max(calc(l), calc(r), calc(m1), calc(m2))
	}

	for _, cmd := range cmds {
		if cmd[0] == 2 {
			res = append(res, get())
		} else {
			a = append(a, cmd[1])
			sum += cmd[1]
			prev = append(prev, sum)
		}
	}

	return res
}

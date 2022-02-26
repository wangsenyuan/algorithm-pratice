package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	readNum(reader)

	query := func(n int, w []int) []int {
		var buf bytes.Buffer
		buf.WriteString("?")
		for i := 0; i < len(w); i++ {
			buf.WriteString(fmt.Sprintf(" %d", w[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		return readNNums(reader, n-1)
	}

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		res := solve(n, m, query)

		if len(res) == 0 {
			fmt.Println("! -1")
		} else {
			var buf bytes.Buffer
			buf.WriteString("!")
			for i := 0; i < m; i++ {
				buf.WriteString(fmt.Sprintf(" %d", res[i]))
			}
			fmt.Println(buf.String())
		}

		readNum(reader)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, m int, query func(int, []int) []int) []int {
	b := make([][]bool, m)
	for i := 0; i < m; i++ {
		b[i] = make([]bool, m)
	}
	w := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			w[j] = (i+j)%m + 1
		}
		q := query(n, w)
		for _, x := range q {
			b[x][i] = true
		}
	}
	var bridge []int
	p := make([]int, m)
	for i := 0; i < m; i++ {
		p[i] = -1
	}

	for i := 0; i < m; i++ {
		var l int
		for l < m && b[i][l] {
			l++
		}
		if l == m {
			bridge = append(bridge, i)
			continue
		}
		for !b[i][l] {
			l = (l + 1) % m
		}
		p[(m-l)%m] = i
	}

	if len(bridge) > 1 {
		return nil
	}

	for _, x := range bridge {
		var i int
		for p[i] != -1 {
			i++
		}
		p[i] = x
	}

	return p
}

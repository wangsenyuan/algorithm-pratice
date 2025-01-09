package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}

	var buf bytes.Buffer
	buf.WriteString("YES\n")
	for i := range res {
		for j := range res[i] {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (res [][]int, a []int, b []int) {
	n, m := readTwoNums(reader)
	a = readNNums(reader, n)
	b = readNNums(reader, m)
	res = solve(a, b)
	return
}

func solve(a []int, b []int) [][]int {
	n := len(a)
	m := len(b)

	res := make([][]int, n)
	for i := range n {
		res[i] = make([]int, m)
	}

	for d := 0; d < 30; d++ {
		var x []int
		for i := range n {
			if (a[i]>>d)&1 == 1 {
				x = append(x, i)
			}
		}

		var y []int
		for i := range m {
			if (b[i]>>d)&1 == 1 {
				y = append(y, i)
			}
		}
		var i, j int
		for i < len(x) && j < len(y) {
			res[x[i]][y[j]] |= 1 << d
			i++
			j++
		}
		if (len(x)-i)%2 == 1 || (len(y)-j)%2 == 1 {
			return nil
		}
		for i < len(x) {
			res[x[i]][0] |= 1 << d
			i++
			res[x[i]][0] |= 1 << d
			i++
		}
		for j < len(y) {
			res[0][y[j]] |= 1 << d
			j++
			res[0][y[j]] |= 1 << d
			j++
		}
	}

	return res
}

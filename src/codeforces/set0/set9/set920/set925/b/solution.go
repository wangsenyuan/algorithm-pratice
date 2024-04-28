package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, x1, x2 := readThreeNums(reader)
	c := readNNums(reader, n)

	res := solve(x1, x2, c)

	if len(res) == 0 {
		fmt.Println("No")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("Yes\n")
	buf.WriteString(fmt.Sprintf("%d %d\n", len(res[0]), len(res[1])))

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

func solve(x1, x2 int, c []int) [][]int {
	n := len(c)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return c[id[i]] > c[id[j]]
	})

	create := func(f int, l int, side int) [][]int {
		res := make([][]int, 2)
		for i := 0; i < f; i++ {
			res[0] = append(res[0], id[i]+1)
		}
		for i := f; i < l; i++ {
			res[1] = append(res[1], id[i]+1)
		}
		if side == 1 {
			res[0], res[1] = res[1], res[0]
		}
		return res
	}

	for p, i := range id {
		b := (x1 + c[i] - 1) / c[i]
		if b <= p {
			j := id[p-b]
			a := (x2 + c[j] - 1) / c[j]
			if a+b <= p+1 {
				return create(a, p+1, 1)
			}
		}
		b = (x2 + c[i] - 1) / c[i]
		if b <= p {
			j := id[p-b]
			a := (x1 + c[j] - 1) / c[j]
			if a+b <= p+1 {
				return create(a, p+1, 0)
			}
		}
	}

	return nil
}

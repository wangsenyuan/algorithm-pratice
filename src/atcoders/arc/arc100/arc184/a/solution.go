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

	// first line not matter
	readString(reader)

	ask := func(x int, y int) int {
		fmt.Printf("? %d %d\n", x, y)
		return readNum(reader)
	}

	res := solve(ask)

	var buf bytes.Buffer
	buf.WriteString("!")
	for _, x := range res {
		buf.WriteString(fmt.Sprintf(" %d", x))
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

func solve(ask func(x int, y int) int) []int {

	specialCase := func() []int {
		i := 981
		res := make([][]int, 2)
		// 最后一个要留下来
		for j := i + 1; j < 1000; j++ {
			x := ask(i, j)
			res[x] = append(res[x], j)
		}
		res[0] = append(res[0], i)

		if len(res[0]) == 10 {
			res[1] = append(res[1], 1000)
		} else {
			res[0] = append(res[0], 1000)
		}
		// 1 肯定不是
		x := ask(1, i)
		if x == 1 {
			return res[0]
		}
		return res[1]
	}

	var cur []int
	for i := 1; i < 1000 && len(cur) < 10; i += 20 {
		if i == 981 && len(cur) == 0 {
			return specialCase()
		}
		res := make([][]int, 2)
		for j := i + 1; j < i+20; j++ {
			x := ask(i, j)
			res[x] = append(res[x], j)
		}
		res[0] = append(res[0], i)
		if len(res[0]) < 10 {
			cur = append(cur, res[0]...)
			continue
		}
		if len(res[1]) < 10 {
			cur = append(cur, res[1]...)
			continue
		}
		// len(res[0]) == 10 && len(res[1]) == 10
		x := ask(i, 1000)
		if x == 0 {
			return res[1]
		}
		// x == 1
		cur = append(cur, res[0]...)
	}
	sort.Ints(cur)

	return cur
}

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
	a := readNNums(reader, n)

	res := solve(a)

	if len(res) == 0 {
		fmt.Println("Impossible")
		return
	}
	var buf bytes.Buffer

	buf.WriteString("Possible\n")

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
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

type pair struct {
	first  int
	second int
}

func solve(a []int) []int {
	n := len(a)
	// 考虑这些人的到达顺序，
	// 要么 a[id[i]] = a[id[i-1]] + 1
	// 要 a[id[i]] % 3 = i % 3
	// 好像这两个是等价的
	// a[id[i]] % 3 = i
	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[a[i]] = append(pos[a[i]], i)
	}

	ans := make([]int, n)

	for i, c := 0, 0; i < n; i++ {
		for c >= 0 && len(pos[c]) == 0 {
			c -= 3
		}
		if c < 0 {
			return nil
		}
		ans[i] = pos[c][0] + 1
		pos[c] = pos[c][1:]
		c++
	}

	return ans
}

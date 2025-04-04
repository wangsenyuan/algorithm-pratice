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
	for range tc {
		l, r := readTwoNums(reader)
		res := solve(l, r)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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

const H = 30

func solve(l int, r int) []int {
	if l+2 == r {
		return []int{l, l + 1, r}
	}

	var a, b, c int

	i := H - 1
	for i > 0 && (l>>i)&1 == (r>>i)&1 {
		if (l>>i)&1 == 1 {
			a |= 1 << i
			b |= 1 << i
			c |= 1 << i
		}
		i--
	}
	// l[i] < r[i] and i > 0 (这样子才有3个数)
	mask := (1 << i) - 1
	a |= l & mask
	b |= mask
	c |= 1 << i
	if a == b {
		b++
	}
	if b == c {
		c++
	}

	return []int{a, b, c}
}

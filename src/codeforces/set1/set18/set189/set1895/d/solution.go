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
	a := readNNums(reader, n-1)
	b := solve(a)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", b[i]))
	}
	fmt.Println(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(a []int) []int {
	// expect b => a[i]  = b[i] ^ b[i+1]
	//  a[n-1] = b[n-1] ^ b[n]
	//  a[n-2] = b[n-2] ^ b[n-1] => a[n-1] ^ a[n-2] = b[n] ^ b[n-2]
	//  a[n-3] = b[n-3] ^ b[n-2] => a[n-1] ^ a[n-2] ^ a[n-3] = b[n] ^ b[n-3]
	// ...
	//  a[n-1] ^ a[n-2] ^ ... ^ a[n-i] = b[n] ^ b[n-i]
	//  a[n-1] ^ .... ^ a[1] = b[n] ^ b[1]
	n := len(a) + 1
	if n == 2 {
		return []int{0, a[0]}
	}
	c := make([]int, n)
	tr := new(Node)
	for i := 1; i < n; i++ {
		c[i] = c[i-1] ^ a[i-1]
		node := tr
		for d := 20; d >= 0; d-- {
			x := (c[i] >> d) & 1
			if node.children[x] == nil {
				node.children[x] = new(Node)
			}
			node = node.children[x]
		}
	}

	first := -1

	for num := 0; num < n; num++ {
		// can set b[1] = x or not
		node := tr
		var mx int
		for d := 20; d >= 0 && mx < n; d-- {
			x := (num >> d) & 1
			if node.children[1^x] != nil {
				mx |= 1 << d
				x ^= 1
			}
			node = node.children[x]
		}
		if mx < n {
			// find one
			first = num
			break
		}
	}
	b := make([]int, n+1)
	b[1] = first
	for i := 2; i <= n; i++ {
		b[i] = c[i-1] ^ b[1]
	}
	return b[1:]
}

type Node struct {
	children [2]*Node
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	query := func(arr []int) bool {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
		ans := readNum(reader)
		return ans == 1
	}

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, query)
		var buf bytes.Buffer
		buf.WriteString("!\n")
		for _, edge := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", edge[0], edge[1]))
		}
		fmt.Print(buf.String())
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
func solve(n int, fn func([]int) bool) [][]int {

	var loop func(nodes []int)

	var res [][]int

	buf := make([]int, 3)

	loop = func(nodes []int) {
		if len(nodes) < 2 {
			return
		}
		if len(nodes) == 2 {
			res = append(res, []int{nodes[0], nodes[1]})
			return
		}
		a := nodes[0]
		b := nodes[1]

		for i := 2; i < len(nodes); i++ {
			x := nodes[i]
			buf[0] = a
			buf[1] = x
			buf[2] = b
			tmp := fn(buf)
			if tmp {
				a = x
			}
		}
		// a direct edge between a & b
		res = append(res, []int{a, b})
		// how to remove a & b
		var A []int
		var B []int
		for i := 0; i < len(nodes); i++ {
			x := nodes[i]
			if x == a || x == b {
				continue
			}
			buf[0] = x
			buf[1] = a
			buf[2] = b
			tmp := fn(buf)
			if tmp {
				A = append(A, x)
			} else {
				B = append(B, x)
			}
		}
		A = append(A, a)
		loop(A)
		B = append(B, b)
		loop(B)
	}
	nodes := make([]int, n)
	for i := 0; i < n; i++ {
		nodes[i] = i + 1
	}
	loop(nodes)

	return res
}

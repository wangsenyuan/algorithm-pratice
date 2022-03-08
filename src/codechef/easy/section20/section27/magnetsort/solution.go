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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		polarity := readString(reader)
		res := solve(n, A, polarity)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(n int, A []int, polarity string) int {
	// 只要有n，s最多两次
	if sort.IntsAreSorted(A) {
		return 0
	}
	// at least 1
	first := make([]int, 2)
	last := make([]int, 2)
	for i := 0; i < 2; i++ {
		first[i] = -1
	}
	for i := 0; i < n; i++ {
		var x int
		if polarity[i] == 'S' {
			x = 1
		}
		if first[x] < 0 {
			first[x] = i
		}
		last[x] = i
	}
	buf := make([]int, n)

	check := func(i, j int) bool {
		copy(buf, A[:i])
		copy(buf[i:], A[i:j+1])
		sort.Ints(buf[i : j+1])
		copy(buf[j+1:], A[j+1:])
		return sort.IntsAreSorted(buf)
	}

	if first[0] >= 0 && last[1] > first[0] {
		if check(first[0], last[1]) {
			return 1
		}
	}
	if first[1] >= 0 && last[0] > first[1] {
		if check(first[1], last[0]) {
			return 1
		}
	}
	if first[0] >= 0 && first[1] >= 0 {
		return 2
	}
	return -1
}

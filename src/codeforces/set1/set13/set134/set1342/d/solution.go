package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	c := readNNums(reader, m)

	res := solve(a, c)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, arr := range res {
		buf.WriteString(fmt.Sprintf("%d", len(arr)))
		for _, x := range arr {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func solve(m []int, c []int) [][]int {
	k := len(c)
	freq := make([]int, k+1)
	for _, x := range m {
		freq[x]++
	}

	for i := k - 1; i > 0; i-- {
		freq[i] += freq[i+1]
	}

	var x int
	for i := 1; i <= k; i++ {
		x = max(x, (freq[i]+c[i-1]-1)/c[i-1])
	}

	res := make([][]int, x)

	sort.Ints(m)
	var id int
	for _, i := range m {
		res[id%x] = append(res[id%x], i)
		id++
	}

	return res
}

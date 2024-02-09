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
		boxes := make([][]int, 2*n-1)
		for i := 0; i < len(boxes); i++ {
			boxes[i] = readNNums(reader, 2)
		}
		res := solve(boxes)
		s := fmt.Sprintf("%v", res)
		s = s[1 : len(s)-1]
		buf.WriteString(fmt.Sprintf("YES\n%s\n", s))
	}
	fmt.Print(buf.String())
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

func solve(boxes [][]int) []int {
	n := len(boxes)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return boxes[id[i]][0] < boxes[id[j]][0]
	})

	sum := make([]int, 2)

	for i := 0; i < n; i += 2 {
		sum[0] += boxes[id[i]][1]
		if i+1 < n {
			sum[1] += boxes[id[i+1]][1]
		}
	}

	tot := sum[0] + sum[1]
	sum[1] += boxes[n-1][1]
	var res []int
	if sum[0]*2 >= tot {
		for i := 0; i < n; i += 2 {
			res = append(res, id[i]+1)
		}
	} else {
		for i := 1; i < n; i += 2 {
			res = append(res, id[i]+1)
		}
		res = append(res, id[n-1]+1)
	}
	return res
}

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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		res := solve(line[0], line[1:])
		if len(res) == 0 {
			buf.WriteString("NO\n")
		} else {
			buf.WriteString("YES\n")
			for _, e := range res {
				buf.WriteString(fmt.Sprintf("%d %d\n", e[0], e[1]))
			}
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, d []int) [][]int {
	// 1， 2， 3 有两种情况
	// 一种它们在一条path上，
	// 它们不在一条path上
	// 假设在一条path上 d12 + d23 = d31 的话，2在中间
	// 如果不在， 假设3在1-2变的外面，那么必然有一个点4，使得 1-3 = 1-4 + 3-4
	// 2-3 = 2 - 4 + 3 - 4
	// 且 1 - 4 + 2 - 4 = 1-2
	sum := d[0] + d[1] + d[2]
	if sum%2 == 1 {
		return nil
	}
	w := make([]int, 3)
	for i := 0; i < 3; i++ {
		w[i] = sum/2 - d[(i+1)%3]
	}
	sw := make([]int, 3)
	copy(sw, w)
	sort.Ints(sw)
	if sw[0] < 0 || sw[1] < 1 {
		return nil
	}
	var edges [][]int
	num := 3
	var center int
	if sw[0] == 0 {
		for i := 0; i < 3; i++ {
			if w[i] < w[center] {
				center = i
			}
		}
	} else {
		center = num
		num++
	}
	for i := 0; i < 3; i++ {
		before := center
		for j := 0; j < w[i]-1; j++ {
			edges = append(edges, []int{before, num})
			before = num
			num++
		}
		if w[i] > 0 {
			edges = append(edges, []int{before, i})
		}
	}
	if num > n {
		return nil
	}
	if num <= n {
		for num < n {
			edges = append(edges, []int{center, num})
			num++
		}
	}
	for _, e := range edges {
		e[0]++
		e[1]++
	}
	return edges
}

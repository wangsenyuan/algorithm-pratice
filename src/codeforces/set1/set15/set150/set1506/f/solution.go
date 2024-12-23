package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		r := readNNums(reader, n)
		c := readNNums(reader, n)
		res := solve(r, c)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(r []int, c []int) int {
	type pair struct {
		first  int
		second int
	}

	n := len(r)

	points := make([]pair, n)
	for i := 0; i < n; i++ {
		points[i] = pair{r[i], c[i]}
	}

	slices.SortFunc(points, func(a, b pair) int {
		return a.first - b.first
	})

	if points[0].first != 1 {
		points = append(points, pair{1, 1})
		slices.SortFunc(points, func(a, b pair) int {
			return a.first - b.first
		})
	}

	var res int

	for i := 0; i+1 < len(points); i++ {
		a, b := points[i].first, points[i].second
		c, d := points[i+1].first, points[i+1].second

		if (a+b)%2 == 0 {
			if a-b == c-d {
				res += c - a
				// 不能往前移动，不然就到不了d了
				continue
			}
			a++
			// 这时候a + b 是奇数，那么a+1+b+1也是奇数
			// 所以可以免费到到 c, b + c - a 这个位置
		}

		res += (b + c - a - d + 1) / 2
	}

	return res
}

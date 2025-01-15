package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		b, w := readTwoNums(reader)
		res := solve(b, w)
		if len(res) == 0 {
			buf.WriteString("NO\n")
			continue
		}
		buf.WriteString("YES\n")
		for _, cur := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
		}
	}

	fmt.Print(buf.String())
}

func solve(b int, w int) [][]int {
	x := min(b, w)
	y := max(b, w)
	if y > 3*x+1 {
		return nil
	}
	// y <= 3 * x + 1
	// wbw
	var res [][]int
	var center []int
	if x == w {
		center = []int{2, 2}
	} else {
		center = []int{2, 3}
	}
	if y == 3*x+1 {
		res = append(res, []int{center[0] - 1, center[1]})
		y--
	}
	// 先搞出 x+x个
	for i := 0; i < x; i++ {
		res = append(res, []int{center[0] + 2*i, center[1]})
		res = append(res, []int{center[0] + 2*i + 1, center[1]})
	}
	y -= x

	for i := 0; i < x && y > 0; i++ {
		res = append(res, []int{center[0] + 2*i, center[1] - 1})
		y--
		if y > 0 {
			res = append(res, []int{center[0] + 2*i, center[1] + 1})
			y--
		}
	}
	return res
}

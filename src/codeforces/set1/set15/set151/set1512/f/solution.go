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

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, c := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, n-1)
		res := solve(a, b, c)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

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

const inf = 1 << 60

func solve(a []int, b []int, c int) int {
	best := inf

	n := len(a)

	type pair struct {
		first  int
		second int
	}

	cur := pair{0, 0}

	for i := 0; i < n; i++ {
		// 如果在职位i上赚够足够的钱
		best = min(best, (c-cur.first+a[i]-1)/a[i]+cur.second)
		if i+1 < n {
			// 到下一个职位
			// cur.first + a[i] * d >= b[i]
			d := max(0, (b[i]-cur.first+a[i]-1)/a[i])
			// +1 是要move的那天
			cur = pair{cur.first + d*a[i] - b[i], cur.second + d + 1}
		}
	}

	return best
}

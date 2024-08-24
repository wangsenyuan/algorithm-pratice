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
		n, k := readTwoNums(reader)
		a := readNNums(reader, 2)
		b := readNNums(reader, 2)
		res := solve(n, k, a, b)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const inf = 1 << 60

func solve(n int, k int, a []int, b []int) int {
	if a[0] > b[0] {
		a, b = b, a
	}
	//如果移动前i个能够得到超过k的（总和）重叠区域
	// 因为都是从a，b移动后得到的区域
	// 那么这n个区间，移动的应该大体是相同的
	// 如果两个区域已经完全重叠了，那么继续增加1个重叠区域的cost是等于2的
	// 所以，如果可以在cost为1的情况下，增加重叠区域，就没有必要使用2
	touch := max(0, b[0]-a[1])
	overlap := max(0, min(a[1], b[1])-max(a[0], b[0]))

	if overlap*n >= k {
		return 0
	}

	k -= n * overlap

	width := max(b[1], a[1]) - min(a[0], b[0])
	width -= overlap

	best := inf

	for i := 1; i <= n; i++ {
		// 如果只是用前i个区域，且前面的区域已经完全被cover住了
		if (i-1)*width >= k {
			// no need to move on
			break
		}
		// 先要接触起来
		tmp := (i - 1) * (touch + width)
		tmp += touch
		x := k - (i-1)*width
		if x <= width {
			tmp += x
		} else {
			tmp += width
			tmp += 2 * (x - width)
		}
		best = min(best, tmp)
	}

	return best
}

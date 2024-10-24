package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func solve(a []int) int {
	n := len(a)

	// 这个只是计算出了最小值
	cur_min := a[0]
	sum := a[0]

	for i := 1; i < n; i++ {
		sum += a[i]
		cur_min = min(cur_min, sum/(i+1))
	}

	check := func(expect int) bool {
		// 能否得到expect的结果
		var sum2 int
		for i := n - 1; i >= 0; i-- {
			if a[i] <= cur_min+expect {
				sum2 += cur_min + expect - a[i]
			} else {
				// 需要填平
				// y = a[i] - (cur_min + expect)
				// y <= sum2
				y := a[i] - (cur_min + expect)
				if y > sum2 {
					// y > sum2 说明a[i]太大了，填完后面的，还会多出来
					return false
				}
				sum2 -= y
			}
		}
		return true
	}

	return sort.Search(slices.Max(a)-cur_min, check)
}

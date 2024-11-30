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

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

type pair struct {
	first  int
	second int
}

func solve(a []int) [][]int {
	// 需要知道最大值和最小值
	freq := make(map[int][]pair)
	// 假设 a[l...i...j] = a[i...j...r]
	// 要怎么把它们找出来呢？

	n := len(a)

	for i := 0; i < n; i++ {
		var sum int
		for j := i; j >= 0; j-- {
			sum += a[j]
			freq[sum] = append(freq[sum], pair{j, i})
		}
	}

	var ans [][]int

	for _, vs := range freq {
		if len(ans) > len(vs) {
			continue
		}
		slices.SortFunc(vs, func(x, y pair) int {
			return x.second - y.second
		})
		last := -1
		var cur [][]int
		for _, v := range vs {
			if last < v.first {
				cur = append(cur, []int{v.first + 1, v.second + 1})
				last = v.second
			}
		}
		if len(ans) < len(cur) {
			ans = cur
		}
	}

	return ans
}

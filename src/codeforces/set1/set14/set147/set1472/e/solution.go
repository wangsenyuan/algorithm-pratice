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
		H := make([]int, n)
		W := make([]int, n)
		for i := 0; i < n; i++ {
			H[i], W[i] = readTwoNums(reader)
		}
		res := solve(H, W)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(H []int, W []int) []int {
	// h[i] < h[j] and w[i] < w[j]
	// or h[i] < w[j] and w[i] < h[j]
	// for each j, find any i
	// h[i] < h[j] and w[i] < w[j] 这个很好处理
	// 但是第二个要咋搞呢？
	// 还是按照高度排序，但是同时记录该高度时，最小的w[i]对应的id

	type Rect struct {
		id     int
		height int
		width  int
	}

	n := len(H)
	rects := make([]Rect, n)
	for i := 0; i < n; i++ {
		rects[i] = Rect{i, H[i], W[i]}
	}

	// 如果考虑高度，应该找到最小的宽度
	sort.Slice(rects, func(i, j int) bool {
		return rects[i].height < rects[j].height || rects[i].height == rects[j].height && rects[i].width < rects[j].width
	})

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	// dp[i] = 在i（包括i）前的box的，拥有最小宽度的下标
	dp := make([]int, n)
	dp[0] = 0

	for i := 1; i < n; i++ {
		// 先查找 h[j] < h[i], and w[j] < w[i]
		// 这个没有问题，因为是按照高度在依次处理的
		j := search(i, func(j int) bool {
			return rects[j].height >= rects[i].height
		}) - 1
		// rects[j].height < rects[i].height
		if j >= 0 && rects[dp[j]].width < rects[i].width {
			ans[rects[i].id] = rects[dp[j]].id + 1
		}

		dp[i] = dp[i-1]
		if rects[i].width < rects[dp[i-1]].width {
			dp[i] = i
		}
	}

	for i := 0; i < n; i++ {
		if ans[rects[i].id] < 0 {
			// not known yet
			j := search(n, func(j int) bool {
				return rects[j].height >= rects[i].width
			}) - 1
			if j < 0 {
				continue
			}
			// rects[j].height < rects[i].width
			if rects[dp[j]].width < rects[i].height {
				ans[rects[i].id] = rects[dp[j]].id + 1
			}
		}
	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

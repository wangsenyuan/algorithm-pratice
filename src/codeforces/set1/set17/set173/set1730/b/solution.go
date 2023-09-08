package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
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
		x := readNNums(reader, n)
		y := readNNums(reader, n)
		res := solve(x, y)
		buf.WriteString(fmt.Sprintf("%.2f\n", res))
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

func solve(x []int, y []int) float64 {
	var mx float64 = -1
	var mn float64 = math.MaxFloat64

	for i := 0; i < len(x); i++ {
		a := float64(x[i] - y[i])
		mn = math.Min(mn, a)
		b := float64(x[i] + y[i])
		mx = math.Max(mx, b)
	}

	sum := mn + mx
	return sum / 2
}
func solve1(x []int, y []int) float64 {
	n := len(x)
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return x[id[i]] < x[id[j]]
	})
	// 在给定x0的情况下，计算到达x0需要的最多的时间
	// x0 主要在x处，或者两个的中间？
	// x0会处在x[i]和x[i+1]的中间吗？
	// 假设在x[i]处的cost为y[i]
	// x[i+1]处的cost时y[i+1]
	// 假设 x0在它们中间, y0, 是否有 y[i]=y0 or y[i+1] = y0
	// y[l] + x0 - x[l]
	// y[r] + x[r] - x0
	// 对于l来说，当x0在x[i]...x[i+1]中间移动时，不会影响 y[l] - x[l] (左边的最大值)
	// 所以 x0 向右边移动时，左边的值cost变大
	// 同样对于右边的r来说也是如此，当x0向右边移动时，它变小，向左边移动时，它变大
	// 因为它们时单调的，所以肯定在某个点处取到最小值
	// 就是找到一个x0, 是的 max(y[l] + x0 - x[l], y[r] + x[r] - x0)最小
	// 假设 y[l] + x0 - x[l] = y[r] + x[r] - x0

	R := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		j := id[i]
		R[i] = y[j] + x[j]
		if i+1 < n {
			R[i] = max(R[i], R[i+1])
		}
	}

	L := -(1 << 30)
	x0 := float64(x[id[0]])
	best := float64(R[0] - x[id[0]])

	for i := 0; i+1 < n; i++ {
		L = max(L, y[id[i]]-x[id[i]])

		j := id[i+1]
		tmp2 := max(L+x[j], R[i+1]-x[j])
		if x0 < 0 || best > float64(tmp2) {
			x0 = float64(x[j])
			best = float64(tmp2)
		}

		mid := float64(R[i+1]-L) / 2
		if float64(x[id[i]]) < mid && mid < float64(x[id[i+1]]) {
			tmp := math.Max(float64(L)+mid, float64(R[i+1])-mid)
			if best > tmp {
				x0 = mid
				best = tmp
			}
		}
	}

	return x0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

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
		a := readNNums(reader, n)
		res := solve(a, k)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
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

func solve(a []int, k int) bool {
	// sum[l...k] >= sum[k+1..]
	// 中间不能出现负值
	// 假设要从k出发，到达n+1位置，计算出最少需要的health
	if a[k-1] < 0 {
		return false
	}
	if k == 1 || k == len(a) {
		return true
	}
	if canReachRight(a, k-1) {
		return true
	}
	reverse(a)
	if canReachRight(a, len(a)-k) {
		return true
	}
	return false
}

type Pair struct {
	first  int
	second int64
}

func canReachRight(arr []int, k int) bool {
	// 主要有个问题解决不了。如果一古脑的往右边前进，会让health下降
	// 当无法前进时，（再前进就成负值了），再返回左边就来不及了
	// 所以合理的貌似
	// 是在可以往左边移动，并获得正收益的情况下，就应该往左，直到无法在现有基础上获得正的收益
	n := len(arr)
	cur := int64(arr[k])
	l, r := k, k+1
	var sum int64
	var mx int64
	for r < n {
		// 尽量往左
		for l > 0 && cur+sum+int64(arr[l-1]) >= 0 {
			sum += int64(arr[l-1])
			mx = max(mx, sum)
			l--
		}
		// 然后利用mx + cur 尽量向右
		if mx+cur+int64(arr[r]) < 0 {
			return false
		}
		cur += int64(arr[r])
		r++
	}
	return true
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}

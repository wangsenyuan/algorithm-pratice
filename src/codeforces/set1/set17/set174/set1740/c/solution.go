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
		w := readNNums(reader, n)
		res := solve(w)
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

func solve(w []int) int64 {
	// a, b, c
	sort.Ints(w)
	// 在第二个中放置一个，是否是最优的方案？还是在两边放置一个，中间放置其他的，是更优的方案呢？
	// 答案是否是始终是，把另外两个bag只有一个时，其中一个有其他的？
	// 先考虑b中只有一个，a和c中有多个
	// 然后假设在这种情况下，Dengklek 做出了选择(x, y, z)，他的目的是最小化 abs(x - y) + abs(y - z)
	// 如果 abs(x - y) > abs(y - z)
	// 那么在保留x, y, z的情况下，把所有a中的其他的都移动到c中，是不改变结果的
	// => a和c中，其中有一个，只需要1个，其他的都在另外一边
	// 然后考虑b的情况
	// 假设b中有两个 y1, y2, 且最终选择了y1, abs(x - y1) + abs(z - y1) < ... y2
	// 然后把y2移动到a中， 有没有可能 abs(x - y1) < abs(y1 - y2) 呢？
	// 其实是不能保证的，但是如果保证y1是y中的最大值，其比y更大的部分在另外两边的话，
	// 就没有问题了
	var best int64
	n := len(w)
	for i := 0; i+2 < n; i++ {
		// <= i全部在b中
		tmp := int64(w[n-1] - w[i])
		tmp += int64(w[i+1] - w[i])
		best = max(best, tmp)
	}

	for i := n - 1; i >= 2; i-- {
		// >= i 都在b中，其他的在a、c中;
		tmp := int64(w[i] - w[0])
		tmp += int64(w[i] - w[i-1])
		best = max(best, tmp)
	}
	// n-1 在a中，0在从中，其他在b中
	best = max(best, int64(w[n-1]-w[0]))

	return best
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

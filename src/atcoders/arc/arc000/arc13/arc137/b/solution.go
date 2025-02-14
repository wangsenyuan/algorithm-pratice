package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)

	res := solve(A)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const inf = 1 << 29

func solve(arr []int) int {
	// 计算出能够得到的最大sum和最小sum，结果为它们的差 + 1
	// 如何计算最大sum呢？
	// 原来为sum, 假设flip了[l...r]
	//   在[l..r]中间的1的个数 flip前是 a = sum[r] - sum[l-1]
	//   flip后变成了  b = r - l + 1 - a
	// 变化为 b - a = r - l + 1 - 2 * a
	// =   r - l + 1 - 2 * sum[r] + 2 * sum[l-1]
	// r - 2 * sum[r] - (l - 1 - 2 * sum[l-1])
	//  对于r 需要找到最小的 l, l - 2 * sum[l] 最小
	n := len(arr)
	var dp, fp int

	var mx, mn int = 0, 0
	var sum int
	for r := 0; r < n; r++ {
		sum += arr[r]
		mx = max(mx, r+1-2*sum-dp)
		mn = min(mn, r+1-2*sum-fp)
		dp = min(r+1-2*sum, dp)
		fp = max(r+1-2*sum, fp)

	}
	// sum 没有关系， 会被抵消掉
	return mx - mn + 1
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

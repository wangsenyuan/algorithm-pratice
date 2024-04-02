package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	cnt, res := solve(a)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(nums []int) (int, []int) {
	//n := len(x)
	mn := nums[0]
	mx := nums[0]
	var sum int
	for _, num := range nums {
		mn = min(mn, num)
		mx = max(mx, num)
		sum += num
	}
	cnt := make([]int, mx-mn+1)
	for _, num := range nums {
		cnt[num-mn]++
	}
	n := len(nums)

	if len(cnt) <= 2 {
		// no choice
		// 只有两个数
		// mn, mx
		// x + y = n
		// a * x + b * y = sum
		// 只有一组解

		return n, nums
	}
	res := make([]int, n)

	// cnt = 3
	// 假设cnt[i] = x个数
	// cnt[i+1] = y 个数
	// cnt[i+2] = z 个数
	// x * mn + y * (mn + 1) + z * (mn + 2) = sum
	// (x + y + z) * mn + (y + z)  + z  = sum
	// mx - 1 != 0
	count := func(arr []int) int {
		var ans int
		for i, x := range arr {
			ans += min(cnt[i], x)
		}
		return ans
	}

	best := []int{-1, -1, -1}
	for z := 0; z < n; z++ {
		tmp := sum - n*mn - z
		if tmp < z {
			continue
		}
		y := tmp - z
		x := n - y - z
		if x < 0 {
			continue
		}
		if best[0] < 0 || count([]int{x, y, z}) < count(best) {
			best = []int{x, y, z}
		}
	}

	for i := 0; i < best[0]; i++ {
		res[i] = mn
	}
	for i := best[0]; i < best[0]+best[1]; i++ {
		res[i] = mn + 1
	}
	for i := best[0] + best[1]; i < n; i++ {
		res[i] = mn + 2
	}
	return count(best), res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

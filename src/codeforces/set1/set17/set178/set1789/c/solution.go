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
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		ops := make([][]int, m)
		for i := 0; i < m; i++ {
			ops[i] = readNNums(reader, 2)
		}

		res := solve(a, ops)

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

func solve(a []int, ops [][]int) int {
	n := len(a)
	// 相邻的两个array, 连接后的肯定是 n + 1 （或者n）
	// 假设在位置i发生了变化,其他的n-1的数，肯定是相同的，同时新的数字v，如果a[i]相同, 相当于没有变化
	// 或者是新的数字，但新的数字，肯定要和其他n-1个数不同，所以，是 n+1
	// a, b, c，
	// a 和 c连接后呢？
	// 可以用贡献法吗？
	// 1，。。。。 n + m 在多少个数组内出现过？
	// 假设i在x个数组中出现过，在剩余的 (m + 1) - x 个数组没有出现
	// 那么i的贡献就是  x * (m + 1 - x) + x * (x - 1) / 2
	m := len(ops)
	// 这是一个差分数组
	prev := make([]int, n+m+1)
	for i := 0; i <= n+m; i++ {
		prev[i] = -2
	}
	for i := 0; i < n; i++ {
		// 这些数在数组0中出现了
		prev[a[i]] = -1
	}

	cnt := make([]int, n+m+1)

	// 除了v，其他的所有的数 + 1
	for i, op := range ops {
		p, v := op[0], op[1]
		u := a[p-1]
		if u == v {
			// no change
			continue
		}
		cnt[u] += i - prev[u]
		prev[v] = i
		a[p-1] = v
		prev[u] = -2
	}

	for i := 1; i <= n+m; i++ {
		if prev[i] >= -1 {
			cnt[i] += m - prev[i]
		}
	}

	var ans int

	for i := 1; i <= n+m; i++ {
		if cnt[i] > 0 {
			ans += cnt[i]*(m+1-cnt[i]) + cnt[i]*(cnt[i]-1)/2
		}
	}

	return ans
}

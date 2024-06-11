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
		a := readNNums(reader, n+m+1)
		b := readNNums(reader, n+m+1)
		res := solve(n, m, a, b)
		s := fmt.Sprintf("%v", res)
		buf.WriteString(s[1 : len(s)-1])
		buf.WriteByte('\n')
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

func solve(n int, m int, a []int, b []int) []int {
	N := len(a)
	// n + m + 1
	// 0 for programmer, 1 for tester

	cnt := make([]int, 2)
	var score int
	assign := make([]int, N)
	next := []int{-1, -1}
	// 假设把i给移除掉了
	// 且i是分配为programmer的，那么必须知道它后面那个第一个符合条件的，可以被配置为programmer
	// 但是却没有被配置为programmer的人
	// 这样的人，应该是cnt[0] = n, 且 a[i] > b[i]
	for i := 0; i < N-1; i++ {
		if cnt[0] == n && a[i] > b[i] && next[0] < 0 {
			// 他可以用来交换
			next[0] = i
		}
		if cnt[1] == m && next[1] < 0 && (a[i] <= b[i] || cnt[0] == n) {
			next[1] = i
		}

		if cnt[1] == m || cnt[0] < n && a[i] > b[i] {
			cnt[0]++
			assign[i] = 0
			score += a[i]
		} else {
			cnt[1]++
			assign[i] = 1
			score += b[i]
		}
	}

	// 最后一个人，肯定没有被选中
	ans := make([]int, N)
	ans[N-1] = score

	for i := 0; i < N-1; i++ {
		x := assign[i]
		tmp := score
		if x == 0 {
			tmp -= a[i]
		} else {
			tmp -= b[i]
		}
		j := next[x]
		if j >= 0 {
			// j需要交换角色，对应的，N-1，也需要变成另外一个角色
			if x == 0 {
				tmp -= b[j]
				tmp += a[j]
			} else {
				tmp -= a[j]
				tmp += b[j]
			}
			x ^= 1
		}
		if x == 0 {
			tmp += a[N-1]
		} else {
			tmp += b[N-1]
		}

		ans[i] = tmp
	}

	return ans
}

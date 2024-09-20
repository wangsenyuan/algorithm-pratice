package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	l := readNNums(reader, m)

	res := solve(n, l)

	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer

	for i := 0; i < m; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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

func solve(n int, ls []int) []int {
	m := len(ls)

	// 计算一下最多需要多少长度来处理后缀
	var dp int
	var sum int
	ans := make([]int, m+1)
	ans[m] = n
	for i := m - 1; i >= 0; i-- {
		l := ls[i]
		// 至少和后面的一样多，且不能低于自己的长度
		if l <= dp {
			dp++
		} else {
			// l > dp
			dp = l
		}
		// 这样子可以保证至少有一个颜色出现
		ans[i] = n - dp
		sum += l
	}

	if sum < n || dp > n {
		return nil
	}
	var pos int
	// dp[i+1]
	for i, l := range ls {
		ans[i] = pos
		if pos+l >= ans[i+1] {
			break
		}
		// pos + l < ans[i+1]
		pos += l
	}

	for i := 0; i < m; i++ {
		ans[i]++
	}

	return ans[:m]
}

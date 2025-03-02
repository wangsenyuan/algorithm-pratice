package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		n, m, k := readThreeNums(reader)
		res := solve(n, m, k)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func solve(n int, m int, k int) string {
	// n for 0, m for 1
	if max(n, m) < k {
		return "-1"
	}
	if k < max(n, m)-min(n, m) {
		// 真个串的value是这个，k不能更小
		return "-1"
	}

	letters := []byte("01")
	if n > m {
		letters[0], letters[1] = letters[1], letters[0]
		n, m = m, n
	}
	buf := make([]byte, n+m)

	// n <= m
	if n == 0 {
		if m != k {
			return "-1"
		}
		// 全部是第二个字符
		for i := range m {
			buf[i] = letters[1]
		}
		return string(buf)
	}
	// k = w - v (把v个0放在w个1的中间)
	// 先放k个1，然后放k个0
	// .... 最后剩下的必然是不足k个0和k个1
	// 但是这样子有个问题是0可能被快速的消耗掉了
	//	k个1，然后 0101 如果差不多相同就好了
	// m - k <= n 的话，似乎没有问题的（多个0，去消耗1）
	// 如果 m - k > n 的话 这个情况已经被排除了
	for i := range k {
		buf[i] = letters[1]
	}
	if m == k {
		// 剩下的全是0
		for i := k; i < n+m; i++ {
			buf[i] = letters[0]
		}
		return string(buf)
	}
	m1 := m - k

	for i := k; i < len(buf); {
		if n > 0 {
			buf[i] = letters[0]
			n--
			i++
		}
		if m1 > 0 {
			buf[i] = letters[1]
			m1--
			i++
		}
	}

	return string(buf)
}

func abs(num int) int {
	return max(num, -num)
}

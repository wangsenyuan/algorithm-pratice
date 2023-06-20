package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
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

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readNNums(reader, 5)
	n := first[0]
	w := readNNums(reader, n)
	res := solve(w, first[1], first[2], first[3], first[4])
	fmt.Println(res)
}

func solve(w []int, l int, r int, L int, R int) int64 {
	// 假设前x个由left hand收集，后面n-x个由right hand收集
	// 基础的cost是 l * sum(w[...x]) + r * sum[w[x...])
	// 如果 x - (n - x) > 1, 则需要而外的 x - (n -x + 1)次L
	var suf int64
	n := len(w)
	for i := 0; i < n; i++ {
		suf += int64(w[i])
	}

	best := int64(r)*suf + int64(n-1)*int64(R)

	var pref int64

	for i := 1; i <= n; i++ {
		pref += int64(w[i-1])
		suf -= int64(w[i-1])
		tmp := int64(l)*pref + int64(r)*suf
		if i-(n-i) > 1 {
			tmp += int64(L) * int64(i-(n-i)-1)
		} else if n-i-i > 1 {
			tmp += int64(R) * int64(n-i-i-1)
		}
		best = min(best, tmp)
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

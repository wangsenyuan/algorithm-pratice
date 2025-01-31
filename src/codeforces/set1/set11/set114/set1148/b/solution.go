package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 5)
	a := readNNums(reader, first[0])
	b := readNNums(reader, first[1])
	res := solve(first[2], first[3], first[4], a, b)

	fmt.Println(res)
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

func solve(ta int, tb int, k int, a []int, b []int) int {
	if k >= len(a) || k >= len(b) {
		return -1
	}
	// tb貌似没用，最后加上去就可以了
	// 应该取消a前面的航班，然后取消b后面尽可能多的航班
	best := b[0]

	for i, j := 0, 0; i < len(a) && i <= k; i++ {
		// 如果取消了前i个航班
		for j < len(b) && b[j] < a[i]+ta {
			j++
		}
		// j == len(b) or b[j] >= a[i] + ta
		cnt := k - i
		// 然后从j开始取消cnt个航班
		j1 := min(len(b), j+cnt)
		if j1 == len(b) {
			return -1
		}
		best = max(best, b[j1])
	}

	return best + tb
}

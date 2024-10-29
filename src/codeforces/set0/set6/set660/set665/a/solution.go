package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a, ta := readTwoNums(reader)
	b, tb := readTwoNums(reader)
	simion := readString(reader)
	res := solve(a, ta, b, tb, simion)

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

func solve(a int, ta int, b int, tb int, simion string) int {
	// simion的出发时刻和到达时刻可以计算出来
	// 那么他遇到的包括某个时刻开始出发，并在他到达时刻前出发的，从B站出发的车辆
	// 假设B站最早出发，且可以和simion遇到的车辆的，出发时间是x, x + tb > departure
	ss := strings.Split(simion, ":")
	var h, m int
	readInt([]byte(ss[0]), 0, &h)
	readInt([]byte(ss[1]), 0, &m)
	departure := h*60 + m
	first := max(5*60, departure-tb)
	// 必须正好是b的倍数
	first = (first-5*60)/b*b + 5*60
	if first+tb <= departure {
		first += b
	}
	arrival := min(departure+ta, 24*60-1)
	arrival = (arrival-5*60)/b*b + 5*60
	res := arrival/b - first/b
	if arrival < departure+ta {
		res++
	}

	return res
}

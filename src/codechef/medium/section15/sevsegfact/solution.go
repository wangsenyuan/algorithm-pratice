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
		line := readString(reader)
		var pos int
		for line[pos] != ' ' {
			pos++
		}
		num := line[:pos]
		pos++
		s := []byte(line)
		cost := make([]int, 7)
		for i := 0; i < 7; i++ {
			pos = readInt(s, pos, &cost[i]) + 1
		}
		res := solve(num, cost)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

var V []int64

func init() {
	V = append(V, 2)
	V = append(V, 5)
	V = append(V, 6)

	for i := 1; i <= 8; i++ {
		for x := 0; x < (1 << i); x++ {
			cnt := 0
			var num int64
			for j := 0; j < i; j++ {
				if (x>>j)&1 == 1 {
					num += 5
					cnt++
				} else {
					num += 6
				}
				num *= 10
			}
			num += 6
			if cnt%3 != 0 {
				V = append(V, num)
			}
		}
	}
}

func solve(num string, cost []int) int64 {

	dc := make([]int64, 9)
	dc[1] = int64(cost[1] + cost[2])
	dc[2] = int64(cost[0]) + int64(cost[1]) + int64(cost[3]) + int64(cost[4]) + int64(cost[6])
	dc[5] = int64(cost[0]) + int64(cost[2]) + int64(cost[3]) + int64(cost[5]) + int64(cost[6])
	dc[6] = dc[5] + int64(cost[4])

	res := dc[1]

	for _, x := range V {
		if div(num, x) {
			var cur int64
			for x > 0 {
				cur += dc[x%10]
				x /= 10
			}
			if cur < res {
				res = cur
			}
		}
	}

	return res
}

func div(num string, x int64) bool {
	var cur int64
	for i := 0; i < len(num); i++ {
		cur *= 10
		cur %= x
		cur += int64(num[i] - '0')
		cur %= x
	}
	return cur == 0
}

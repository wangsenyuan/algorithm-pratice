package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line := readNNums(reader, 6)
	n := line[0]
	P1 := line[1]
	P2 := line[2]
	P3 := line[3]
	T1 := line[4]
	T2 := line[5]
	events := make([][]int, n)
	for i := 0; i < n; i++ {
		events[i] = readNNums(reader, 2)
	}
	res := solve(n, P1, P2, P3, T1, T2, events)

	fmt.Println(res)
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
func solve(n int, P1 int, P2 int, P3 int, T1 int, T2 int, events [][]int) int64 {

	var res int64

	for i, event := range events {
		l, r := event[0], event[1]
		res += int64(P1) * int64(r-l)
		if i+1 < n {
			diff := events[i+1][0] - r
			if diff <= T1 {
				res += int64(diff) * int64(P1)
				continue
			}
			// diff > T1
			res += int64(T1) * int64(P1)
			diff -= T1
			if diff <= T2 {
				res += int64(diff) * int64(P2)
				continue
			}
			// diff > T2
			res += int64(T2) * int64(P2)
			diff -= T2
			res += int64(diff) * int64(P3)
		}
	}

	return res
}

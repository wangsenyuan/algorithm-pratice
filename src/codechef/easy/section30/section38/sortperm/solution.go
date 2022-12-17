package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		ans, swaps := solve(A)

		buf.WriteString(fmt.Sprintf("%d\n", ans))
		buf.WriteString(fmt.Sprintf("%d\n", len(swaps)))
		for _, swap := range swaps {
			buf.WriteString(fmt.Sprintf("%d %d\n", swap[0], swap[1]))
		}
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

func solve(P []int) (int, [][]int) {
	var ans int
	n := len(P)

	for i := 1; i <= n; i++ {
		ans += abs(P[i-1] - i)
	}
	ans /= 2
	var swaps [][]int
	dif := make([]int, n+1)
	pos := make([]int, n+1)

	for {
		ok := true
		for i := 1; i <= n; i++ {
			ok = ok && (P[i-1] == i)
			dif[P[i-1]] = P[i-1] - i
			pos[P[i-1]] = i
		}

		if ok {
			break
		}
		for i := 1; i <= n; i++ {
			if dif[P[i-1]] >= 0 {
				continue
			}
			for j := i - 1; j > 0; j-- {
				if dif[P[j-1]] > 0 {
					swaps = append(swaps, []int{i, j})
					P[i-1], P[j-1] = P[j-1], P[i-1]
					break
				}
			}
			break
		}
	}
	return ans, swaps
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

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
		n := readNum(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, n)
		}
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(A [][]int) string {
	n := len(A)

	best := make([]bool, n)
	min := make([]int, n)
	max := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i][j] == 1 {
				min[i]++
				max[i]++
			} else if A[i][j] == 0 {
				min[j]++
				max[j]++
			} else {
				max[i]++
				max[j]++
			}
		}
	}

	var score int
	for i := 0; i < n; i++ {
		if min[i] > score {
			score = min[i]
		}
	}

	for i := 0; i < n; i++ {
		if min[i] == score {
			best[i] = true
		}
	}

	for i := 0; i < n; i++ {
		if best[i] {
			var j int
			for ; j < n; j++ {
				if A[i][j] == 2 && best[j] {
					score++
					break
				}
			}
			if j < n {
				break
			}
		}
	}

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		if max[i] >= score {
			buf.WriteByte('1')
		} else {
			buf.WriteByte('0')
		}
	}
	return buf.String()
}

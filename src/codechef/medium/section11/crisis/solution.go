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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--

		N, X := readTwoNums(reader)
		A := readNNums(reader, N)
		res := solve(N, X, A)

		fmt.Printf("%.8f\n", res)
	}
}

func solve(N, X int, C []int) float64 {
	var num int64
	var den int64

	use := make([]bool, N)

	for i := 0; i < N; i++ {
		num += int64(i+1) * int64(C[i])
		den += int64(i + 1)
		use[i] = true
	}

	for j := 1; j <= N-X; j++ {
		var best float64
		var pos int
		for i := 0; i < N; i++ {
			if use[i] {
				num2 := num - int64(i+1)*int64(C[i])
				den2 := den - int64(i+1)
				tmp := float64(num2) / float64(den2)
				if tmp > best {
					best = tmp
					pos = i
				}
			}
		}
		num -= int64(pos+1) * int64(C[pos])
		den -= int64(pos + 1)
		use[pos] = false
	}

	return float64(num) / float64(den)
}

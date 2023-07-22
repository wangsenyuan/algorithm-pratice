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

		A := readNNums(reader, n)

		res := solve(A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(A []int) int64 {
	var ans int64

	var sum int64
	n := len(A)
	for i := n - 1; i >= 0; i-- {
		sum += int64(A[i])
		if sum > ans {
			ans = sum
		}
	}

	return ans
}

func solve1(A []int) int64 {
	n := len(A)
	st := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		if A[i] > 0 {
			st = append(st, i)
		}
	}
	var res int64
	var sum int64
	for i := 0; i < n; i++ {
		if A[i] >= 0 {
			sum += int64(A[i])
			continue
		}
		// A[i] < 0
		if int64(-A[i]) <= sum {
			sum += int64(A[i])
			continue
		}
		// -A[i] >= sum
		A[i] += int(sum)
		res += int64(-A[i])
		sum = 0

		for A[i] < 0 {
			j := st[0]
			if A[j] >= -A[i] {
				A[j] += A[i]
				A[i] = 0
				break
			}
			A[i] += A[j]
			A[j] = 0
			st = st[1:]
		}
	}
	return res
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := 1

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

const H = 30

func solve(A []int) int {
	n := len(A)
	if n == 2 {
		return -1
	}

	B := make([]int, n)

	for i := 0; i < n; i++ {
		cur := A[i]
		for cur > 0 {
			B[i]++
			cur >>= 1
		}
	}

	for i := 1; i+1 < n; i++ {
		if B[i-1] == B[i] && B[i] == B[i+1] {
			return 1
		}
	}
	// n <= 60
	ans := -1
	for i := 0; i < n; i++ {
		var a int
		for j := i; j+1 < n; j++ {
			a ^= A[j]
			var b int
			for k := j + 1; k < n; k++ {
				b ^= A[k]
				if a > b {
					if ans < 0 || ans > j-i+k-j-1 {
						ans = j - i + k - j - 1
					}
				}
			}
		}
	}

	return ans
}

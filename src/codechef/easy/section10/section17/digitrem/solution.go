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
		n, d := readTwoNums(reader)
		res := solve(n, d)
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

func solve(N int, D int) int {
	buf := make([]int, 10)
	var p int
	for N > 0 {
		buf[p] = N % 10
		p++
		N /= 10
	}
	buf = buf[:p]
	reverse(buf)

	if D == 0 {
		return solve0(buf)
	}
	if D == 9 {
		return solve9(buf)
	}
	return solveX(buf, D)
}

func solve0(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			// find 1111 -
			var a, b int
			for j := i; j < len(arr); j++ {
				a = a*10 + 1
				b = b*10 + arr[j]
			}
			return a - b
		}
	}

	return 0
}

func solve9(arr []int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 9 {
			a := 1
			b := 0
			for j := i; j < len(arr); j++ {
				a *= 10
				b = b*10 + arr[j]
				arr[j] = 0
			}
			ans := a - b
			if i > 0 {
				arr[i-1]++
				return ans + solve9(arr)
			}
			// i == 0
			return ans
		}
	}

	return 0
}

func solveX(arr []int, D int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == D {
			a := arr[i] + 1
			b := arr[i]
			for j := i + 1; j < len(arr); j++ {
				a = a * 10
				b = b*10 + arr[j]
			}

			return a - b
		}
	}

	return 0
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

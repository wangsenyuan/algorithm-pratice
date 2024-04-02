package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, a, b := readThreeNums(reader)

	res := solve(n, a, b)

	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		s := fmt.Sprintf("%v", res)
		fmt.Println(s[1 : len(s)-1])
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\r' || s[i] == '\n' {
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

func solve(n int, A int, B int) []int {

	for x := 0; x*A <= n; x++ {
		// x * A + y * B = n
		tmp := n - x*A
		if tmp%B == 0 {
			return permute(n, x, A, tmp/B, B)
		}
	}

	return nil
}

func permute(n int, x int, a int, y int, b int) []int {
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = i + 1
	}

	for i := 0; i < x; i++ {
		shift(res[i*a:(i+1)*a], 1)
	}

	offset := x * a
	for j := 0; j < y; j++ {
		shift(res[j*b+offset:(j+1)*b+offset], 1)
	}

	return res
}

func shift(arr []int, k int) {
	reverse(arr[:k])
	reverse(arr[k:])
	reverse(arr)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

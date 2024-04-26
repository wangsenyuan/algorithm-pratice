package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := readString(reader)

	res := solve(s)

	s = fmt.Sprintf("%v", res)

	fmt.Print(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(s string) (res []int) {
	n := len(s)
	var sum int
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' {
			sum++
		}
	}
	res = make([]int, n)
	// 0 for a, 1 for b
	pos := []int{-1, -1}
	for i := 0; i < n && sum > 0; i++ {
		if s[i] == 'a' {
			// 如果后面有a, 那么我们就应该将a放在suf的位置上
			if pos[1] >= 0 {
				res[pos[1]] = 1
			}
			pos[1] = -1
			if sum == 1 {
				// 如果这是最后一个a
				res[i] = 1
			}
			sum--
			pos[0] = i
		} else {
			if pos[0] >= 0 {
				res[pos[0]] = 1
			}
			pos[1] = i
			pos[0] = -1
		}
	}

	return
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readNum(reader)
	s := readString(reader)

	res := solve(s)

	fmt.Println(res)
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

func solve(s string) string {
	n := len(s)
	arr := make([]int8, n)
	for i := 0; i < n; i++ {
		arr[i] = int8(s[i] - '0')
	}
	best := make([]int8, n)

	tmp := make([]int8, n)
	get := func(i int) []int8 {
		// 第i位变成0，并且将其shift到最左边
		clear(tmp)

		var pushs int8
		if arr[i] != 0 {
			pushs = 10 - arr[i]
		}

		for j := 0; j < n; j++ {
			tmp[j] = (arr[(i+j)%n] + pushs) % 10
		}

		return tmp
	}

	comp := func(a, b []int8) bool {
		for i := 0; i < n; i++ {
			if a[i] < b[i] {
				return true
			}
			if a[i] > b[i] {
				return false
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		tmp := get(i)
		if i == 0 || comp(tmp, best) {
			copy(best, tmp)
		}
	}
	ans := make([]byte, n)

	for i := 0; i < n; i++ {
		ans[i] = byte(best[i] + '0')
	}

	return string(ans)
}

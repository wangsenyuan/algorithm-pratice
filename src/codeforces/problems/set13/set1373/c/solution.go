package main

import (
	"bufio"
	"fmt"
)

func main() {
	var tc int
	fmt.Scanf("%d", &tc)

	for tc > 0 {
		tc--
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(solve(len(s), []byte(s)))
	}
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

func solve(n int, s []byte) int64 {
	//fmt.Printf("%s\n", string(s))
	f := make([]int, n)
	var cur int
	var mn int
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			cur++
		} else if s[i] == '-' {
			cur--
		} else {
			break
		}
		if cur <= 0 {
			f[i] = abs(mn)
		}

		if cur < mn {
			mn = cur
		}
	}
	end := abs(mn)
	// end could be n, when init = end, then ok will be true, and it will terminate
	var res int64
	for i := 0; i < n; i++ {
		res += int64(end - f[i] + 1)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

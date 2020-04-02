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
		A, K := readTwoNums(reader)
		s, _ := reader.ReadString('\n')
		fmt.Println(solve(s[:len(s)-1], A, K))
	}
}

func solve(s string, A int, K int) int {
	n := len(s)
	buf := make([]byte, n)
	var res int

	good := func(k int, a int) bool {
		x := byte(a + 'a')
		for d := 1; d < n; d++ {
			j := k - d
			i := j - d
			if i < 0 {
				break
			}

			if buf[i] == x && buf[j] == x {
				return false
			}
		}
		return true
	}

	var loop func(i int, cnt int)

	loop = func(i int, cnt int) {
		if cnt > K {
			return
		}
		if i == n {
			res++
			return
		}

		for a := 0; a < A; a++ {
			if good(i, a) {
				buf[i] = byte(a + 'a')
				var cnt1 = cnt
				if buf[i] != s[i] {
					cnt1++
				}
				loop(i+1, cnt1)
			}
		}

	}

	loop(0, 0)

	return res
}

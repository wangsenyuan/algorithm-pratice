package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		role := func(s string) {
			fmt.Println(s)
		}
		play := func(num []int) (int, int) {
			if len(num) > 0 {
				fmt.Printf("%d %d\n", num[0], num[1])
			}
			return readTwoNums(reader)
		}

		solve(n, role, play)
	}

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

func solve(n int, role func(string), play func([]int) (int, int)) {
	if bits.OnesCount(uint(n))%2 == 0 {
		role("first")
	} else {
		role("second")
		a, b := play(nil)
		if bits.OnesCount(uint(a))%2 == 0 {
			n = a
		} else {
			n = b
		}
	}

	for n > 0 {
		a := 1 << (bits.Len(uint(n)) - 1)
		a, b := play([]int{a, n ^ a})
		if bits.OnesCount(uint(a))%2 == 0 {
			n = a
		} else {
			n = b
		}
	}
}

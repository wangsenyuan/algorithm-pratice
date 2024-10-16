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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
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

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	s := readString(reader)

	res, swaps := solve(s)

	fmt.Printf("%d\n%d %d\n", res, swaps[0], swaps[1])
}

func solve(s string) (int, []int) {
	n := len(s)

	var cnt int

	for _, x := range []byte(s) {
		if x == '(' {
			cnt++
		} else {
			cnt--
		}
	}

	if cnt != 0 {
		return 0, []int{1, 1}
	}

	buf := []byte(s)

	check := func(l int, r int) int {
		buf[l], buf[r] = buf[r], buf[l]

		var ans int
		var mv = 2
		var level int
		for i := 0; i < n; i++ {
			if buf[i] == '(' {
				level++
			} else {
				level--
			}
			if level < mv {
				ans = 1
				mv = level
			} else if level == mv {
				ans++
			}
		}

		buf[l], buf[r] = buf[r], buf[l]

		return ans
	}

	best := check(0, 0)
	swaps := []int{1, 1}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := check(i, j)
			if tmp > best {
				best = tmp
				swaps = []int{i + 1, j + 1}
			}
		}
	}

	return best, swaps
}

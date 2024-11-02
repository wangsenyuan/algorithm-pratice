package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, m)

	fmt.Println(res)
}

func solve(a []int, m int) int {
	n := len(a)
	check := func(k int) bool {
		// 能否在最大k操作后，得到一个非递减的序列
		x := a[0]
		if a[0]+k >= m {
			x = 0
		}
		for i := 1; i < n; i++ {
			// 当前这个数，进行k次操作后，仍然比前面的尽量小的数，小
			if a[i]+k < x {
				return false
			}

			if a[i]+k < m {
				x = max(x, a[i])
				continue
			}

			if (a[i]+k)%m < x {
				x = max(x, a[i])
			}
		}

		return true
	}

	return sort.Search(m, check)
}

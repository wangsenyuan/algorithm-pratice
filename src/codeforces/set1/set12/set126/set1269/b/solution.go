package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)

	return solve(m, a, b)
}

func solve(m int, a []int, b []int) int {

	groupByFreq := func(arr []int) [][]int {
		n := len(arr)
		freq := make(map[int]int)
		for _, x := range arr {
			freq[x]++
		}

		res := make([][]int, n+1)

		for k, v := range freq {
			res[v] = append(res[v], k)
		}
		for i := range res {
			sort.Ints(res[i])
		}
		return res
	}

	ax := groupByFreq(a)

	sort.Ints(b)

	bx := groupByFreq(b)
	n := len(a)

	c := make([]int, n)

	check := func(x int) bool {
		for i := 0; i < n; i++ {
			c[i] = (a[i] + x) % m
		}
		sort.Ints(c)
		for i := 0; i < n; i++ {
			if c[i] != b[i] {
				return false
			}
		}
		return true
	}

	for i := n; i > 0; i-- {
		if len(ax[i]) > 0 {
			// len(bx[i]) == len(ax[i])
			u := ax[i][0]
			res := m
			for j := 0; j < len(bx[i]); j++ {
				v := bx[i][j]
				x := v - u
				if x < 0 {
					x += m
				}
				if check(x) {
					res = min(res, x)
				}
			}
			return res
		}
	}

	return m
}

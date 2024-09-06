package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)

	res := solve(a, b, k)

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

func solve(a []int, b []int, k int) int {

	pre := func(arr []int) []int {
		n := len(arr)
		cnt := make([]int, n+2)
		for i := 0; i < n; i++ {
			if arr[i] == 0 {
				continue
			}
			j := i
			for i < n && arr[i] == 1 {
				i++
			}
			ln := i - j
			for j := 1; j <= ln; j++ {
				cnt[j] += ln - j + 1
			}
		}

		return cnt
	}

	pa := pre(a)
	pb := pre(b)

	calc := func(h int, w int) int {
		// h is for a, w is for b
		if h >= len(pa) || w >= len(pb) {
			return 0
		}
		return pa[h] * pb[w]
	}

	var res int

	for i := 1; i <= k/i; i++ {
		if k%i == 0 {
			res += calc(i, k/i)
			if i*i != k {
				res += calc(k/i, i)
			}
		}
	}

	return res

}

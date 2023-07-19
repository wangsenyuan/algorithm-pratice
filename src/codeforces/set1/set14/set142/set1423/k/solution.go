package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
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

const N = 1_000_010

var ans [N]int

func init() {
	pf := make([]int, N)
	for i := 2; i < N; i++ {
		if pf[i] == 0 {
			pf[i] = i
		}
		if N/i < i {
			continue
		}
		for j := i * i; j < N; j += i {
			if pf[j] == 0 {
				pf[j] = i
			}
		}
	}
	ans[1] = 1
	ans[2] = 2
	ans[3] = 3
	l := 2
	for j := 4; j < N; j++ {
		ans[j] = ans[j-1]
		for l <= j/l {
			if pf[l] == l && l*l == j {
				ans[j]--
			}
			l++
		}
		if pf[j] == j {
			ans[j]++
		}
	}
}

func solve(nums []int) []int {
	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = ans[n]
	}
	return res
}

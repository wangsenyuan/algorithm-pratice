package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	res := solve(a, k)
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

func solve(a []int, k int) string {
	n := len(a)

	prev := make([]int, n+1)
	next := make([]int, n+1)
	pos := make([]int, n+1)
	for i, x := range a {
		prev[i] = i - 1
		next[i] = i + 1
		pos[x] = i
	}

	ans := make([]byte, n)
	marked := make([]bool, n+1)
	var coach int

	id := "12"

	for x := n; x > 0; x-- {
		if marked[x] {
			continue
		}
		marked[x] = true
		i := pos[x]
		ans[i] = id[coach]
		l := prev[i]
		for j := 0; j < k && l >= 0; j++ {
			marked[a[l]] = true
			ans[l] = id[coach]
			l = prev[l]
		}
		r := next[i]
		for j := 0; j < k && r < n; j++ {
			marked[a[r]] = true
			ans[r] = id[coach]
			r = next[r]
		}

		if l >= 0 {
			next[l] = r
		}
		if r < n {
			prev[r] = l
		}

		coach ^= 1
	}

	return string(ans)
}

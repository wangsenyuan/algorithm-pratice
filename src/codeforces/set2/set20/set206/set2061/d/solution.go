package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) []bool {
	tc := readNum(reader)
	res := make([]bool, tc)
	for i := range tc {
		n, m := readTwoNums(reader)
		a := readNNums(reader, n)
		b := readNNums(reader, m)
		res[i] = solve(a, b)
	}
	return res
}

func solve(a []int, b []int) bool {
	freq := make(map[int]int)
	for _, x := range a {
		freq[x]++
	}

	var dfs func(x int) bool

	dfs = func(x int) bool {
		if freq[x] > 0 {
			freq[x]--
			return true
		}
		if x == 1 {
			return false
		}
		// freq[x] == 0
		return dfs(x/2) && dfs(x-x/2)
	}

	for _, x := range b {
		if !dfs(x) {
			return false
		}
	}
	for _, v := range freq {
		if v > 0 {
			return false
		}
	}
	return true
}

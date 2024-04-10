package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	p := readNNums(reader, n)

	res := solve(p)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

func solve(p []int) int {
	n := len(p)
	for i := 0; i < n; i++ {
		p[i]--
	}

	var cycles []int
	vis := make([]bool, n)

	for i := 0; i < n; i++ {
		if !vis[i] {
			if p[i] == i {
				vis[i] = true
				cycles = append(cycles, 1)
				continue
			}
			var cnt int
			j := i
			for !vis[j] {
				cnt++
				vis[j] = true
				j = p[j]
			}
			cycles = append(cycles, cnt)
		}
	}
	if len(cycles) <= 2 {
		return n * n
	}
	sort.Ints(cycles)

	m := len(cycles)

	res := (cycles[m-2] + cycles[m-1]) * (cycles[m-2] + cycles[m-1])
	for i := m - 3; i >= 0; i-- {
		res += cycles[i] * cycles[i]
	}
	return res
}

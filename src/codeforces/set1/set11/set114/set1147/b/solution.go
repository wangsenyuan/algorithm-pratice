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
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
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

func process(reader *bufio.Reader) bool {
	n, m := readTwoNums(reader)
	segs := make([][]int, m)
	for i := range m {
		segs[i] = readNNums(reader, 2)
	}
	return solve(n, segs)
}

func solve(n int, segs [][]int) bool {
	g := make([][]int, n)
	for _, seg := range segs {
		u, v := seg[0]-1, seg[1]-1
		u, v = min(u, v), max(u, v)
		g[u] = append(g[u], v-u)
		g[v] = append(g[v], u+n-v)
	}

	var arr []int
	for _, vs := range g {
		arr = append(arr, len(vs))
		sort.Ints(vs)
		arr = append(arr, vs...)
	}
	m := len(arr)
	lps := make([]int, m)
	for i := 1; i < m; i++ {
		j := lps[i-1]
		for j > 0 && arr[j] != arr[i] {
			j = lps[j-1]
		}
		if arr[j] == arr[i] {
			j++
		}
		lps[i] = j
	}

	var j int
	for i := 0; i < 2*m; i++ {
		for j > 0 && arr[i%m] != arr[j] {
			j = lps[j-1]
		}
		if arr[i%m] == arr[j] {
			j++
		}
		if (i+1)%m != 0 && j == m {
			return true
		}
		if j == m {
			j = lps[m-1]
		}
	}
	return false
}

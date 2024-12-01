package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (p []int, q [][]int) {
	n := readNum(reader)
	q = make([][]int, n-2)
	for i := range n - 2 {
		q[i] = readNNums(reader, 3)
	}
	p = solve(q)
	return
}

type pair struct {
	first  int
	second int
}

func solve(q [][]int) []int {
	n := len(q) + 2

	occ := make([][]int, n+1)
	var first int

	freq := make(map[pair]int)

	occ2 := make(map[pair][]int)

	get := func(a, b int) pair {
		a, b = min(a, b), max(a, b)
		return pair{a, b}
	}

	add := func(a, b, i int) {
		k := get(a, b)
		freq[k]++
		occ2[k] = append(occ2[k], i)
	}

	for i, cur := range q {
		for _, u := range cur {
			occ[u] = append(occ[u], i)
		}
		add(cur[0], cur[1], i)
		add(cur[0], cur[2], i)
		add(cur[1], cur[2], i)
	}
	for i := 1; i <= n; i++ {
		if len(occ[i]) == 1 {
			first = i
		}
	}

	ans := make([]int, n)
	id := occ[first][0]
	ans[0] = first
	// 然后把2，3，4 给排出来
	var arr []int
	for _, u := range q[id] {
		if u == first {
			continue
		}
		arr = append(arr, u)
	}

	tmp := occ2[get(arr[0], arr[1])]
	for _, j := range tmp {
		if j != id {
			id = j
			break
		}
	}
	i := 1
	c := sum(q[id]) - arr[0] - arr[1]
	if freq[get(arr[0], c)] > freq[get(arr[1], c)] {
		arr[0], arr[1] = arr[1], arr[0]
	}
	arr = append(arr, c)
	for _, x := range arr {
		ans[i] = x
		i++
	}

	for i < n {
		a, b := ans[i-2], ans[i-1]
		tmp := occ2[get(a, b)]
		for _, j := range tmp {
			if j != id {
				id = j
				break
			}
		}
		c := sum(q[id]) - a - b
		ans[i] = c
		i++
	}

	return ans
}

func sum(arr []int) int {
	var res int
	for _, x := range arr {
		res += x
	}
	return res
}

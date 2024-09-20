package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	res := solve(a)

	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(f []int) []int {
	n := len(f)

	a := make([]int, n+1)
	copy(a[1:], f)

	gift := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		gift[a[i]] = true
	}

	// 那么没有receive的才需要处理
	var arr []int
	for i := 1; i <= n; i++ {
		if !gift[i] {
			arr = append(arr, i)
		}
	}

	vis := make([]bool, n+1)
	head := make([]int, n+1)
	first := -1

	for _, i := range arr {
		j := i
		for j != 0 && !vis[j] {
			vis[j] = true
			head[i] = j
			j = a[j]
		}
		if first < 0 {
			first = i
		} else {
			// 把第一个和它连起来
			a[head[first]] = i
			head[first] = head[i]
		}
	}
	if first > 0 {
		a[head[first]] = first
	}
	return a[1:]
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	w := readNNums(reader, n)
	b := readNNums(reader, m)
	res := solve(w, b)
	fmt.Println(res)
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
func solve(w []int, b []int) int {
	n := len(w)
	stack := make([]int, n)

	var it int
	vis := make([]bool, n+1)
	for _, i := range b {
		if vis[i] {
			continue
		}
		stack[it] = i
		it++
		vis[i] = true
	}

	var res int
	for _, i := range b {
		// find position of i
		j := 0
		for j < n && stack[j] != i {
			j++
		}
		// stack[j] = i
		for j > 0 {
			res += w[stack[j-1]-1]
			stack[j] = stack[j-1]
			j--
		}
		stack[0] = i
	}

	return res
}

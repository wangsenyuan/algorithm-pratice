package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
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

func process(reader *bufio.Reader) (res []int, friends [][]int, alex []int) {
	n := readNum(reader)
	friends = make([][]int, n)
	for i := range n {
		friends[i] = readNNums(reader, n)
	}
	alex = readNNums(reader, n)
	res = solve(friends, alex)
	return
}

func solve(friends [][]int, alex []int) []int {
	n := len(friends)
	ans := make([]int, n)
	pos := make([][]int, n)
	for i := range n {
		ans[i] = -1
		pos[i] = make([]int, n)
		for j, v := range friends[i] {
			pos[i][v-1] = j
		}
	}

	p := make([]int, n)
	for i, v := range alex {
		p[v-1] = i
	}
	first, second := -1, -1

	res := make([]int, n)
	for i := range n {
		if first == -1 || p[i] < p[first] {
			second = first
			first = i
		} else if second == -1 || p[i] < p[second] {
			second = i
		}
		// 更新答案
		for j := range n {
			card := first
			if j == first {
				card = second
			}
			if card < 0 {
				continue
			}
			if ans[j] == -1 || pos[j][card] < pos[j][ans[j]] {
				res[j] = i + 1
				ans[j] = card
			}
		}
	}
	return res
}

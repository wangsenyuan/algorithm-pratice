package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, t, r := readThreeNums(reader)
	w := readNNums(reader, n)
	res := solve(t, r, w)

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

func solve(t int, r int, w []int) int {
	sort.Ints(w)
	if t < r {
		return -1
	}
	n := len(w)
	var res int

	offset := t + 1

	marked := make([]bool, offset+w[n-1]+1)

	for i := 0; i < n; i++ {
		var cnt int
		for j := 1; j <= t; j++ {
			if marked[w[i]-j+offset] {
				cnt++
			}
		}
		for j := 1; j <= t && cnt < r; j++ {
			if !marked[w[i]-j+offset] {
				marked[w[i]-j+offset] = true
				cnt++
				res++
			}
		}
	}

	return res
}

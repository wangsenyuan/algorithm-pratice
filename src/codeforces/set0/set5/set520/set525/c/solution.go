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
	a := readNNums(reader, n)
	res := solve(a)

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

func solve(a []int) int {
	n := len(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	// 如果到i为止，能够组成一个新的矩形，就组成一个？
	// 假设不是这样，假设i可以和后面的组成一个矩形，但是必须减少一个前面的一个矩形，
	// 能在后面增加更多的矩形吗？
	// h * a[i], a[i] * w1
	var res int

	for i := 0; i+1 < n; i++ {
		if a[i]-a[i+1] > 1 {
			continue
		}
		// a[i] - a[i+1] < 1
		j := i + 2
		for j+1 < n && a[j]-a[j+1] > 1 {
			j++
		}
		if j+1 >= n {
			break
		}
		res += a[i+1] * a[j+1]
		i = j + 1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

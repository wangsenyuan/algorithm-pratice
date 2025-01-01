package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}

	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) []int {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	bases := make([][]int, m)
	for i := range m {
		bases[i] = readNNums(reader, 2)
	}
	return solve(a, bases)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, bases [][]int) []int {
	n := len(a)
	ships := make([]pair, n)
	for i := range n {
		ships[i] = pair{a[i], i}
	}

	slices.SortFunc(bases, func(a, b []int) int {
		return a[0] - b[0]
	})

	slices.SortFunc(ships, func(a, b pair) int {
		return a.first - b.first
	})

	ans := make([]int, n)
	var sum int
	for i, j := 0, 0; i < n; i++ {
		for j < len(bases) && bases[j][0] <= ships[i].first {
			sum += bases[j][1]
			j++
		}
		ans[ships[i].second] = sum
	}
	return ans
}

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
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
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
	a := make([][]int, n)
	b := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	for i := range n {
		b[i] = readNNums(reader, m)
	}
	return solve(a, b)
}

func solve(a [][]int, b [][]int) bool {
	n := len(a)
	m := len(a[0])
	arr1 := make([][]int, n+m)
	arr2 := make([][]int, n+m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			arr1[i+j] = append(arr1[i+j], a[i][j])
			arr2[i+j] = append(arr2[i+j], b[i][j])
		}
	}

	for i := 0; i < n+m; i++ {
		sort.Ints(arr1[i])
		sort.Ints(arr2[i])
		for j := 0; j < len(arr1[i]); j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}

	return true
}

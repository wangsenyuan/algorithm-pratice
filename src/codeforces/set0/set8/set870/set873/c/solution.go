package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m, k := readThreeNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	res := solve(mat, k)

	fmt.Printf("%d %d\n", res[0], res[1])
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

func solve(mat [][]int, k int) []int {
	ans := make([]int, 2)
	n := len(mat)
	m := len(mat[0])
	col := make([]int, n)
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			col[i] = mat[i][j]
		}
		tmp := solveCol(col, k)
		ans[0] += tmp[0]
		for i := 0; i < tmp[1]; i++ {
			// 在top上面的1要删除
			ans[1] += mat[i][j]
		}
	}
	return ans
}

func solveCol(arr []int, k int) []int {
	n := len(arr)
	var sum int
	ans := []int{0, n}

	for i, r := n-1, n-1; i >= 0; i-- {
		sum += arr[i]
		if r-i+1 > k {
			// 长度不超过k
			sum -= arr[r]
			r--
		}
		if arr[i] == 1 && sum >= ans[0] {
			ans[0] = sum
			ans[1] = i
		}
	}
	return ans
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		mat := make([][]int, n)
		for i := 0; i < n; i++ {
			mat[i] = readNNums(reader, m)
		}
		res := solve(mat)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(mat [][]int) int64 {
	// 选择1/4的区域，让另外三个区域和它形成回文
	// 中间的+需要单独处理
	n := len(mat)
	m := len(mat[0])
	a := n / 2
	b := m / 2

	var res int64
	arr := make([]int, 4)
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			arr[0] = mat[i][j]
			arr[1] = mat[i][m-1-j]
			arr[2] = mat[n-1-i][j]
			arr[3] = mat[n-1-i][m-1-j]
			sort.Ints(arr)
			// arr[1] or arr[2]
			if (arr[1]+arr[2])%2 == 0 {
				res += count(arr, (arr[1]+arr[2])/2)
			} else {
				res += min(count(arr, arr[1]), count(arr, arr[2]))
			}
		}
	}
	if n%2 == 1 {
		res += makePalindrome(mat[n/2])
	}
	if m%2 == 1 {
		col := make([]int, n)
		for i := 0; i < n; i++ {
			col[i] = mat[i][m/2]
		}
		res += makePalindrome(col)
	}

	return res
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func count(arr []int, mid int) int64 {
	var res int64
	for _, x := range arr {
		res += abs(int64(x - mid))
	}
	return res
}

func makePalindrome(arr []int) int64 {
	var res int64
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		res += abs(int64(arr[i] - arr[j]))
	}
	return res
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		nums := readNNums(reader, n)
		res := solve(nums, k)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int, k int) int {

	n := len(A)

	// 如果 A[i] == A[i-1], 它们可以按照同样的规则区分
	// A[i] > A[i-1] 肯定会增加一个数
	var cnt int
	for i := 1; i <= n; i++ {
		if i == n || A[i] > A[i-1] {
			cnt++
		}
	}

	if cnt <= k {
		return 1
	}
	if k == 1 {
		return -1
	}
	// 至少需要两个数组
	// 考虑处理到A[i]时，A[i] > A[i-1],把A[i]分到前面不同的分组后, 尽量保持和A[i-1]的划分一致（这样可以不增加新的数字)
	// 剩下 A[i] - A[i-1]， 这时有两个选择，一个选择时把它加到最后一个数组中, B中， 此时B的计数+1
	//  1, 2, 3, 4  & k = 2
	//  [1] => [1, 2] => [1, 2, 2], [0, 0, 1] => [1, 2, 2, 2], [0, 0, 1, 1], [0, 0, 0, 1]
	m := 1
	var dist = 1
	for i := 1; i < n; i++ {
		if A[i] == A[i-1] {
			continue
		}
		// A[i] > A[i-1]
		if dist == k {
			m++
			// 前面的需要使用0补充
			dist = 1
		}
		dist++
	}

	return m
}

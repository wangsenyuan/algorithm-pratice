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
		A := readNNums(reader, n)
		B := readNNums(reader, n)
		res := solve(n, m, A, B)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, m int, A []int, B []int) int {
	M := int64(m)

	var all int64

	for i := 0; i < n; i++ {
		all += int64(A[i])
	}

	if all < M {
		return -1
	}

	que := make([][]int, 2)
	for i := 0; i < 2; i++ {
		que[i] = make([]int, 0, n)
	}

	for i := 0; i < n; i++ {
		b := B[i] - 1
		que[b] = append(que[b], A[i])
	}

	for i := 0; i < 2; i++ {
		sort.Ints(que[i])
		reverse(que[i])
	}
	preSum := make([]int64, len(que[1])+1)
	for i := 0; i < len(que[1]); i++ {
		preSum[i+1] = preSum[i] + int64(que[1][i])
	}
	// if use all type 2
	j := len(preSum) - 1
	var best = 2 * n
	for j >= 0 && preSum[j] >= M {
		best = j * 2
		if j == 0 || preSum[j-1] < M {
			break
		}
		j--
	}
	var sum int64
	for i := 1; i <= len(que[0]); i++ {
		sum += int64(que[0][i-1])
		if sum+preSum[j] < M {
			// not a valid solution
			continue
		}
		tmp := M - sum
		// need tmp
		for j > 0 && preSum[j-1] >= tmp {
			j--
		}
		if best > i+j*2 {
			best = i + j*2
		}
	}

	return best
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func search(n int, fn func(int) bool) int {
	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		if fn(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

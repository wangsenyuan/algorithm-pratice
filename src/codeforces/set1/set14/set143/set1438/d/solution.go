package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	cnt, res := solve(A)

	if cnt < 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("YES\n%d\n", cnt))
	for i := 0; i < cnt; i++ {
		for j := 0; j < len(res[i]); j++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i][j]))
		}
		buf.WriteByte('\n')
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

func solve(A []int) (int, [][]int) {
	n := len(A)
	var x int
	for i := 0; i < n; i++ {
		x ^= A[i]
	}
	if n%2 == 0 && x != 0 {
		return -1, nil
	}
	if n%2 == 0 {
		n--
	}
	arr := make([]int, n+1)
	copy(arr[1:], A)
	var res [][]int
	for i := 1; i < n; i += 2 {
		res = append(res, []int{i, i + 1, i + 2})
		a := arr[i] ^ arr[i+1] ^ arr[i+2]
		arr[i] = a
		arr[i+1] = a
		arr[i+2] = a
	}
	for i := n - 2; i > 2; i -= 2 {
		res = append(res, []int{i - 2, i - 1, i})
		a := arr[i-2] ^ arr[i-1] ^ arr[i]
		arr[i-2] = a
		arr[i-1] = a
		arr[i] = a
	}

	return len(res), res
}

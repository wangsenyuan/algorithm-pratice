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
		R, G, B := readThreeNums(reader)
		ok, s, e := solve(R, G, B)
		if !ok {
			buf.WriteString("-1\n")
		} else {
			buf.WriteString(s)
			buf.WriteByte('\n')
			for _, edge := range e {
				buf.WriteString(fmt.Sprintf("%d %d\n", edge[0], edge[1]))
			}
		}
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
func solve(R, G, B int) (bool, string, [][]int) {
	if B > R+G {
		return false, "", nil
	}

	//B <= R + G
	// 0...R-1
	n := R + G + B
	arr := make([]byte, n)
	arr[0] = 'R'
	arr[1] = 'G'
	var res [][]int
	res = append(res, []int{1, 2})
	pos := 3
	for R > 1 {
		arr[pos-1] = 'R'
		res = append(res, []int{2, pos})
		pos++
		R--
	}

	for G > 1 {
		arr[pos-1] = 'G'
		res = append(res, []int{1, pos})
		pos++
		G--
	}

	for i := 1; i <= R+G && B > 0; i++ {
		arr[pos-1] = 'B'
		res = append(res, []int{i, pos})
		pos++
		B--
	}

	return true, string(arr), res
}

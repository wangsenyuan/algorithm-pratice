package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	var buf bytes.Buffer

	query := func(arr []int) uint64 {
		m := len(arr)

		buf.WriteString(fmt.Sprintf("? %d", m))

		for i := 0; i < len(arr); i++ {
			buf.WriteString(fmt.Sprintf(" %d", arr[i]))
		}
		fmt.Println(buf.String())
		buf.Reset()

		s, _ := reader.ReadBytes('\n')
		var r uint64
		readUint64(s, 0, &r)
		return r
	}

	res := solve(n, query)

	buf.WriteString("!")

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Println(buf.String())
}

const Q = 13

func solve(n int, query func([]int) uint64) []uint64 {

	ask := make([][]int, Q)

	for i := 0; i < Q; i++ {
		ask[i] = make([]int, 0, n)
	}

	masks := make([]int, n+1)

	var assigned int
	for i := 1; i < 1<<Q && assigned < n; i++ {
		if popCount(i) != Q/2 {
			continue
		}
		masks[assigned] = i

		for j := 0; j < Q; j++ {
			if i>>j&1 == 0 {
				ask[j] = append(ask[j], assigned+1)
			}
		}

		assigned++
	}

	val := make([]uint64, Q)

	for i := 0; i < Q; i++ {
		if len(ask[i]) > 0 {
			val[i] = query(ask[i])
		}
	}

	ans := make([]uint64, n)

	for i := 0; i < n; i++ {
		for j := 0; j < Q; j++ {
			if masks[i]>>j&1 == 1 {
				ans[i] |= val[j]
			}
		}
	}

	return ans
}

func popCount(num int) int {
	var res int

	for num > 0 {
		res += num & 1
		num >>= 1
	}

	return res
}

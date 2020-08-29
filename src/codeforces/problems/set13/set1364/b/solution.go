package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		P := readNNums(reader, n)
		res := solve(n, P)
		fmt.Println(len(res))
		var buf bytes.Buffer
		for i := 0; i < len(res); i++ {
			buf.WriteString(strconv.Itoa(res[i]))
			buf.WriteByte(' ')
		}
		fmt.Println(buf.String())
	}
}

func solve(n int, P []int) []int {
	// find moutain peek
	var res = make([]int, 0, n)
	res = append(res, P[0])
	for i := 0; i < n-1; {
		j := i + 1
		if P[j] > P[i] {
			for j < n && P[j] > P[j-1] {
				j++
			}

		} else {
			for j < n && P[j] < P[j-1] {
				j++
			}
		}
		// j - 1 is the peek
		j--

		if i == j {
			break
		}

		res = append(res, P[j])
		i = j
	}
	return res
}

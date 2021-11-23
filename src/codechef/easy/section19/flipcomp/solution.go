package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		S, _ := reader.ReadString('\n')
		S = S[:len(S)-1]
		res := solve(S)
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

func solve(S string) int {
	a := flip("00"+S+"00", '0')
	b := flip("11"+S+"11", '1')
	if a < b {
		return a
	}
	return b
}

func flip(s string, x byte) int {
	n := len(s)
	var res int
	buf := []byte(s)
	for i := 2; i+2 < n; i++ {
		// 01010, or 10101
		flip := buf[i] == x && buf[i+1] != x && buf[i-1] != x && (buf[i-2] != x || buf[i+2] != x)
		if flip {
			buf[i] = byte('0' + '1' - buf[i])
			res++
		}
	}

	// compress and flip it
	for i := 0; i < n; i++ {
		if buf[i] == x {
			continue
		}
		j := i
		for i+2 < n && buf[i] != x {
			i++
		}

		res++
		if i-j > 1 {
			res++
		}
	}
	return res
}

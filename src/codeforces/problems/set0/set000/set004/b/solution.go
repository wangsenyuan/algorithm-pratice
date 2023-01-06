package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	d, s := readTwoNums(reader)

	time := make([][]int, d)

	for i := 0; i < d; i++ {
		time[i] = readNNums(reader, 2)
	}

	ok, res := solve(d, s, time)

	if ok {
		var buf bytes.Buffer
		buf.WriteString("YES\n")
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		fmt.Println(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(day int, sum int, time [][]int) (bool, []int) {
	var x, y int
	res := make([]int, day)

	for i := 0; i < day; i++ {
		x += time[i][0]
		y += time[i][1]
		res[i] = time[i][0]
	}

	ok := x <= sum && sum <= y
	if !ok {
		return false, nil
	}
	sum -= x
	for i := 0; i < day; i++ {
		tmp := min(time[i][1]-time[i][0], sum)
		res[i] += tmp
		sum -= tmp
	}

	if sum != 0 {
		return false, nil
	}

	return true, res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

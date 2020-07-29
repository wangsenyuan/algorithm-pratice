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
		n, k := readTwoNums(reader)
		S := make([][]int, k)
		for i := 0; i < k; i++ {
			var c int
			b, _ := reader.ReadBytes('\n')
			pos := readInt(b, 0, &c)
			S[i] = make([]int, c)
			for j := 0; j < c; j++ {
				pos = readInt(b, pos+1, &S[i][j])
			}
		}
		res := solve(n, k, S, reader)
		var buf bytes.Buffer
		buf.WriteByte('!')
		for i := 0; i < k; i++ {
			buf.WriteByte(' ')
			buf.WriteString(strconv.Itoa(res[i]))
		}
		fmt.Println(buf.String())
		ok, _ := reader.ReadString('\n')
		if ok[0] != 'C' {
			break
		}
	}
}

const N = 1010

var arr []int

func init() {
	arr = make([]int, N)
}

func solve(n int, k int, S [][]int, reader *bufio.Reader) []int {
	comp := make([]int, n+1)
	for i, s := range S {
		for _, x := range s {
			comp[x] = i
		}
	}
	var buf bytes.Buffer

	X := askBetween(1, n, buf, reader)
	// X is the largest number

	var begin, end = 1, n

	for begin < end {
		mid := (begin + end) / 2
		tmp := askBetween(begin, mid, buf, reader)
		if tmp == X {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	// we know the index A[end] == X
	var j int
	for i := 1; i <= n; i++ {
		if comp[i] != comp[end] {
			arr[j] = i
			j++
		}
	}
	Y := ask(arr[:j], buf, reader)
	// result is XXXYXXX
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		if comp[end] == i {
			ans[i] = Y
		} else {
			ans[i] = X
		}
	}
	return ans
}

func askBetween(begin, end int, buf bytes.Buffer, reader *bufio.Reader) int {
	n := end - begin + 1
	for i := 0; i < n; i++ {
		arr[i] = begin + i
	}
	return ask(arr[:n], buf, reader)
}

func ask(arr []int, buf bytes.Buffer, reader *bufio.Reader) int {
	buf.WriteString(fmt.Sprintf("? %d", len(arr)))
	for i := 0; i < len(arr); i++ {
		buf.WriteString(fmt.Sprintf(" %d", arr[i]))
	}
	fmt.Println(buf.String())
	buf.Reset()

	return readNum(reader)
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		e, o := readTwoNums(reader)
		res := solve(n, e, o)

		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < len(res); i++ {
				buf.WriteString(fmt.Sprintf("%d ", res[i]))
			}
			buf.WriteByte('\n')
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

func solve(n int, e int, o int) []int {
	N := int64(n)
	//E := int64(e)
	O := int64(o)

	// x + y = N + 1 = a
	// x * y = O = b
	// x - a * x + b = 0 => x = (a +/- sqrt(a * a - 4 * b)) / 2
	sq := sqrt((N+1)*(N+1) - 4*O)

	if sq < 0 {
		return nil
	}
	var x int
	if ((N+1)+sq)%2 == 0 {
		x = int((N + 1 + sq) / 2)
	} else if ((N+1)-sq)%2 == 0 {
		x = int((N + 1 - sq) / 2)
	} else {
		return nil
	}
	//y := (n + 1) - x

	res := make([]int, n)

	if x > 0 && o > 0 {
		res[x-1] = 1
	}

	return res
}

func sqrt(n int64) int64 {
	x := int64(math.Sqrt(float64(n)))

	if x*x == n {
		return x
	}

	if (x+1)*(x+1) == n {
		return x + 1
	}

	if (x-1)*(x-1) == n {
		return x - 1
	}

	return -1
}

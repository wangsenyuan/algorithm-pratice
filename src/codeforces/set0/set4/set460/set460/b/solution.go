package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
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

	a, b, c := readThreeNums(reader)

	res := solve(a, b, c)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func solve(a, b, c int) []int {
	// x = b * s(x) ** a + c
	// s(x) = digits sum of x
	// x <= 1e9
	// s(x) <= 9 * 9 = 81
	var res []int
	for sx := 0; sx < 81; sx++ {
		x := int64(b)*pow(sx, a) + int64(c)
		if x < 1e9 && x > 0 && digits_sum(x) == sx {
			res = append(res, int(x))
		}
	}
	sort.Ints(res)
	return res
}

func digits_sum(num int64) int {
	var res int
	for num > 0 {
		res += int(num % 10)
		num /= 10
	}
	return res
}

func pow(a, b int) int64 {
	var res int64 = 1
	A := int64(a)
	for b > 0 {
		if b&1 == 1 {
			res *= A
		}
		A *= A
		b >>= 1
	}
	return res
}

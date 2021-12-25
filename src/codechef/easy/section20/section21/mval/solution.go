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

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(n, A)
		buf.WriteString(fmt.Sprintf("%d\n", getMaxSum(n, A)))
		buf.WriteString(fmt.Sprintf("%d", len(res)))
		for i := 0; i < len(res); i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func solve(n int, A []int) []int {
	res := make([]int, 0, n)

	var i, j = 0, n - 1

	for i < j {
		if A[i] > 0 {
			i++
			continue
		}
		//A[i] <= 0
		for j > i && A[j] < 0 {
			j--
			continue
		}
		if j == i {
			break
		}
		A[i], A[j] = A[j], A[i]
		//A[j] >= 0
		res = append(res, i+1)
		res = append(res, j+1)
		i++
		j--
	}
	sort.Ints(res)

	return res
}

func getMaxSum(n int, A []int) int64 {
	var sum int64
	best := sum

	for i := 0; i < n; i++ {
		sum += int64(A[i])
		if best < sum {
			best = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return best
}

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

	H := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	var pos int
	for i := 0; i < n; i++ {
		var x uint64
		pos = readUint64(s, pos, &x) + 1
		H[i] = int64(x)
	}

	res := solve(n, H)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')
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
func solve(n int, H []int64) []int64 {
	if n == 1 {
		return H
	}
	var sum int64
	for i := 0; i < n; i++ {
		sum += H[i]
	}
	// if (a + a + n - 1) * n / 2 == sum
	// (2 * a + n - 1) = 2 * sum / n
	N := int64(n)
	if 2*sum%N == 0 {
		m := 2 * sum / N
		m += 1 - N
		if m > 0 && m%2 == 0 {
			a := m / 2
			for i := 0; i < n; i++ {
				H[i] = a + int64(i)
			}
			return H
		}
	}

	// then there should be one adjacent pair that H[i] == H[i+1] in the result.
	// (a + a + n - 2) * (n - 1) / 2 + H[k] = sum
	// where k is the same pair index
	// H[k] = a + k
	// (2 * a + n - 2) * (n - 1) / 2 + a + k == sum
	// (2 * a + n - 2) * (n - 1) + 2 * a  + 2 * k == 2 * sum
	sum *= 2
	// 2 * a * (n - 1) + 2 * a + (n - 2) * (n - 1) + 2 * k == 2 * sum
	sum -= (N - 2) * (N - 1)
	// 2 * a * n + 2 * k == 2 * sum - (N - 2) * (N - 1)
	for k := 0; k < n-1; k++ {
		// if k == i
		tmp := sum - 2*int64(k)
		if tmp%(2*N) == 0 {
			// found
			a := tmp / (2 * N)
			//a := m / 2
			for i := 0; i < n; i++ {
				H[i] = a + int64(i)
				if i >= k + 1 {
					H[i]--
				}
			}

			break
		}
	}
	return H
}

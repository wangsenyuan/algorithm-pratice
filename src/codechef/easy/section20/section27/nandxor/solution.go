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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int) []int {
	n := len(A)

	if n >= 62 {
		return solve1(A)
	}
	return solve2(A)
}

func solve2(A []int) []int {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := A[i] ^ A[j]
			for k := i + 1; k < n; k++ {
				if k == j {
					continue
				}
				for l := k + 1; l < n; l++ {
					if l == j {
						continue
					}
					y := A[k] ^ A[l]
					if popcount(x) == popcount(y) {
						return []int{i + 1, j + 1, k + 1, l + 1}
					}
				}
			}
		}
	}
	return nil
}

func solve1(A []int) []int {
	found := make([]int, 30)
	for i := 0; i < 30; i++ {
		found[i] = -1
	}
	for i := 0; ; i += 2 {
		x := A[i] ^ A[i+1]
		c := popcount(x)
		if found[c] >= 0 {
			return []int{found[c] + 1, found[c] + 2, i + 1, i + 2}
		}
		found[c] = i
	}
}

func popcount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

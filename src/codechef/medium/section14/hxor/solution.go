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
	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		solve(n, k, A)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", A[i]))
		}
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

const H = 30

func solve(n int, k int, A []int) {
	bits := make([][]int, H)
	for i := 0; i < H; i++ {
		bits[i] = make([]int, 0, 10)
		for j := 0; j < n; j++ {
			num := A[j]
			if (num>>uint(i))&1 == 1 {
				bits[i] = append(bits[i], j)
			}
		}
	}

	pos := make([]int, H)

	for i := 0; i < n-1 && k > 0; i++ {
		//num := A[i]
		for j := H - 1; j >= 0 && k > 0; j-- {
			if (A[i]>>uint(j))&1 == 1 {
				// this position is set
				a := bits[j][pos[j]]
				pos[j]++
				b := n - 1
				if pos[j] < len(bits[j]) {
					b = bits[j][pos[j]]
					pos[j]++
				}
				//a should be i
				A[a] ^= 1 << uint(j)
				A[b] ^= 1 << uint(j)
				k--
			}
		}
	}

	if n == 2 && k&1 == 1 {
		// has to set one
		A[0] ^= 1
		A[1] ^= 1
	}
}

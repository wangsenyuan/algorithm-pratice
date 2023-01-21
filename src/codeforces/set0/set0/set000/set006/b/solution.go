package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	s, _ := reader.ReadBytes('\n')
	pos := readInt(s, 0, &n)
	pos = readInt(s, pos+1, &m)
	P := s[pos+1]

	R := make([]string, n)

	for i := 0; i < n; i++ {
		R[i], _ = reader.ReadString('\n')
	}

	fmt.Println(solve(n, m, P, R))
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(n, m int, P byte, R []string) int {
	mem := make(map[byte]bool)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if R[i][j] == P {
				for k := 0; k < 4; k++ {
					u, v := i+dd[k], j+dd[k+1]
					if u >= 0 && u < n && v >= 0 && v < m && R[u][v] != P && R[u][v] != '.' {
						mem[R[u][v]] = true
					}
				}
			}
		}
	}

	return len(mem)
}

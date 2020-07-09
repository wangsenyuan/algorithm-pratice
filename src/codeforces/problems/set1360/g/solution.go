package main

import (
	"bufio"
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

	for tc > 0 {
		tc--
		line := readNNums(reader, 4)
		n, m, a, b := line[0], line[1], line[2], line[3]
		ok, g := solve(n, m, a, b)
		if ok {
			fmt.Println("YES")
			for i := 0; i < n; i++ {
				fmt.Println(string(g[i]))
			}
		} else {
			fmt.Println("NO")
		}
	}
}

func solve(n, m, a, b int) (bool, [][]byte) {
	if n*a != b*m {
		return false, nil
	}

	G := make([][]byte, n)
	var d = 1
	for (d*n)%m != 0 {
		d++
	}

	var j int
	for i := 0; i < n; i++ {
		G[i] = make([]byte, m)
		var x int
		for k := 0; k < m; k++ {
			if x < a {
				G[i][(j+k)%m] = '1'
				x++
			} else {
				G[i][(j+k)%m] = '0'
			}
		}
		j = (j + d) % m
	}

	return true, G
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)
	}

	res := solve(n, mat)

	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(n int, mat []string) int {
	grid := make([][]uint8, n)

	for i := 0; i < n; i++ {
		// 4位一个，4个一组
		grid[i] = make([]uint8, n/4)
		for j := 0; j < n/4; j++ {
			if mat[i][j] <= '9' && mat[i][j] >= '0' {
				grid[i][j] = uint8(mat[i][j] - '0')
			} else {
				grid[i][j] = 10 + uint8(mat[i][j]-'A')
			}
		}
	}

	get := func(r int, j int) uint8 {
		c, k := j/4, j%4
		val := grid[r][c]
		return (val >> (3 - k)) & 1
	}
	sum := make([][]uint32, n)
	for i := 0; i < n; i++ {
		sum[i] = make([]uint32, n)
		for j := 0; j < n; j++ {
			sum[i][j] = uint32(get(i, j))
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	check := func(x int) bool {
		for i := 0; i < n; i += x {
			for j := 0; j < n; j += x {
				tmp := sum[i+x-1][j+x-1]
				if i > 0 {
					tmp -= sum[i-1][j+x-1]
				}
				if j > 0 {
					tmp -= sum[i+x-1][j-1]
				}
				if i > 0 && j > 0 {
					tmp += sum[i-1][j-1]
				}
				if tmp != uint32(x*x) && tmp != 0 {
					return false
				}
			}
		}
		return true
	}

	for x := n; x > 1; x-- {
		if n%x == 0 && check(x) {
			return x
		}
	}

	return 1
}

func solve1(n int, mat []string) int {
	grid := make([][]uint8, n)

	for i := 0; i < n; i++ {
		// 4位一个，4个一组
		grid[i] = make([]uint8, n/4)
		for j := 0; j < n/4; j++ {
			if mat[i][j] <= '9' && mat[i][j] >= '0' {
				grid[i][j] = uint8(mat[i][j] - '0')
			} else {
				grid[i][j] = 10 + uint8(mat[i][j]-'A')
			}
		}
	}

	get := func(r int, j int) uint8 {
		c, k := j/4, j%4
		val := grid[r][c]
		return (val >> (3 - k)) & 1
	}

	x := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; {
			k := j
			expect := get(i, k)
			for j < n && get(i, j) == expect {
				j++
			}
			x = gcd(x, j-k)
		}
	}

	for j := 0; j < n; j++ {
		for i := 0; i < n; {
			k := i
			expect := get(k, j)
			for i < n && get(i, j) == expect {
				i++
			}
			x = gcd(x, i-k)
		}
	}

	return x
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

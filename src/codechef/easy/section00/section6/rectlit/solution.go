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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		k, n := readTwoNums(scanner)

		if k >= 4 {
			for i := 0; i < k; i++ {
				scanner.Scan()
			}
			fmt.Println("yes")
		} else {
			coords := make([][]int, k)
			for i := 0; i < k; i++ {
				coords[i] = readNNums(scanner, 2)
			}
			res := solve(k, n, coords)
			if res {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}

	}
}

func solve(k int, n int, coords [][]int) bool {
	if k >= 4 {
		return true
	}

	if k == 0 {
		return false
	}
	all := 1
	for i := 0; i < k; i++ {
		all *= 4
	}

	rect := make([][]int, k)

	N := int64(n)

	for mask := 0; mask < all; mask++ {
		x := mask
		for i := 0; i < k; i++ {
			xx := x % 4
			var p1 int
			if xx/2 == 1 {
				p1 = n - 1
			}
			var p2 int
			if xx%2 == 1 {
				p2 = n - 1
			}
			rect[i] = []int{min(p1, coords[i][0]), min(p2, coords[i][1]), max(p1, coords[i][0]), max(p2, coords[i][1])}

			x /= 4
		}

		var area int64

		for mask2 := 1; mask2 < (1 << uint(k)); mask2++ {
			minx, miny := 0, 0
			maxx, maxy := n-1, n-1
			for j := 0; j < k; j++ {
				if mask2&(1<<uint(j)) > 0 {
					minx = max(minx, rect[j][0])
					miny = max(miny, rect[j][1])
					maxx = min(maxx, rect[j][2])
					maxy = min(maxy, rect[j][3])
				}
			}
			tmp := int64(maxx-minx+1) * int64(maxy-miny+1)

			if tmp < 0 {
				tmp = 0
			}

			if bitCount(mask2)%2 == 0 {
				area -= tmp
			} else {
				area += tmp
			}
		}
		if area == N*N {
			return true
		}
	}

	return false
}

func bitCount(num int) int {
	var res int
	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

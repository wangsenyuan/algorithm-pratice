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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var tmp int64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNInt64Nums(scanner *bufio.Scanner, n int) []int64 {
	res := make([]int64, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt64(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		r0, c0 := readTwoNums(scanner)

		steps := solve(r0, c0)

		fmt.Println(len(steps))

		for i := 0; i < len(steps); i++ {
			fmt.Printf("%d %d\n", steps[i][0], steps[i][1])
		}
	}
}

func solve(r0, c0 int) [][]int {

	steps := make([][]int, 0, 64)
	// any way move r0, c0 to main diagonal
	if r0+c0 == 8 {
		// already in bottom main diagonal
		if r0 != 7 {
			// move there first
			r0, c0 = 7, 1
			steps = append(steps, []int{r0, c0})
		}
	} else if r0+c0 == 10 {
		// already in up main diagonal
		if r0 != 8 {
			r0, c0 = 8, 2
			steps = append(steps, []int{r0, c0})
		}
	} else {
		// rx - cx = r0 - c0
		for rx := 1; rx <= 7; rx++ {
			cx := 8 - rx
			if rx-cx == r0-c0 {
				steps = append(steps, []int{rx, cx})
				if rx != 7 {
					steps = append(steps, []int{7, 1})
					rx, cx = 7, 1
				}
				r0, c0 = rx, cx
				break
			}
		}
	}

	// now r0, c0, is at 7, 1, or 8, 2
	for i := 1; i <= 7; i++ {
		x, y := r0, c0
		var xx, yy int

		if x+y == 8 {
			// in bottom diag, go down first, then to right-most, then down
			xx = x - 1
			for xx > 0 {
				yy = y + xx - x
				if yy <= 0 || yy > 8 {
					break
				}
				steps = append(steps, []int{xx, yy})
				xx--
			}
			xx = min(8, 8-y+x)
			for xx > x {
				yy = y + xx - x
				if yy <= 0 || yy > 8 {
					break
				}
				steps = append(steps, []int{xx, yy})
				xx--
			}
			x++
			y++
			r0 = x - 1
			c0 = 10 - r0
			if r0 > 0 && r0 <= 8 && c0 > 0 && c0 <= 8 {
				steps = append(steps, []int{r0, c0})
			}
		} else {
			// in up diag, go up first, then go bottom-most, then up
			xx = x + 1
			for xx <= 8 {
				yy = y + xx - x
				if yy <= 0 || yy > 8 {
					break
				}
				steps = append(steps, []int{xx, yy})
				xx++
			}
			xx = max(1, 1-y+x)
			for xx < x {
				yy = y + xx - x
				if yy <= 0 || yy > 8 {
					break
				}
				steps = append(steps, []int{xx, yy})
				xx++
			}
			x--
			y--
			r0 = x - 1
			c0 = 8 - r0
			if r0 > 0 && r0 <= 8 && c0 > 0 && c0 <= 8 {
				steps = append(steps, []int{r0, c0})
			}
		}
	}

	return steps
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

package main

import (
	"bufio"
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
		var A, M uint64
		scanner.Scan()
		pos := readUint64(scanner.Bytes(), 0, &A)
		readUint64(scanner.Bytes(), pos+1, &M)
		res := solve(int64(A), int64(M))

		fmt.Println(len(res))
		if len(res) == 0 {
			fmt.Println()
		} else {
			for i := 0; i < len(res); i++ {
				if i < len(res)-1 {
					fmt.Printf("%d ", res[i])
				} else {
					fmt.Printf("%d\n", res[i])
				}
			}
		}
	}
}

func solve(A, M int64) []int64 {
	// s := int64(math.Sqrt(float64(M)))
	var res []int64
	for d := int64(1); d*d <= M; d++ {
		if M%d == 0 {
			x := M / d
			if x > 1 && (x-1)%A == 0 {
				p := (x - 1) / A
				res = append(res, d*p)
			}
			if d > 1 && (d-1)%A == 0 {
				p := (d - 1) / A
				x := M / d
				res = append(res, p*x)
			}
		}
	}

	sort.Sort(Nums(res))

	var j int

	for i := 1; i <= len(res); i++ {
		if i == len(res) || res[i] > res[i-1] {
			res[j] = res[i-1]
			j++
		}
	}

	return res[:j]
}

type Nums []int64

func (nums Nums) Len() int {
	return len(nums)
}

func (nums Nums) Less(i, j int) bool {
	return nums[i] < nums[j]
}

func (nums Nums) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

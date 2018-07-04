package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	n := readNum(scanner)

	monsters := readNNums(scanner, n)

	q := readNum(scanner)

	actions := make([][]int, q)

	for i := 0; i < q; i++ {
		actions[i] = readNNums(scanner, 2)
	}
	res := solve(n, monsters, q, actions)

	for _, ans := range res {
		fmt.Println(ans)
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
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

const INF = 1e9 + 3
const MASK = 1<<17 - 1
const Q_MAX = 1<<18 + 1
const BIT = 18

var bitMask = make([]int, Q_MAX)
var zeros = make([]int, Q_MAX)

func solve(n int, monsters []int, q int, actions [][]int) []int {

	for i := 0; i < q; i++ {
		// If the value is greater than
		// 1<<17 trim it. As index can
		// be at max 1<<17
		actions[i][0] &= MASK
	}

	blockSize := 1000
	totalBlock := (q + blockSize - 1) / blockSize
	index := 0
	dead := make([]int, q)
	for blockIndex := 0; blockIndex < totalBlock; blockIndex++ {
		copy(bitMask, zeros)

		duplicate := index

		for i := 0; i < blockSize && index < q; i, index = i+1, index+1 {
			bitMask[actions[index][0]] += actions[index][1]
		}

		for bit := 0; bit < BIT; bit++ {
			for i := MASK; i >= 0; i-- {
				if i&(1<<uint(bit)) == 0 {
					bitMask[i] += bitMask[i^(1<<uint(bit))]
					if bitMask[i] > INF {
						bitMask[i] = INF
					}
				}
			}
		}
		pos := duplicate

		for i := 0; i < n; i++ {
			if monsters[i] <= 0 {
				continue
			}
			monsters[i] -= bitMask[i]
			if monsters[i] > 0 {
				// survive
				continue
			}
			cur := 0
			duplicate = pos
			for j := 0; j < blockSize && duplicate < q; j, duplicate = j+1, duplicate+1 {
				if i&actions[duplicate][0] == i {
					cur += actions[duplicate][1]
				}
				// get back health: bitMask[i] + monsters[i]
				if bitMask[i]+monsters[i] <= cur {
					dead[duplicate]++
					break
				}
			}
		}
	}
	for i := 1; i < q; i++ {
		dead[i] += dead[i-1]
	}
	res := make([]int, q)

	for i := 0; i < q; i++ {
		res[i] = n - dead[i]
	}
	return res
}

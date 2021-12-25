package main

import (
	"bufio"
	"fmt"
	"math"
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

	input := make([]int, 4)
	for {
		fillNNums(scanner, 4, input)
		if input[0] == 0 && input[1] == 0 && input[2] == 0 && input[3] == 0 {
			break
		}
		res := solve(input[0], input[1], input[2], input[3])
		fmt.Printf("%.3f\n", res)
	}
}

func solve(dh, dl, dr, k int) float64 {
	const e = 1e-5

	var ans float64
	fk := float64(k)
	fh := float64(dh)
	fl := float64(dl)
	fr := float64(dr)

	low := 0.0
	high := fr - e

	for i := 0; i < 100; i++ {
		mid := (low + high) / 2
		mide := mid + e

		first := fk*math.Sqrt(fl*fl+mid*mid) + math.Sqrt(fh*fh+(fr-mid)*(fr-mid))
		second := fk*math.Sqrt(fl*fl+mide*mide) + math.Sqrt(fh*fh+(fr-mide)*(fr-mide))

		if first < second {
			ans = math.Sqrt(fl*fl+mid*mid) + math.Sqrt(fh*fh+(fr-mid)*(fr-mid))
		} else {
			ans = math.Sqrt(fl*fl+mide*mide) + math.Sqrt(fh*fh+(fr-mide)*(fr-mide))
		}

		if first < second {
			high = mid
		} else {
			low = mid
		}
	}

	return ans
}

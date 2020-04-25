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
		n := readNum(reader)
		res := solve(n)
		fmt.Println(len(res))
		for i := 0; i < len(res); i++ {
			fmt.Printf("%d ", len(res[i]))
			for j := 0; j < len(res[i]); j++ {
				if j < len(res[i])-1 {
					fmt.Printf("%d ", res[i][j])
				} else {
					fmt.Printf("%d\n", res[i][j])
				}
			}
		}
	}
}

func solve(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	res := make([][]int, n/2)

	if n%2 == 0 {
		for i := 0; i < n; i += 2 {
			res[i/2] = []int{i + 1, i + 2}
		}
	} else {
		res[0] = []int{1, 2, n}

		for i := 1; i < n-2; i += 2 {
			res[i/2+1] = []int{i + 2, i + 3}
		}
	}

	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(reader)
	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
	//buf.WriteString(fmt.Sprintf("%d\n", res))
	if len(res) > 0 {
		fmt.Println("YES")
		fmt.Printf("%d %d %d %d\n", res[0], res[1], res[2], res[3])
	} else {
		fmt.Println("NO")
	}
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

func solve(A []int) []int {
	var a, b int
	for _, num := range A {
		if num > a {
			b = a
			a = num
		} else if num > b {
			b = num
		}
	}

	S := a + b
	x := make([]int, S+1)
	y := make([]int, S+1)
	for i := 0; i <= S; i++ {
		x[i] = -1
		y[i] = -1
	}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			sum := A[i] + A[j]
			if x[sum] >= 0 && x[sum] != i && y[sum] != j && y[sum] != i {
				return []int{x[sum] + 1, y[sum] + 1, i + 1, j + 1}
			}
			x[sum] = i
			y[sum] = j
		}
	}
	return nil
}

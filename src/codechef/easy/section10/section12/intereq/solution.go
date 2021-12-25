package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, _ := readTwoNums(reader)
		ask := func(arr []int, x int) int {
			var buf bytes.Buffer
			buf.WriteString(fmt.Sprintf("? %d", len(arr)+1))
			for i := 0; i < len(arr); i++ {
				buf.WriteString(fmt.Sprintf(" %d", arr[i]))
			}
			buf.WriteString(fmt.Sprintf(" %d", x))
			fmt.Println(buf.String())
			return readNum(reader)
		}
		A := solve(n, ask)

		var buf bytes.Buffer
		buf.WriteString("!")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf(" %d", A[i]))
		}
		fmt.Println(buf.String())

		readNum(reader)
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

func solve(n int, ask func([]int, int) int) []int {
	A := make([]int, n)
	A[0] = 1
	arr := make([]int, 0, 101)
	arr = append(arr, 1)
	cur := 1
	query := func(i int) int {
		if ask(arr, i) == 1 {
			// a new element
			cur++
			arr = append(arr, i)
			return cur
		}
		l, r := 0, len(arr)-1
		for l < r {
			mid := (l + r) / 2
			if ask(arr[l:mid+1], i) == 1 {
				l = mid + 1
			} else {
				r = mid
			}
		}
		j := arr[l]
		return A[j-1]
	}

	for i := 1; i < n; i++ {
		A[i] = query(i + 1)
	}
	return A
}

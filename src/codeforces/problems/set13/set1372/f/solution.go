package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	ask := func(l, r int) (int, int) {
		fmt.Printf("? %d %d\n", l, r)
		return readTwoNums(reader)
	}

	res := solve(n, ask)
	var buf bytes.Buffer
	buf.WriteString("!")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	fmt.Println(buf.String())
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

func solve(n int, ask func(l int, r int) (int, int)) []int {
	arr := make([]int, n)
	occ := make(map[int]int)

	var query func(l, r int) int

	query = func(l, r int) int {
		if l <= r {
			x, f := ask(l, r)

			if ff, found := occ[x]; found {
				end := r - f + ff
				for j := end - ff + 1; j <= end; j++ {
					arr[j-1] = x
				}
				delete(occ, x)
				query(l, r-f)
				return end
			}
			occ[x] = f
			var j = l
			for occ[x] > 0 {
				j = query(j, j+f-1) + 1
			}
			return query(j, r)
		}
		return l - 1
	}

	query(1, n)

	return arr
}

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

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		nums := readNInt64s(reader, 2)
		A, B := nums[0], nums[1]
		res := solve(A, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

var nums map[int64]bool

func init() {
	nums = map[int64]bool{
		125874:     true,
		1025874:    true,
		1245789:    true,
		10025874:   true,
		10245789:   true,
		12356784:   true,
		12457899:   true,
		100025874:  true,
		100245789:  true,
		102356784:  true,
		102457899:  true,
		123456789:  true,
		124578999:  true,
		1000025874: true,
		1000245789: true,
		1002356784: true,
		1002457899: true,
		1023456789: true,
		1024578999: true,
		1122556284: true,
		1233566784: true,
		1234567899: true,
		1245789999: true,
		1573788744: true,
	}
}

func solve(A int64, B int64) int {
	var res int
	for num := range nums {
		res += find(A, B, num)
	}

	return res
}

func find(L, U int64, num int64) int {
	var arr []int
	for num > 0 {
		arr = append(arr, int(num%10))
		num /= 10
	}

	l := len(arr)

	mask := (1 << l) - 1

	res := make(map[int64]bool)

	var dfs func(cur int64, flag int)

	dfs = func(cur int64, flag int) {
		if cur > U {
			return
		}
		if flag == mask {
			if cur >= L {
				res[cur] = true
			}
			return
		}

		for i := 0; i < l; i++ {
			if (flag>>uint64(i))&1 == 0 && (arr[i] != 0 || flag > 0) {
				dfs(cur*10+int64(arr[i]), flag|(1<<uint64(i)))
			}
		}

	}

	dfs(0, 0)

	return len(res)
}

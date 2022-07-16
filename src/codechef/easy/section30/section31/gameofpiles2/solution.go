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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
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

const CHEF = "CHEF"
const CHEFINA = "CHEFINA"

func solve(A []int) string {
	cnt := make([]int, 3)
	n := len(A)
	var sum int64
	for i := 0; i < n; i++ {
		if A[i] < 3 {
			cnt[A[i]]++
		}
		sum += int64(A[i])
	}
	if cnt[1] >= 2 {
		sum -= int64(n)
		if sum&1 == 1 {
			return CHEF
		}
		return CHEFINA
	}
	if cnt[1] == 1 {
		if sum&1 == 1 {
			return CHEF
		}
		// sum is even
		if cnt[2] > 0 && (sum-int64(n))&1 == 1 {
			return CHEF
		}

		return CHEFINA
	}
	// 如果有一个pile是奇数，chef从上面取走一个后，
	// 减去chef取走的1个，剩余的部分是偶数，所以chef总可以第二个取空pile；
	// chefina选择偶数的pile，chef也选择；
	// chefina选择奇数的pile，车费也选择；
	if sum&1 == 1 {
		return CHEF
	}

	return CHEFINA
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

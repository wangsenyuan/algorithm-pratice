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
		S, _ := reader.ReadString('\n')
		res := solve(S)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(S string) int64 {
	var cnt int
	for i := 0; i < len(S); i++ {
		if S[i] >= 'a' && S[i] <= 'z' {
			cnt ^= 1 << uint(S[i]-'a')
		}
	}
	mem := make([]map[int]int, 2)
	mem[0] = make(map[int]int)
	mem[1] = make(map[int]int)

	mem[0][0] = 1

	var res int64

	var cnt2 int
	var qm int
	for i := 0; i < len(S); i++ {
		if S[i] >= 'a' && S[i] <= 'z' {
			cnt2 ^= 1 << uint(S[i]-'a')
		} else if S[i] == '?' {
			qm ^= 1
		} else {
			break
		}

		cnt1 := cnt ^ cnt2

		res += int64(mem[qm][cnt1])

		for j := 0; j < 26; j++ {
			res += int64(mem[qm^1][cnt1^(1<<uint(j))])
		}

		mem[qm][cnt2]++
	}

	return res
}

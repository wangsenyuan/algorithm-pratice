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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func fillNNums(bs []byte, n int, res []int) {
	x := 0
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

const MOD = 1000000007

func solve(n int, A []int) int {
	N := 1 << uint(n)

	mem := make([][][]int, 31)

	for i := 0; i < 31; i++ {
		mem[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			mem[i][j] = make([]int, N)
			for k := 0; k < N; k++ {
				mem[i][j][k] = -1
			}
		}
	}

	var count func(bit int, mask1 int, mask2 int) int

	count = func(bit int, mask1 int, mask2 int) int {
		if bit < 0 {
			return 1
		}

		if mem[bit][mask1][mask2] >= 0 {
			return mem[bit][mask1][mask2]
		}

		var res int
		// cur is the numbers set on each B number'is bit position
		for cur := 0; cur < 1<<uint(n); cur++ {
			nmask1, nmask2 := mask1, mask2
			var ok = true
			for i := 0; i < n-1; i++ {
				b1 := (cur >> uint(i)) & 1
				b2 := (cur >> uint(i+1)) & 1
				if b1 == 1 && b2 == 0 && (mask1>>uint(i))&1 == 0 {
					// can'set B[i] as 1, and B[i+1] = 0, when B[i] == B[i+1]
					ok = false
					break
				}
				a1 := ((A[i] >> uint(bit)) & 1) ^ b1
				a2 := ((A[i+1] >> uint(bit)) & 1) ^ b2
				if a1 == 1 && a2 == 0 && (mask2>>uint(i))&1 == 0 {
					ok = false
					break
				}

				if b1 == 0 && b2 == 1 {
					nmask1 |= 1 << uint(i)
				}
				if a1 == 0 && a2 == 1 {
					nmask2 |= 1 << uint(i)
				}
			}
			if ok {
				res += count(bit-1, nmask1, nmask2)
				if res >= MOD {
					res -= MOD
				}
			}
		}
		mem[bit][mask1][mask2] = res
		return res
	}

	return count(30, 0, 0)
}

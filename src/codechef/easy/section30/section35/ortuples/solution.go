package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		P, Q, R := readThreeNums(reader)
		res := solve(P, Q, R)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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
		if s[i] == '\n' || s[i] == '\r' {
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

const H = 20

func solve(P, Q, R int) int64 {

	var res int64 = 1

	for i := 0; i < H; i++ {
		x := (P >> i) & 1
		y := (Q >> i) & 1
		z := (R >> i) & 1
		cnt := x + y + z
		if cnt == 0 {
			// only one options, all be zero
			continue
		}
		if cnt == 1 {
			return 0
		}
		if cnt == 2 {
			// a, b, or c be zero
			continue
		}
		res *= 4
	}

	return res
}

func solve1(P, Q, R int) int64 {
	// max(P, Q, R) < 1 << 20
	// A | B = P, B | C = Q, C | A = R
	// A, B should be subset of P
	// B, C should be subset of Q
	// A, C should be subset of R
	nums := []int{P, Q, R}
	sort.Ints(nums)
	P, Q, R = nums[0], nums[1], nums[2]

	// B := P ^ A,
	// C ：= R ^ A
	var res int64

	for a := P; ; a = (a - 1) & P {
		if a&R == a {
			// b 至少要包含所有P中为1，且a中不为1的部分，
			b := P ^ a
			// c 至少要包含所有R中为1，且a中不为1的部分，
			c := R ^ a
			x := b | c
			if x&Q == x {
				// x 是 Q的一个子集
				var tmp int64 = 1
				for i := H - 1; i >= 0; i-- {
					if (Q>>i)&1 == 1 {
						// 如果Q[i] = 1
						if (x>>i)&1 == 0 {
							// 这一位需要为1，要么在b中为1， 要么在c中为1
							pw := 1
							if (P>>i)&1 == 1 {
								// b中可以为1
								pw *= 2
							}
							if (R>>i)&1 == 1 {
								// c 中可以为1
								pw *= 2
							}
							// pow(2, cnt) - 1
							tmp *= int64(pw - 1)
						}
						//else 这一位已经为1，（b中为1，or c = 1)
					}
					// else Q[i] = 0, then x[i] = 0, and force b[i] = 0, c[i] = 0
				}
				res += tmp
			}
		}
		if a == 0 {
			break
		}
	}

	return res
}

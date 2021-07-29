package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n)
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

const H = 30

func solve(n int) int {
	N := int64(n)
	B := make([]int, H)
	dp1 := make([][]int64, H)
	for i := 0; i < H; i++ {
		dp1[i] = make([]int64, 2)
	}

	dp2 := make([][][]int64, H)

	for i := 0; i < H; i++ {
		dp2[i] = make([][]int64, 2)
		for j := 0; j < 2; j++ {
			dp2[i][j] = make([]int64, 3)
		}
	}

	check := func(k int) int64 {
		for i := 0; i < H; i++ {
			for j := 0; j < 2; j++ {
				dp1[i][j] = 0
				for u := 0; u < 3; u++ {
					dp2[i][j][u] = 0
				}
			}
		}
		var l int
		for ; k > 0; k >>= 1 {
			B[l] = k & 1
			l++
		}

		reverse(B[:l])
		dp1[l][0] = 1
		dp1[l][1] = 1

		for i := l - 1; i >= 0; i-- {
			// already less
			dp1[i][0] = 2 * dp1[i+1][0]

			if B[i] == 0 {
				// can only set 0
				dp1[i][1] = dp1[i+1][1]
			} else {
				dp1[i][1] = dp1[i+1][0] + dp1[i+1][1]
			}
		}
		for idx := l - 1; idx >= 0; idx-- {
			for tight := 0; tight < 2; tight++ {
				for parity := 0; parity < 3; parity++ {
					last := 1
					if tight == 1 {
						last = B[idx]
					}
					for i := 0; i <= last; i++ {
						newParity := parity
						if parity == 2 && i == 1 {
							newParity = idx % 2
						}
						var cur int
						if i == 1 {
							idxFromLeft := idx - newParity
							cur = 1 - (idxFromLeft & 1)
						}
						var newTight int
						if tight == 1 && i == last {
							newTight = 1
						}
						dp2[idx][tight][parity] += int64(cur)*dp1[idx+1][newTight] + dp2[idx+1][newTight][newParity]
					}
				}
			}
		}

		return dp2[0][1][2]
	}

	l, r := 1, n

	for l < r {
		mid := (l + r) / 2
		if check(mid) >= N {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m, h := readThreeNums(reader)
	S := readNNums(reader, m)
	res := solve(n, h, S)
	fmt.Printf("%.8f\n", res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(n int, h int, S []int) float64 {
	var sum int
	for i := 0; i < len(S); i++ {
		sum += S[i]
	}
	if sum < n {
		return -1
	}
	h--
	if S[h] == 1 {
		return 0
	}

	if sum-S[h] < n-1 {
		// always must select one from this room
		return 1.0
	}
	ans := big.NewFloat(1.0)
	for i := sum - 1; i > sum-S[h]; i-- {
		ans = ans.Quo(ans, big.NewFloat(float64(i)))
	}
	for i := sum - 1 - (n - 1); i > sum-S[h]-(n-1); i-- {
		ans = ans.Mul(ans, big.NewFloat(float64(i)))
	}

	//ans = ans.Sub(big.NewFloat(1.0), ans)
	//res := ans.String()
	//fmt.Println(res)
	res, _ := strconv.ParseFloat(ans.String(), 64)
	return 1.0 - res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
	long long ans = 1;
    int i;

    for(i = 1; i <= r; i++) {
        ans *= (n - r + i);
        ans /= i;
    }
*/

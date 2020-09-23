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
		S, _ := reader.ReadBytes('\n')
		ok := solve(n, S)
		if !ok {
			buf.WriteString("No\n")
		} else {
			buf.WriteString("Yes\n")
			buf.WriteString(string(S))
		}
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

const PAT = "abacaba"
const N = 100

var BUF []byte
var lps []int
var tmp []byte

func init() {
	BUF = make([]byte, N)
	lps = make([]int, N)
	tmp = make([]byte, len(PAT))
}

func solve(n int, S []byte) bool {
	// let's find whether there is already a abacaba already
	for i := 0; i+len(PAT) <= n; i++ {
		// if we make S[i:i + len(PAT)] as PAT
		var j int
		for j < len(PAT) && (S[i+j] == PAT[j] || S[i+j] == '?') {
			j++
		}
		if j < len(PAT) {
			continue
		}
		copy(tmp, S[i:i+j])
		copy(S[i:i+j], PAT)
		var cnt int
		p := kmp(n, S)
		for k := len(PAT) - 1; k < n; k++ {
			if p[k] == len(PAT) {
				cnt++
			}
		}
		if cnt == 1 {
			for k := 0; k < n; k++ {
				if S[k] == '?' {
					S[k] = 'd'
				}
			}
			return true
		}
		//revert
		copy(S[i:i+j], tmp)
	}

	return false
}

func kmp(n int, S []byte) []int {
	copy(BUF, PAT)
	BUF[len(PAT)] = '#'
	copy(BUF[len(PAT)+1:], S)
	n += len(PAT) + 1
	//lps := make([]int, n)
	for i := 1; i < n; i++ {
		j := lps[i-1]

		for j > 0 && BUF[i] != BUF[j] {
			j = lps[j-1]
		}

		lps[i] = j
		if BUF[i] == BUF[j] {
			lps[i]++
		}
	}

	return lps[len(PAT)+1 : n]
}

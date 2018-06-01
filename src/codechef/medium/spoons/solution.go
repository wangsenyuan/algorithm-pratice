package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const MAX_N = 65

var C [][]uint64

// const MAX_X = 1e18

func init() {
	C = make([][]uint64, MAX_N)
	C[0] = make([]uint64, MAX_N)
	C[0][0] = 1

	for i := 1; i < MAX_N; i++ {
		C[i] = make([]uint64, MAX_N)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j] + C[i-1][j-1]
		}
	}
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readUInt64(bytes []byte, from int, val *uint64) int {
	i := from
	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var t int
	readInt(scanner.Bytes(), 0, &t)

	var buf bytes.Buffer

	for t > 0 {
		var n uint64
		scanner.Scan()
		readUInt64(scanner.Bytes(), 0, &n)
		buf.WriteString(strconv.Itoa(solve(n)))
		buf.WriteByte('\n')
		t--
	}
	fmt.Print(buf.String())
}

func solve(n uint64) int {
	i := sort.Search(MAX_N, func(i int) bool {
		return C[i][i/2] >= n
	})
	return i
}

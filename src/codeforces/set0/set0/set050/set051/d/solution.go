package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	A := readNInt64s(reader, n)

	res := solve(A)

	fmt.Println(res)
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

func solve(A []int64) int {
	if progress(A) {
		return 0
	}

	n := len(A)

	if n == 2 {
		return 1
	}
	// A[0] != 0, A[1] != 0
	// d := A[1] / A[0]

	// first find some A[i] = 0, have to remove it
	var zero_count int
	zero_pos := -1
	for i := 0; i < n; i++ {
		if A[i] == 0 {
			zero_count++
			zero_pos = i
		}
	}
	buf := make([]int64, n-1)

	remove_and_check := func(i int) bool {
		copy(buf, A)
		copy(buf[i:], A[i+1:])
		return progress(buf)
	}

	if zero_count == n-1 {
		// remove the un-zero one
		return 1
	}
	if zero_count == n-2 {
		if A[0] != 0 {
			// remove either 0 or 1 could work
			// A[0] != 0, remove another 1
			return 1
		}
	}

	if zero_count > 1 {
		// 0 have to remove
		return 2
	}

	if zero_count == 1 {
		if remove_and_check(zero_pos) {
			return 1
		}
		return 2
	}
	// no zeros
	// deno = A[i+1] / A[i] all holds

	if remove_and_check(0) || remove_and_check(1) {
		return 1
	}

	// remove 0 & 1 will not work
	d1 := A[1] / A[0]

	if d1 != 0 && A[0]*d1 == A[1] {
		for i := 1; i+1 < n; i++ {
			d2 := A[i+1] / A[i]

			if d2 != d1 || d2*A[i] != A[i+1] {
				if remove_and_check(i) || remove_and_check(i+1) {
					return 1
				}
				break
			}
		}
	}

	d1 = A[0] / A[1]

	if d1 != 0 && d1*A[1] == A[0] {
		for i := 1; i+1 < n; i++ {
			d2 := A[i] / A[i+1]
			if d2 != d1 || d2*A[i+1] != A[i] {
				if remove_and_check(i) || remove_and_check(i+1) {
					return 1
				}
				break
			}
		}
	}

	return 2
}

func progress(arr []int64) bool {
	if len(arr) <= 1 {
		return true
	}
	if arr[0] == 0 {
		return all_zero(arr)
	}
	d := arr[1] / arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1]*d {

			if arr[1] != 0 {
				d = arr[0] / arr[1]
				for i := 1; i < len(arr); i++ {
					if arr[i-1] != arr[i]*d {
						return false
					}
				}
			} else {
				return false
			}
			break
		}
	}

	return true
}

func all_zero(arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

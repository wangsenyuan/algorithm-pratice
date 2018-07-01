package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N, Q := readTwoNums(scanner)
	queries := make([][]byte, Q)

	for i := 0; i < Q; i++ {
		scanner.Scan()
		queries[i] = scanner.Bytes()
	}
	res := solve(N, Q, queries)
	for _, ans := range res {
		fmt.Println(ans)
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

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func solve(N int, Q int, queries [][]byte) []int {
	RS0TS := make([]int, N) // RowSet i 0 timestamp
	RS1TS := make([]int, N)

	CS0TS := make([]int, N)
	CS1TS := make([]int, N)
	for i := 0; i < N; i++ {
		RS1TS[i] = -1
		CS1TS[i] = -1
	}
	RS0 := CreateBIT(Q)
	RS1 := CreateBIT(Q)
	CS0 := CreateBIT(Q)
	CS1 := CreateBIT(Q)

	rowSet0 := func(t, i int) {
		if RS0TS[i] > RS1TS[i] {
			if RS0TS[i] > 0 {
				//delete the previously stored time-stamp of RS0TS[i]
				RS0.Update(RS0TS[i], -1)
			}
		} else {
			RS1.Update(RS1TS[i], -1)
		}
		RS0.Update(t, 1)
		RS0TS[i] = t
	}

	rowSet1 := func(t, i int) {
		if RS0TS[i] > RS1TS[i] {
			RS0.Update(RS0TS[i], -1)
		} else {
			RS1.Update(RS1TS[i], -1)
		}
		RS1.Update(t, 1)
		RS1TS[i] = t
	}

	colSet0 := func(t, i int) {
		if CS0TS[i] > CS1TS[i] {
			CS0.Update(CS0TS[i], -1)
		} else {
			CS1.Update(CS1TS[i], -1)
		}
		CS0.Update(t, 1)
		CS0TS[i] = t
	}

	colSet1 := func(t, i int) {
		if CS0TS[i] > CS1TS[i] {
			CS0.Update(CS0TS[i], -1)
		} else {
			CS1.Update(CS1TS[i], -1)
		}
		CS1.Update(t, 1)
		CS1TS[i] = t
	}

	rowQuery := func(t, i int) int {
		if RS0TS[i] > RS1TS[i] {
			ones := CS1.Get(t) - CS1.Get(RS0TS[i])
			return N - ones
		}
		zeros := CS0.Get(t) - CS0.Get(RS1TS[i])
		return zeros
	}

	colQuery := func(t, i int) int {
		if CS0TS[i] > CS1TS[i] {
			return N - RS1.Get(t) + RS1.Get(CS0TS[i])
		}
		return RS0.Get(t) - RS0.Get(CS1TS[i])
	}

	res := make([]int, 0, Q)

	for t := 1; t <= Q; t++ {
		q := queries[t-1]
		var i, x, pos int
		if q[0] == 'R' && q[3] == 'S' {
			pos = readInt([]byte(q), 7, &i)
			pos = readInt([]byte(q), pos+1, &x)
			i--
			if x == 0 {
				rowSet0(t, i)
			} else {
				rowSet1(t, i)
			}
		} else if q[0] == 'C' && q[3] == 'S' {
			pos = readInt([]byte(q), 7, &i)
			pos = readInt([]byte(q), pos+1, &x)
			i--
			if x == 0 {
				colSet0(t, i)
			} else {
				colSet1(t, i)
			}
		} else if q[0] == 'R' && q[3] == 'Q' {
			readInt([]byte(q), 9, &i)
			i--
			tmp := rowQuery(t, i)
			res = append(res, tmp)
		} else {
			readInt([]byte(q), 9, &i)
			i--
			tmp := colQuery(t, i)
			res = append(res, tmp)
		}
	}

	return res
}

type BIT struct {
	arr  []int
	size int
}

func CreateBIT(n int) BIT {
	arr := make([]int, n+1)
	return BIT{arr, n}
}

func (bit *BIT) Update(pos int, val int) {
	if pos == 0 {
		return
	}
	arr := bit.arr
	size := bit.size
	for pos <= size {
		arr[pos] += val
		pos += pos & (-pos)
	}
}

func (bit *BIT) Get(pos int) int {
	arr := bit.arr
	var res int
	if pos > bit.size {
		return 0
	}
	for pos > 0 {
		res += arr[pos]
		pos -= pos & (-pos)
	}
	return res
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	T, N, H := readThreeNums(reader)

	query := func(ps []string) int {
		for _, s := range ps {
			fmt.Println(s)
		}
		return readNum(reader)
	}

	solve(T, N, H, query)
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

func solve(T int, N int, H int, query func([]string) int) {
	guess := make([]string, T/(N-H)+2)

	r := rand.New(rand.NewSource(99))

	for i := 0; i < len(guess); i++ {
		L := 8 + r.Intn(5)
		buf := make([]byte, L)
		for j := 0; j < L; j++ {
			buf[j] = byte('a' + r.Intn(26))
		}

		guess[i] = string(buf)
	}

	tested := make(map[Key]int)

	passwords := make([]string, N)
	var tc int
	score := func(arr []int) int {
		key := hashArrayAsKey(len(guess), arr)
		if v, ok := tested[key]; ok {
			return v
		}
		tc++
		for i := 0; i < N; i++ {
			passwords[i] = guess[arr[i]]
		}

		res := query(passwords)
		tested[key] = res
		return res
	}

	index := make([]int, N)
	honeyPot := make([]bool, N)
	best := make([]int, N)

	findHoney := func() {
		groupSize := 50

		m := (N + groupSize - 1) / groupSize
		// all use the first password to guess
		ref0 := score(index)

		var valid int

		for i := 0; i < m; i++ {
			// this group use the second password
			for j := i * groupSize; j < min((i+1)*groupSize, N); j++ {
				index[j] = 1
			}
			tmp := score(index)
			// change back
			for j := i * groupSize; j < min((i+1)*groupSize, N); j++ {
				index[j] = 0
			}
			if tmp == ref0 {
				// all honey pot
				for j := i * groupSize; j < min((i+1)*groupSize, N); j++ {
					honeyPot[j] = true
				}
			} else {
				// may have honey pot
				for j := i * groupSize; j < min((i+1)*groupSize, N); j++ {
					index[j] = 1
					// change only this one
					tmp = score(index)

					index[j] = 0

					if tmp == ref0 {
						honeyPot[j] = true
					} else {
						honeyPot[j] = false
						valid++
						if tmp > ref0 {
							best[j] = 1
						} else {
							best[j] = 0
						}
						if valid == H {
							for j < H {
								honeyPot[j] = true
								j++
							}
							return
						}
					}
				}
			}
			if valid == N-H {
				for j := (i + 1) * groupSize; j < N; j++ {
					honeyPot[j] = true
				}
				break
			}
		}
	}

	findHoney()

	copy(best, index)

	for i := 0; i < N; i++ {
		if honeyPot[i] {
			continue
		}
		for j := 2; j < len(guess); j++ {
			index[i] = j
			if score(index) > score(best) {
				best[i] = j
			}
		}
		index[i] = best[i]
	}

	for tc < T {
		tc++
		for i := 0; i < N; i++ {
			passwords[i] = "aaaaaaaaaa"
			query(passwords)
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

const MOD1 = 1e9 + 7
const MOD2 = 1e9 + 9
const MOD3 = 1e10 + 7

var mods = [3]int64{MOD1, MOD2, MOD3}

type Key struct {
	vals [3]int64
}

func NewKey(v int64) Key {
	var vals [3]int64
	for i := 0; i < len(vals); i++ {
		vals[i] = v % mods[i]
	}
	return Key{vals}
}

func (this Key) Mul(x int64) Key {
	var vals [3]int64
	for i := 0; i < len(vals); i++ {
		vals[i] = this.vals[i] * x % mods[i]
	}

	return Key{vals}
}

func (this Key) Add(x int64) Key {
	var vals [3]int64
	for i := 0; i < len(vals); i++ {
		vals[i] = (this.vals[i] + x) % mods[i]
	}

	return Key{vals}
}

func hashArrayAsKey(b int, arr []int) Key {
	var res Key
	for i := 0; i < len(arr); i++ {
		res = res.Mul(int64(b))
		res = res.Add(int64(arr[i]))
	}
	return res
}

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
		a, b := readTwoNums(reader)
		score, sets := solve(a, b)

		if len(score) == 0 {
			buf.WriteString("Impossible\n")
			continue
		}
		buf.WriteString(fmt.Sprintf("%d:%d\n", score[0], score[1]))
		for _, set := range sets {
			buf.WriteString(fmt.Sprintf("%d:%d ", set[0], set[1]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

const inf = 1 << 30

type pair struct {
	first  int
	second int
}

var f [5][5][201][201]*pair

func init() {

	// f[s][t][x][y]表示是否能够达到这个状态

	f[0][0][0][0] = &pair{-1, -1}

	for s := 0; s <= 4; s++ {
		for t := 0; t <= 4; t++ {
			if s+t >= 5 || s+t >= 4 && abs(s-t) == 2 || s+t == 3 && abs(s-t) == 3 {
				continue
			}

			id := s + t
			full := 25
			if id == 4 {
				full = 15
			}
			for x := 0; x <= 200; x++ {
				for y := 0; y <= 200; y++ {
					if f[s][t][x][y] == nil {
						continue
					}
					// 如果这场A获胜了
					if checkPossible(s+1, t) {
						for z := 0; z < full-2 && x+full <= 200 && y+z <= 200; z++ {
							f[s+1][t][x+full][y+z] = &pair{full, z}
						}
						// 差2分
						for z := full - 2; x+z+2 <= 200 && y+z <= 200; z++ {
							f[s+1][t][x+z+2][y+z] = &pair{z + 2, z}
						}
					}

					if checkPossible(s, t+1) {
						for z := 0; z < full-2 && x+z <= 200 && y+full <= 200; z++ {
							f[s][t+1][x+z][y+full] = &pair{z, full}
						}
						for z := full - 2; x+z <= 200 && y+z+2 <= 200; z++ {
							f[s][t+1][x+z][y+z+2] = &pair{z, z + 2}
						}
					}
				}
			}
		}
	}

}

func checkPossible(a, b int) bool {
	a, b = max(a, b), min(a, b)
	if a+b > 5 {
		return false
	}
	if a <= 3 {
		return true
	}
	// 4:1 not allowed
	return b >= 2
}

func abs(num int) int {
	return max(num, -num)
}

func solve(a int, b int) ([]int, [][]int) {

	var sets [][]int

	for s := 4; s >= 0; s-- {
		for t := 0; t <= 4; t++ {
			if (s >= 3 || t >= 3) && checkPossible(s, t) && f[s][t][a][b] != nil {
				score := []int{s, t}
				x, y := a, b
				for s >= 0 && t >= 0 && x >= 0 && y >= 0 && f[s][t][x][y] != nil {
					prev := f[s][t][x][y]
					if prev.first < 0 {
						break
					}

					sets = append(sets, []int{prev.first, prev.second})
					if prev.first > prev.second {
						s--
					} else {
						t--
					}
					x -= prev.first
					y -= prev.second
				}
				reverse(sets)

				return score, sets
			}
		}
	}

	return nil, nil
}

func reverse(arr [][]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

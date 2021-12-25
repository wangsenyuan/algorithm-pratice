package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for i := 0; i < t; i++ {
		n := readNum(scanner)
		hs, ms := make([]int, n), make([]int, n)
		for j := 0; j < n; j++ {
			tmp := readNNums(scanner, 2)
			hs[j] = tmp[0]
			ms[j] = tmp[1]
		}
		res := solve(n, hs, ms)
		fmt.Println(len(res))
		for _, ans := range res {
			outputAnswer("", ans[0])
			for k := 1; k < len(ans); k++ {
				outputAnswer(" ", ans[k])
			}
			fmt.Println()
		}
	}
}

func outputAnswer(space string, v int) {
	if len(space) > 0 {
		fmt.Print(space)
	}
	if v == INF {
		fmt.Print("Inf")
	} else {
		fmt.Printf("%d", v)
	}
}
func solve(n int, hs []int, ms []int) [][]int {
	firstHigh := process(n, hs, ms, true)
	firstLow := process(n, hs, ms, false)

	if len(firstHigh) == 0 && len(firstLow) == 0 {
		return nil
	}
	if len(firstHigh) == 0 {
		return [][]int{firstLow}
	}

	if len(firstLow) == 0 {
		return [][]int{firstHigh}
	}
	a, b := firstHigh[0], firstHigh[1]
	c, d := firstLow[0], firstLow[1]

	if b < c {
		return [][]int{{a, b}, {c, d}}
	}
	return [][]int{{min(a, c), max(b, d)}}
}

const INF = math.MaxInt32 - 1

func process(n int, hs []int, ms []int, firstHigh bool) []int {
	prevHigh := firstHigh
	begin, after := 0, INF
	for i := 1; i < n; i++ {
		from, end := INF, INF
		if ms[i-1] == ms[i] {
			if hs[i-1] == hs[i] {
				return nil
			}
			if hs[i-1] > hs[i] && prevHigh {
				from = 0
				end = INF
			} else if hs[i-1] < hs[i] && !prevHigh {
				from = 0
				end = INF
			} else {
				return nil
			}
		} else if ms[i-1] < ms[i] {
			if hs[i-1] == hs[i] {
				if prevHigh {
					return nil
				}
				from = 1
				end = INF
			} else if hs[i-1] < hs[i] {
				if prevHigh {
					return nil
				}
				from = 0
				end = INF
			} else {
				// hs[i-1] > hs[i]
				if prevHigh {
					from = 0
					end = floorDiv((hs[i-1] - hs[i]), (ms[i] - ms[i-1]))
				} else {
					from = ceilDiv(hs[i-1]-hs[i], ms[i]-ms[i-1])
					end = INF
				}
			}
		} else {
			// ms[i-1] > ms[i]
			if hs[i-1] == hs[i] {
				if !prevHigh {
					return nil
				}
				from = 1
				end = INF
			} else if hs[i-1] > hs[i] {
				if !prevHigh {
					return nil
				}
				from = 0
				end = INF
			} else {
				// hs[i-1] < hs[i]
				if prevHigh {
					from = ceilDiv(hs[i]-hs[i-1], ms[i-1]-ms[i])
					end = INF
				} else {
					from = 0
					end = floorDiv(hs[i]-hs[i-1], ms[i-1]-ms[i])
				}
			}
		}

		if from > after || end < begin {
			// no overlap
			return nil
		}
		begin = max(from, begin)
		after = min(end, after)
		prevHigh = !prevHigh
	}

	return []int{begin, after}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func ceilDiv(a, b int) int {
	c := a / b
	return c + 1
}

func floorDiv(a, b int) int {
	c := a / b
	if c*b < a {
		return c
	}
	return c - 1
}

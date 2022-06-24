package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		readString(reader)
		n, m := readTwoNums(reader)
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(reader, 3)
		}
		computers := make([][]int, m)
		for i := 0; i < m; i++ {
			computers[i] = readNNums(reader, 2)
		}
		res := solve(points, computers)

		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%.3f\n", res[i]))
		}
		buf.WriteByte('\n')

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

const CAPACITY = 32

func solve(router [][]int, computers [][]int) []float64 {
	S := len(router)
	mask := make([]int, S)
	for i := 0; i < S; i++ {
		mask[i] = 1 << uint(i)
	}

	C := len(computers)

	graph := make([][]int, S+1)
	for i := 0; i <= S; i++ {
		graph[i] = make([]int, C)
	}

	match := make([]int, C)

	for i := 0; i < C; i++ {
		match[i] = -1
	}

	smatch := make([][]int, S)
	for i := 0; i < S; i++ {
		smatch[i] = make([]int, CAPACITY)
	}

	aug := make([][]int, S+1)
	for i := 0; i <= S; i++ {
		aug[i] = make([]int, S+1)
	}

	count := make([]int, S)

	clist := make([][]int, S)

	for i := 0; i < S; i++ {
		clist[i] = make([]int, S*CAPACITY)
	}

	connected := make([]int, S)

	var dfs func(s int, mark *int) bool

	dfs = func(s int, mark *int) bool {
		*mark = ((*mark) | 1<<uint(s))
		if aug[s][S] > 0 {
			for match[clist[s][count[s]-1]] >= 0 {
				count[s]--
			}
			i := clist[s][count[s]-1]
			match[i] = s
			smatch[s][connected[s]] = i
			connected[s]++
			for j := 0; j < S; j++ {
				if graph[j][i] > 0 {
					aug[j][s]++
					aug[j][S]--
				}
			}
			return true
		}

		for i := 0; i < S; i++ {
			if ((*mark)>>uint(i))&1 == 0 && aug[s][i] > 0 {
				if dfs(i, mark) {
					for j := 0; j < connected[i]; j++ {
						if graph[s][smatch[i][j]] > 0 {
							c := smatch[i][j]
							for k := 0; k < S; k++ {
								if graph[k][c] > 0 {
									aug[k][s]++
									aug[k][i]--
								}
							}
							match[c] = s
							connected[i]--
							smatch[i][j] = smatch[i][connected[i]]
							smatch[s][connected[s]] = c
							connected[s]++
							return true
						}
					}
				}
			}
		}

		return false
	}

	addEdge := func(s, c int) bool {
		graph[s][c] = 1

		if match[c] < 0 {
			aug[s][S]++
			for i := 0; i < S; i++ {
				if (mask[i]>>(uint(s)))&1 == 1 {
					mask[i] |= 1 << uint(S)
				}
			}
			clist[s][count[s]] = c
			count[s]++
		} else {
			aug[s][match[c]]++
			if (mask[s]>>(uint(match[c])))&1 == 0 {
				for i := 0; i < S; i++ {
					if (mask[i]>>(uint(s)))&1 == 1 {
						mask[i] |= mask[match[c]]
					}
				}
			}
		}

		for t := 0; t < S; t++ {
			if connected[t] < router[t][2] && (mask[t]>>uint(S))&1 == 1 {
				dfs(t, new(int))
				for i := 0; i < S; i++ {
					mask[i] = 1 << uint(i)
					for j := 0; j <= S; j++ {
						if aug[i][j] > 0 {
							mask[i] |= 1 << uint(j)
						}
					}
				}
				for i := 0; i < S; i++ {
					for j := 0; j < S; j++ {
						if mask[j]&(1<<uint(i)) > 0 {
							mask[j] |= mask[i]
						}
					}
				}
				return true
			}
		}
		return false
	}

	ans := make([]float64, C)

	var latency int64

	edges := make(Edges, 0, S*C)

	for i, com := range computers {
		for j := 0; j < S; j++ {
			edge := Edge{distance(com, router[j]), j, i}
			heap.Push(&edges, edge)
		}

		for {
			top := heap.Pop(&edges).(Edge)
			latency = max(latency, top.dist)
			if addEdge(top.router, top.computer) {
				break
			}
		}
		ans[i] = math.Sqrt(float64(latency))
	}

	return ans
}

type Edge struct {
	dist     int64
	router   int
	computer int
}

func distance(a, b []int) int64 {
	dx := int64(a[0]) - int64(b[0])
	dy := int64(a[1]) - int64(b[1])
	return dx*dx + dy*dy
}

type Edges []Edge

func (edges Edges) Len() int {
	return len(edges)
}

func (edges Edges) Less(i, j int) bool {
	return edges[i].dist < edges[j].dist
}

func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func (edges *Edges) Push(x interface{}) {
	e := x.(Edge)
	*edges = append(*edges, e)
}

func (edges *Edges) Pop() interface{} {
	old := *edges
	n := len(old)
	ret := old[n-1]
	*edges = old[:n-1]
	return ret
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

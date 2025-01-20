package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	gates := make([]string, n)
	for i := range n {
		gates[i] = readString(reader)
	}
	return solve(n, gates)
}

func solve(n int, gates []string) string {

	res := make([]int, n)
	dep := make([][]int, n)

	deg := make([]int, n)
	fa := make([]int, n)
	fa[0] = -1
	for u, gate := range gates {
		if gate[0] == 'I' {
			// IN[i]
			readInt([]byte(gate), 3, &res[u])
		} else if gate[0] == 'N' {
			var v int
			readInt([]byte(gate), 4, &v)
			v--
			dep[u] = []int{v}
			fa[v] = u
			deg[u] = 1
		} else {
			var a, b int

			pos := 4
			if gate[0] == 'O' {
				pos = 3
			}

			pos = readInt([]byte(gate), pos, &a) + 1
			readInt([]byte(gate), pos, &b)
			a--
			b--
			dep[u] = []int{a, b}
			fa[a] = u
			fa[b] = u
			deg[u] = 2
		}
	}

	que := make([]int, n)
	var head, tail int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		if fa[u] >= 0 {
			v := fa[u]
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++

				if gates[v][0] == 'N' {
					// u是v的唯一子节点
					res[v] = res[u] ^ 1
				} else {
					a, b := dep[v][0], dep[v][1]
					if gates[v][0] == 'A' {
						res[v] = res[a] & res[b]
					} else if gates[v][0] == 'O' {
						res[v] = res[a] | res[b]
					} else {
						res[v] = res[a] ^ res[b]
					}
				}
			}
		}
	}

	head, tail = 0, 0
	marked := make([]bool, n)

	que[head] = 0
	head++
	for tail < head {
		u := que[tail]
		tail++
		gate := gates[u]
		if gate[0] == 'I' {
			marked[u] = true
			continue
		}
		if gate[0] == 'N' {
			v := dep[u][0]
			que[head] = v
			head++
			continue
		}
		a, b := dep[u][0], dep[u][1]
		if gate[0] == 'A' {
			if res[a] == 0 && res[b] == 0 {
				continue
			}
			if res[a] == 1 && res[b] == 1 {
				que[head] = a
				head++
				que[head] = b
				head++
			} else if res[b] == 1 {
				que[head] = a
				head++
			} else {
				que[head] = b
				head++
			}
		} else if gate[0] == 'O' {
			if res[a] == 1 && res[b] == 1 {
				continue
			}
			if res[a] == 0 && res[b] == 0 {
				que[head] = a
				head++
				que[head] = b
				head++
			} else if res[b] == 0 {
				// res[a] == 1
				que[head] = a
				head++
			} else {
				que[head] = b
				head++
			}
		} else {
			que[head] = a
			head++
			que[head] = b
			head++
		}
	}

	var ans []byte
	for i, cur := range gates {
		if cur[0] == 'I' {
			tmp := res[0]
			if marked[i] {
				tmp ^= 1
			}
			ans = append(ans, byte('0'+tmp))
		}
	}

	return string(ans)
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

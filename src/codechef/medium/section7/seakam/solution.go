package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, m := readTwoNums(scanner)
		edges := make([][]int, m)
		for i := 0; i < m; i++ {
			edges[i] = readNNums(scanner, 2)
		}
		fmt.Println(solve(n, edges))
	}
}

const MX = 100005

const MOD = 1000000007

var fact []int64
var ifact []int64

func init() {
	fact = make([]int64, MX)

	fact[0] = 1
	fact[1] = 1

	for i := 2; i < MX; i++ {
		fact[i] = (int64(i) * fact[i-1]) % MOD
	}

	ifact = make([]int64, MX)
	ifact[MX-1] = pow(int(fact[MX-1]), MOD-2)

	for i := MX - 2; i >= 0; i-- {
		ifact[i] = int64(i+1) * ifact[i+1]
		ifact[i] %= MOD
	}
}

func pow(a, b int) int64 {
	r := int64(1)
	A := int64(a)
	for b > 0 {
		if b&1 == 1 {
			r *= A
			r %= MOD
		}
		A *= A
		A %= MOD
		b >>= 1
	}

	return r
}

func nCr(n int, r int) int64 {
	if n < r {
		return 0
	}
	res := fact[n]
	res *= ifact[r]
	res %= MOD
	res *= ifact[n-r]
	res %= MOD
	return res
}

func getKey(u, v int) int64 {
	if u > v {
		u, v = v, u
	}
	return int64(u)*MX + int64(v)
}

func solve(n int, notConnectedEdges [][]int) int {
	if len(notConnectedEdges) == 0 {
		return int(fact[n])
	}

	faulty := make(map[int]bool)

	banned := make(map[int64]bool)

	for _, edge := range notConnectedEdges {
		u, v := edge[0]-1, edge[1]-1
		faulty[u] = true
		faulty[v] = true
		k := getKey(u, v)
		banned[k] = true
	}

	m := len(faulty)

	M := 1 << uint(m)

	dp := make([][][]int64, M)
	for i := 0; i < M; i++ {
		dp[i] = make([][]int64, m)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int64, m+1)
		}
	}

	for i := 0; i < m; i++ {
		mask := 1 << uint(i)
		dp[mask][i][0] = 1
	}

	nodes := make([]int, 0, m)

	for k := range faulty {
		nodes = append(nodes, k)
	}

	var ans int64

	for mask := 1; mask < M; mask++ {
		for first := 0; first < m; first++ {
			for tied := 0; tied < m; tied++ {
				if dp[mask][first][tied] > 0 {
					for second := 0; second < m; second++ {
						if mask&(1<<uint(second)) > 0 {
							continue
						}
						next := mask | 1<<uint(second)
						var edgeExists int
						if banned[getKey(nodes[first], nodes[second])] {
							edgeExists++
						}
						dp[next][second][tied+edgeExists] += dp[mask][first][tied]
						if dp[next][second][tied+edgeExists] >= MOD {
							dp[next][second][tied+edgeExists] -= MOD
						}
					}
				}

				if mask == M-1 {
					val := dp[mask][first][tied]
					val *= nCr(n-tied, m)
					val %= MOD
					val *= fact[n-m]
					val %= MOD

					ans += val
					if ans >= MOD {
						ans -= MOD
					}
				}

			}
		}
	}

	return int(ans)
}

package main

func main() {

}

const MAX_N = 1000000002

var f = make([]int, MAX_N)

func init() {
	f[1] = 1
	//1, 1
	f[2] = 1
	//1,1,1, 1, 2
	// if n % x == x - 1, then f[n] += f[x-1]
	// n = m * x + x - 1, n + 1 = x * (m + 1)
	for i := 2; i <= 100000; i++ {
		for j := 2 * i; j <= MAX_N; j += i {
			f[j-1] += f[i-1]
		}
	}
}

func solve(n int) int {
	return f[n] + 1
}

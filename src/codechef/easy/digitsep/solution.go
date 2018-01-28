package main

func main() {

}

func solve(N int, digits []byte, M int, X int, Y int) int64 {
	f := make([][][]int64, N)
	for i := 0; i < N; i++ {
		f[i] = make([][]int64, Y+2)
		for j := 0; j <= Y+1; j++ {
			f[i][j] = make([]int64, 0, 10)
		}
	}
	var num int64
	for i := 0; i < M && i < N; i++ {
		num = num*10 + int64(digits[i]-'0')
		f[i][1] = append(f[i][1], num)
	}

	for i := 0; i < N; i++ {
		for y := 1; y < Y+1; y++ {
			var x int64
			for k := i + 1; k < N && k <= i+M; k++ {
				x = x*10 + int64(digits[k]-'0')
				for _, p := range f[i][y] {
					tmp := gcd(x, p)
					f[k][y+1] = append(f[k][y+1], tmp)
				}
			}
		}
	}

	var ans int64
	for y := X + 1; y <= Y+1; y++ {
		for _, g := range f[N-1][y] {
			if g > ans {
				ans = g
			}
		}
	}

	return ans
}

func gcd(a, b int64) int64 {
	// fmt.Printf("[debug] gcd(%d, %d)", a, b)
	if a < 0 {
		panic("a should not be less than zero")
	}
	if b < 0 {
		panic("b should not be less than zero")
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

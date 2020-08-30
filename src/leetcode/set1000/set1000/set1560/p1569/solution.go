package p1569

const MOD = 1000000007
const N = 1001

var C [N][N]int64

func init() {
	C[0][0] = 1
	for i := 1; i < N; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = (C[i-1][j] + C[i-1][j-1]) % MOD
		}
	}
}

func numOfWays(nums []int) int {
	var loop func(arr []int) int64
	loop = func(arr []int) int64 {
		if len(arr) <= 1 {
			return 1
		}
		a := make([]int, 0, len(arr)/2)
		b := make([]int, 0, len(arr)/2)
		for i := 1; i < len(arr); i++ {
			if arr[i] > arr[0] {
				a = append(a, arr[i])
			} else {
				b = append(b, arr[i])
			}
		}
		res := (loop(a) * loop(b)) % MOD * C[len(arr)-1][len(a)]
		return res % MOD
	}
	res := loop(nums)
	res = (res - 1 + MOD) % MOD
	return int(res)
}

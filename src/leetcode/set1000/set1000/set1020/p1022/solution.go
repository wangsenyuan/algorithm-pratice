package p1022

func smallestRepunitDivByK(K int) int {
	vis := make([]bool, K)
	cur := 1
	l := 1
	for {
		rem := cur % K
		if rem == 0 {
			return l
		}
		if vis[rem] {
			return -1
		}
		vis[rem] = true
		cur = ((cur * 10) + 1) % K
		l++
	}
}

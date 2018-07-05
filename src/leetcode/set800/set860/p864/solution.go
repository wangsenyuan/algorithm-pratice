package p864

import (
	"math/rand"
)

type Solution struct {
	N      int
	avoids map[int]int
}

func Constructor(N int, blacklist []int) Solution {
	avoids := make(map[int]int)
	// sort.Ints(blacklist)
	for _, num := range blacklist {
		avoids[num] = N
	}
	M := N - len(avoids)
	for i := len(blacklist) - 1; i >= 0; i-- {
		num := blacklist[i]
		if num < M {
			_, found := avoids[N-1]
			for found {
				N--
				_, found = avoids[N-1]
			}
			avoids[num] = N - 1
			N--
		}
	}
	return Solution{N: M, avoids: avoids}
}

func (this *Solution) Pick() int {
	num := rand.Intn(this.N)
	if this.avoids[num] > 0 {
		return this.avoids[num]
	}
	return num
}

package p3435

import (
	"math"
	"math/bits"
)

func supersequences(words []string) (ans [][]int) {
	all := 0
	g := [26][]int{}
	for _, s := range words {
		x, y := int(s[0]-'a'), int(s[1]-'a')
		all |= 1<<x | 1<<y
		g[x] = append(g[x], y)
	}

	set := map[int]struct{}{}
	minSize := math.MaxInt
	// 枚举 all 的所有子集 sub
	for sub, ok := all, true; ok; ok = sub != all {
		size := bits.OnesCount(uint(sub))
		if size > minSize { // 剪枝
			sub = (sub - 1) & all
			continue
		}

		// 判断是否有环
		color := [26]int8{}
		var hasCycle func(int) bool
		hasCycle = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				// 只遍历不在 sub 中的字母
				if sub>>y&1 > 0 {
					continue
				}
				c := color[y]
				if c == 1 || c == 0 && hasCycle(y) {
					return true
				}
			}
			color[x] = 2
			return false
		}
		ok2 := true
		for i, c := range color {
			// 只遍历不在 sub 中的字母
			if c == 0 && sub>>i&1 == 0 && hasCycle(i) {
				ok2 = false
				break
			}
		}

		if ok2 { // 无环
			if size < minSize {
				minSize = size
				set = map[int]struct{}{sub: {}}
			} else {
				set[sub] = struct{}{}
			}
		}

		sub = (sub - 1) & all
	}

	base := [26]int{}
	for s := uint(all); s > 0; s &= s - 1 {
		base[bits.TrailingZeros(s)] = 1
	}
	for mask := range set {
		cnt := base
		for s := uint(mask); s > 0; s &= s - 1 {
			cnt[bits.TrailingZeros(s)] = 2
		}
		ans = append(ans, cnt[:])
	}
	return
}

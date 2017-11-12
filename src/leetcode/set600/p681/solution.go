package main

import (
	"fmt"
	"math"
)

func nextClosestTime(time string) string {
	a, b, c, d := int(time[0]-'0'), int(time[1]-'0'), int(time[3]-'0'), int(time[4]-'0')

	asMin := func(a, b, c, d int) int {
		return (a*10+b)*60 + c*10 + d
	}

	start := asMin(a, b, c, d)

	nums := []int{a, b, c, d}

	ans := math.MaxInt32
	res := ""
	for i := 0; i < 4; i++ {
		x := nums[i]
		if x > 2 {
			continue
		}
		for j := 0; j < 4; j++ {
			y := nums[j]
			if x == 2 && y > 4 {
				continue
			}

			for k := 0; k < 4; k++ {
				w := nums[k]
				if w > 5 {
					continue
				}

				for l := 0; l < 4; l++ {
					v := nums[l]
					tmp := asMin(x, y, w, v)

					if tmp <= start {
						tmp += 24 * 60
					}
					if tmp < ans {
						ans = tmp
						res = fmt.Sprintf("%d%d:%d%d", x, y, w, v)
					}
				}
			}
		}
	}
	return res
}

func main() {
	//fmt.Println(nextClosestTime("19:34"))
	//fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("00:13"))
	fmt.Println(nextClosestTime("00:00"))

}

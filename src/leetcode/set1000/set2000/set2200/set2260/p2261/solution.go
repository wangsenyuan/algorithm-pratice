package p2261

func countDistinct(nums []int, k int, p int) int {
	n := len(nums)

	mem := make(map[Key]bool)

	for i := 0; i < n; i++ {
		var cnt int
		var key Key
		for j := i; j < n; j++ {
			key = key.Mul(201).Add(int64(nums[j]))
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			mem[key] = true
		}
	}

	return len(mem)
}

const MOD1 = 10000000007
const MOD2 = 10000000009

type Key struct {
	first  int64
	second int64
}

func (this Key) Mul(num int64) Key {
	first := this.first * num % MOD1
	second := this.second * num % MOD2
	return Key{first, second}
}

func (this Key) Add(num int64) Key {
	first := (this.first + num) % MOD1
	second := (this.second + num) % MOD2
	return Key{first, second}
}

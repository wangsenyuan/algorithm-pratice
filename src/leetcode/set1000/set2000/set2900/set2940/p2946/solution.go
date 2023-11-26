package p2946

func beautifulSubstrings(s string, k int) int {
	k = sqrt(4 * k)
	var ans int64
	var sum int

	freq := make(map[Pair]int)

	freq[Pair{k - 1, 0}]++

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'o' || s[i] == 'i' || s[i] == 'u' || s[i] == 'e' {
			sum++
		} else {
			sum--
		}
		p := Pair{i % k, sum}
		ans += int64(freq[p])
		freq[p]++
	}

	return int(ans)
}

type Pair struct {
	first  int
	second int
}

func sqrt(num int) int {
	var res = 1
	for i := 2; i*i <= num; i++ {
		i2 := i * i
		for num%i2 == 0 {
			res *= i
			num /= i2
		}
		if num%i == 0 {
			res *= i
			num /= i
		}
	}

	res *= num

	return res
}

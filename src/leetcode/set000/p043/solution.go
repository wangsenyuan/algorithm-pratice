package p043

import "bytes"

func multiply(num1 string, num2 string) string {
	a := toNumArray(num1)
	b := toNumArray(num2)
	res := mul(a, b)
	if len(res) == 0 {
		return "0"
	}
	return arrayToString(res)
}

func normalize(num []int, n int) []int {
	res := make([]int, n)
	copy(res[n-len(num):], num)
	return res
}

func toNumArray(str string) []int {
	n := len(str)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(str[i] - '0')
	}
	return res
}

func arrayToString(num []int) string {
	var buf bytes.Buffer
	for i := 0; i < len(num); i++ {
		buf.WriteByte(byte(num[i] + '0'))
	}
	return buf.String()
}

func mul(num1 []int, num2 []int) []int {
	if len(num1) == 0 || len(num2) == 0 {
		return nil
	}

	if num1[0] < 0 && num2[0] < 0 {
		return mul(num1[1:], num2[1:])
	}

	if num1[0] < 0 {
		return neg(mul(num1[1:], num2))
	}
	if num2[0] < 0 {
		return neg(mul(num1, num2[1:]))
	}

	m := len(num1)
	n := len(num2)
	if m < 3 || n < 3 {
		return mulDirectly(num1, num2)
	}
	k := max(m, n)

	if m != n {
		num1 = normalize(num1, k)
		num2 = normalize(num2, k)
	}

	k2 := k / 2
	k3 := k - k2
	high1, low1 := num1[:k2], num1[k2:]
	high2, low2 := num2[:k2], num2[k2:]
	z0 := mul(low1, low2)
	z1 := mul(add(low1, high1), add(low2, high2))
	z2 := mul(high1, high2)
	first := shift(z2, k3*2)
	second := shift(sub(sub(z1, z2), z0), k3)
	res := add(add(first, second), z0)
	return dropPrefixZero(res)
}

func shift(num []int, cnt int) []int {
	res := make([]int, len(num)+cnt)
	copy(res, num)
	return res
}

func clone(num []int) []int {
	res := make([]int, len(num))
	copy(res, num)
	return res
}

func add(num1, num2 []int) []int {
	if len(num1) == 0 {
		return clone(num2)
	}
	if len(num2) == 0 {
		return clone(num1)
	}

	if num2[0] < 0 {
		//negative
		return sub(num1, num2[1:])
	}
	if num1[0] < 0 {
		return sub(num2, num1[1:])
	}

	i := len(num1) - 1
	j := len(num2) - 1
	res := make([]int, max(len(num1), len(num2))+1)
	var c int
	var p int
	for i >= 0 && j >= 0 {
		s := num1[i] + num2[j] + c
		res[p] = s % 10
		p++
		c = s / 10
		i--
		j--
	}

	for i >= 0 {
		s := num1[i] + c
		res[p] = s % 10
		p++
		c = s / 10
		i--
	}

	for j >= 0 {
		s := num2[j] + c
		res[p] = s % 10
		p++
		c = s / 10
		j--
	}

	if c > 0 {
		res[p] = c
		p++
	}
	res = res[:p]
	for i, j := 0, p-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func sub(num1, num2 []int) []int {
	if len(num2) == 0 {
		return clone(num1)
	}
	if len(num1) == 0 {
		// 0 - num2
		if num2[0] < 0 {
			// get postive
			return clone(num2[1:])
		}
		return neg(num2)
	}

	if num2[0] < 0 {
		return add(num1, num2[1:])
	}

	if num1[0] < 0 {
		tmp := add(num1[1:], num2)
		return neg(tmp)
	}
	// both postive and have digits
	if less(num1, num2) {
		tmp := sub(num2, num1)
		return neg(tmp)
	}
	// num1 >= num2
	i := len(num1) - 1
	j := len(num2) - 1
	res := make([]int, len(num1))
	for i >= 0 && j >= 0 {
		a := num1[i]
		b := num2[j]
		if a >= b {
			res[i] = a - b
		} else {
			//need to borrow
			k := i - 1
			for num1[k] == 0 {
				k--
			}
			num1[k]--
			k++
			for k < i {
				num1[k] = 9
				k++
			}
			res[i] = 10 + a - b
		}
		i--
		j--
	}
	for i >= 0 {
		res[i] = num1[i]
		i--
	}

	return dropPrefixZero(res)
}

func less(num1, num2 []int) bool {
	if len(num1) < len(num2) {
		return true
	} else if len(num1) == len(num2) {
		for i := 0; i < len(num1); i++ {
			if num1[i] < num2[i] {
				return true
			} else if num1[i] > num2[i] {
				return false
			}
		}
	}
	return false
}

func neg(nums []int) []int {
	res := make([]int, len(nums)+1)
	res[0] = -1
	copy(res[1:], nums)
	return res
}

func mulDirectly(num1, num2 []int) []int {
	res := make([]int, len(num1)+len(num2))
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			res[i+j+1] += num1[i] * num2[j]
		}
	}

	for k := len(res) - 1; k >= 0; k-- {
		if k > 0 {
			res[k-1] += res[k] / 10
		}
		res[k] %= 10
	}

	return dropPrefixZero(res)
}

func dropPrefixZero(num []int) []int {
	var i = 0
	for i < len(num) && num[i] == 0 {
		i++
	}
	return num[i:]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

package main

import "fmt"

func main() {
	fmt.Println(addStrings("123", "32"))
}

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	var res []byte
	carray := 0
	k := 0
	for i-k >= 0 && j-k >= 0 {
		a, b := int(num1[i-k]-'0'), int(num2[j-k]-'0')
		c := a + b + carray
		if c >= 10 {
			res = append(res, byte(c-10+'0'))
			carray = 1
		} else {
			res = append(res, byte(c+'0'))
			carray = 0
		}
		k++
	}

	for i-k >= 0 {
		c := int(num1[i-k]-'0') + carray
		if c >= 10 {
			res = append(res, byte(c-10+'0'))
			carray = 1
		} else {
			res = append(res, byte(c+'0'))
			carray = 0
		}
		k++
	}

	for j-k >= 0 {
		c := int(num2[j-k]-'0') + carray
		if c >= 10 {
			res = append(res, byte(c-10+'0'))
			carray = 1
		} else {
			res = append(res, byte(c+'0'))
			carray = 0
		}
		k++
	}
	if carray > 0 {
		res = append(res, byte(carray+'0'))
	}

	for a, b := 0, len(res)-1; a < b; a, b = a+1, b-1 {
		res[a], res[b] = res[b], res[a]
	}

	return string(res)
}

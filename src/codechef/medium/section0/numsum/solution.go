package main

import (
	"bytes"
	"strconv"
	"bufio"
	"os"
	"fmt"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	t := readNum(scanner)

	for t > 0 {
		scanner.Scan()
		fmt.Println(solve(scanner.Text()))
		t--
	}
}

func solve(s string) string {
	X := OP{toNum(s), 1, len(s)}
	ONE := OP{[]int{1}, 1, 1}
	Y := add(X, ONE)
	if X.rep[X.size-1]%2 == 0 {
		X = half(X)
	} else {
		Y = half(Y)
	}
	if X.size < Y.size {
		X = expandSize(X, Y.size)
	} else if Y.size < X.size {
		Y = expandSize(Y, X.size)
	}
	Z := mul(X, Y)

	return toStr(Z.rep)
}

func expandSize(X OP, n int) OP {
	res := make([]int, n)
	i := n - X.size
	copy(res[i:], X.rep)
	return OP{res, 1, n}
}

func toStr(num []int) string {
	var buf bytes.Buffer
	for i := 0; i < len(num); i++ {
		buf.WriteByte(byte(num[i] + '0'))
	}
	return buf.String()
}

func toNum(s string) []int {
	n := len(s)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(s[i] - '0')
	}
	return res
}

func toNumX(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res = res*10 + nums[i]
	}
	return res
}

type OP struct {
	rep  []int
	sign int
	size int
}

func (op OP) String() string {
	s := toStr(op.rep)
	if op.sign < 0 {
		return "-" + s
	}
	return s
}

func (op OP) left(n int) OP {
	return OP{op.rep[:n], op.sign, n}
}

func (op OP) right(n int) OP {
	return OP{op.rep[n:], op.sign, op.size - n}
}

func mul(X, Y OP) OP {
	n := X.size
	if n < 3 {
		x, y := toNumX(X.rep), toNumX(Y.rep)
		z := x * y
		zs := strconv.Itoa(z)
		return OP{toNum(zs), X.sign * Y.sign, len(zs)}
	}

	XL, XR := X.left(n/2), X.right(n/2)
	YL, YR := Y.left(n/2), Y.right(n/2)

	A := mul(XL, YL)
	B := mul(XR, YR)
	a := add(XL, XR)
	b := add(YL, YR)
	C := mul(a, b)
	D := sub(sub(C, A), B)

	return add(add(pow(A, (n+1)/2*2), pow(D, (n+1)/2)), B)
}

func add(X, Y OP) OP {
	if (X.sign > 0) != (Y.sign > 0) {
		if Y.sign < 0 {
			return sub(X, OP{Y.rep, 1, Y.size})
		}

		if X.sign < 0 {
			return sub(Y, OP{X.rep, 1, X.size})
		}
	}
	sign := X.sign
	// treat both as postive
	n := X.size
	if n < Y.size {
		n = Y.size
	}

	res := make([]int, n+1)
	i, j := X.size-1, Y.size-1
	k := n
	var carry int
	for i >= 0 && j >= 0 {
		sum := X.rep[i] + Y.rep[j] + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		res[k] = sum
		k--
		i--
		j--
	}

	for i >= 0 {
		sum := X.rep[i] + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		res[k] = sum
		k--
		i--
	}
	for j >= 0 {
		sum := Y.rep[j] + carry
		if sum >= 10 {
			sum -= 10
			carry = 1
		}
		res[k] = sum
		k--
		j--
	}
	if carry > 0 {
		res[k] = carry
		k--
	}
	k++
	for k < n+1 && res[k] == 0 {
		k++
	}
	return OP{res[k:], sign, (n + 1 - k)}
}

func sub(X, Y OP) OP {
	if X.sign < 0 {
		Z := sub(OP{X.rep, 1, X.size}, OP{Y.rep, -Y.sign, Y.size})
		return OP{Z.rep, -Z.sign, Z.size}
	}
	// X.sign > 0 now
	if Y.sign < 0 {
		return add(X, OP{Y.rep, 1, Y.size})
	}

	// X.sign > 0 && Y.sign > 0
	// X < Y
	if less(X, Y) {
		// switch ops
		Z := sub(Y, X)
		return OP{Z.rep, -Z.sign, Z.size}
	}
	n := X.size
	res := make([]int, n)
	// X >= Y
	i, j := X.size-1, Y.size-1
	k := n - 1

	for i >= 0 && j >= 0 {
		x := X.rep[i]
		if x < Y.rep[j] {
			p := i - 1
			for p >= 0 && X.rep[p] == 0 {
				p--
			}
			//X.rep[p] > 0
			X.rep[p]--
			for q := p + 1; q < i; q++ {
				X.rep[q] = 9
			}
			// borrow from previous digit
			x = 10 + x
		}
		res[k] = x - Y.rep[j]
		k--
		i--
		j--
	}
	for i >= 0 {
		res[k] = X.rep[i]
		k--
		i--
	}
	k++
	for k < n-1 && res[k] == 0 {
		k++
	}
	//res[k] > 0
	return OP{res[k:], 1, n - k}
}

func less(X, Y OP) bool {
	if X.size < Y.size {
		return true
	}
	if X.size > Y.size {
		return false
	}
	for i := 0; i < X.size; i++ {
		if X.rep[i] < Y.rep[i] {
			return true
		}
		if X.rep[i] > Y.rep[i] {
			return false
		}
	}
	return false
}

func pow(X OP, n int) OP {
	res := make([]int, X.size+n)
	copy(res, X.rep)
	return OP{res, X.sign, X.size + n}
}

func half(X OP) OP {
	n := X.size
	res := make([]int, n)
	var rem int
	for i := 0; i < n; i++ {
		rem *= 10
		res[i] = (rem + X.rep[i]) / 2
		rem = (rem + X.rep[i]) - 2*res[i]
	}
	if res[0] == 0 {
		return OP{res[1:], 1, n - 1}
	}
	return OP{res, 1, n}
}

func calculate(n string) string {
	if len(n) > 18 {
		panic("can't handle digits more than 18")
	}
	var x uint64
	for i := 0; i < len(n); i++ {
		x = x*10 + uint64(n[i]-'0')
	}
	s := x * (x + 1) / 2
	var buf bytes.Buffer

	for s > 0 {
		buf.WriteByte(byte('0' + s%10))
		s /= 10
	}
	res := buf.Bytes()
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

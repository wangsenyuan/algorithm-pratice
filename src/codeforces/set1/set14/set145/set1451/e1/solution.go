package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	ask := func(tp string, i int, j int) int {
		fmt.Printf("%s %d %d\n", tp, i, j)
		return readNum(reader)
	}

	res := solve(n, ask)

	var buf bytes.Buffer

	buf.WriteString("!")
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf(" %d", res[i]))
	}
	buf.WriteByte('\n')
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(n int, ask func(tp string, i int, j int) int) []int {
	// i & j, i | j, i xor j
	// a and b
	// a xor b
	// a or b
	//  假设  a & b = 1  a ^ b = 0
	//       a & b = 0, a ^ b = 0
	//       a & b = 0, a | b = 1 => a ^ b = 1
	//       但是这个1在a，还是b中，就不好说了
	// 只有两个数的时候是没法判断的
	// 如果有c
	//     必须找到c[i] 肯定为1的情况
	//     或者找到c[i]为0的情况
	// xor 有个好处是， a ^ b ^ c => a ^ c
	// 可以省掉一次查询
	// 先用n-1次查询xor，和前面的进行
	// 然后按位进行查询，对于d
	// 找到和 0, xor(1, j) [d] = 1 的位置j
	// 如果全部都是0，说明这一位大家都一致，（需要而外的一次or 或者 and查询)
	// 如果知道 A[1]是什么，那么xor就可以知道A[j]是什么
	// 现在就是如何在3次额外的查询中，知道A[1]是什么？
	// 如果用A[1]和某个数进行or操作，为1的部分，是 A[1]中有可能为1的部分， 为0的部分，肯定为0
	//     A[1]... and 操作， 为1的部分，肯定为1， 为0的部分，可能为0
	// 通过这两次操作，可以知道哪些肯定为0，肯定为1，尚不确定的部分
	// 不确定的位是  A[1] | A[j] = 1
	//             A[1] & A[k] = 0
	// 进行一次 A[j] and A[k] 因为已经知道 A[j] xor A[k] => 也可以推导出 A[j] or A[k]
	//   如果 A[j] and A[k] = 1, 那么A[k]在这位肯定位1 => A[1] 在这一位肯定位0
	//    A[j] and A[k] = 0, A[j] 在这一位肯定位0， => A[1] 在这一位肯定为1
	xor := make([]int, n+1)

	for i := 2; i <= n; i++ {
		xor[i] = ask("XOR", 1, i)
	}

	or2 := ask("OR", 1, 2)
	and3 := ask("AND", 1, 3)
	and23 := ask("AND", 2, 3)
	xor23 := xor[2] ^ xor[3]

	ans := make([]int, n+1)

	for d := 0; 1<<d < n; d++ {
		if (or2>>d)&1 == 0 {
			// ans[1][d] must be 0
			continue
		}
		if (and3>>d)&1 == 1 {
			ans[1] |= 1 << d
			continue
		}
		// need to check
		if (and23>>d)&1 == 1 {
			// ans[2][d] = 1, ans[3][d] = 1
			if (and3>>d)&1 == 1 {
				ans[1] |= 1 << d
				continue
			}
			// else ans[1][d] = 0
		} else {
			// and23[d] = 0 至少2和3中有一个为0
			if (xor23>>d)&1 == 0 {
				// both be 0
				if (or2>>d)&1 == 1 {
					ans[1] |= 1 << d
				}
			} else {
				// 恰好有一个0/1，但现在不确定是2还是3
				// 如果 ans[1][d] = 0, 则 ans[2][d] = 1, ans[3][d] = 0
				// 如果 ans[1][d] = 1, 则 ans[3][d] = 0, ans[2][d] = 1
				if (xor[2]>>d)&1 == 0 {
					ans[1] |= 1 << d
				}
			}
		}
	}

	for i := 2; i <= n; i++ {
		ans[i] = xor[i] ^ ans[1]
	}

	return ans[1:]
}

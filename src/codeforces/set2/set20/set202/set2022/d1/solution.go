package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	ask := func(i int, j int) int {
		fmt.Printf("? %d %d\n", i, j)
		return readNum(reader)
	}

	for range tc {
		n := readNum(reader)
		res := solve(n, ask)
		fmt.Printf("! %d\n", res)
	}

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

func solve(n int, ask func(int, int) int) int {

	check := func(i int, x int, y int, z int) int {
		// 如果i是Imposter，i+1是Knave，i+2是Knight 101
		// 如果i是Imposter，i+1是Knave， i+2也是Knave, 110
		// 如果i是Imposter，i+1是Knight，i+2也是Knight, 011
		// 如果i是Imposter，i+1是Knight, i+2是Knave 000
		// 这里有两种情况，一种是3个职业都出现了， 且是 I -> N -> T 的顺序（循环）
		// 还有一种情况是，只出现了I和两个T, I -> T -> T
		// 或者出现了I和两个N, I -> N -> N
		u := ask(i+1, i)
		v := ask(i+2, i+1)
		w := ask(i, i+2)
		// ITT 和 INN， u+v+w = 2 还是成立的， 但是如果是I->N->T 这个顺序得到的是0
		if u+v+w == 0 {
			if x == 0 {
				// 只有 N -> T 是0
				return i + 2
			}
			if y == 0 {
				return i
			}
			return i + 1
		}
		// 如果是ITT， 那么 x = 0, y = 1, z = 1
		// 如果是INN， 那么 x = 1, y = 1, z = 0
		// TIT 呢？
		if x == u && x == 1 {
			// i 和 i+1肯定是一样的
			return i + 2
		}
		if y == v && y == 1 {
			return i
		}
		return i + 1
	}

	check2 := func(i int, x int, y int, z int) int {
		// 全部为false的情况下，3个职业都不一样，
		// 如果i是Imposter, i+1是Knight，i+2是Knave是ok的
		// 如果i是Knight，i+1是Knave, i+2是Imposter也是ok的
		// 如果i是Knave， i+1是Imposter, i+2是Knight也是ok的
		// 循环顺序肯定是 ITN, NIT, TNI
		u := ask(i+1, i)
		v := ask(i, i+2)

		if x == 0 && u == 1 {
			//只能是前两种情况
			// z = 0 肯定等于0
			if v == 1 {
				return i
			}
			return i + 1
		}
		// y肯定是为0的（后面两种情况）
		// NT和，TI的回答是不一样的
		if v == 1 {
			return i + 2
		}
		return i + 1
	}

	i := 1
	for i+2 <= n {
		x := ask(i, i+1)
		y := ask(i+1, i+2)
		z := ask(i+2, i)

		if x+y+z == 0 {
			return check2(i, x, y, z)
		} else if x+y+z == 2 {
			return check(i, x, y, z)
		}

		i += 3
	}
	// 最后一个了
	if i == n {
		return i
	}
	// 最后两个 i, i + 1
	x := ask(i-1, i)
	y := ask(i, i+1)
	z := ask(i+1, i-1)
	if x+y+z == 0 {
		return check2(i-1, x, y, z)
	}
	return check(i-1, x, y, z)
}

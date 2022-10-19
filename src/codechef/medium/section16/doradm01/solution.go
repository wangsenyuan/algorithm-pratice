package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		s := make([]string, n)
		for i := 0; i < n; i++ {
			s[i] = readString(reader)
		}
		blocks := make([][]int, k)
		for i := 0; i < k; i++ {
			blocks[i] = readNNums(reader, 2)
		}
		res := solve(n, s, blocks)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}

	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

	for len(bs) == 0 || bs[0] == '\n' || bs[0] == '\r' {
		bs, _ = reader.ReadBytes('\n')
	}

	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(n int, s []string, blocks [][]int) bool {
	inv := make([][]bool, n)
	for i := 0; i < n; i++ {
		inv[i] = make([]bool, n)
	}
	for _, block := range blocks {
		a, b := block[0], block[1]
		a--
		b--
		inv[a][b] = true
	}

	equenow := make([]*BitSet, n)
	equepre := make([]*BitSet, n)
	equetmp := make([]*BitSet, n)
	var pro []*BitSet

	for i := 0; i < n; i++ {
		equenow[i] = NewBitSet(256)
		equenow[i].Set(i)

		equepre[i] = NewBitSet(256)
		equetmp[i] = NewBitSet(256)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			equetmp[j].Copy(equenow[j])
			if i > 0 {
				equetmp[j].Xor(equepre[j])
			}
			if j > 0 {
				equetmp[j].Xor(equenow[j-1])
			}
			if j+1 < n {
				equetmp[j].Xor(equenow[j+1])
			}
			if s[i][j] == '1' {
				equetmp[j].Flip(n)
			}
			if inv[i][j] {
				tmp := NewBitSet(256)
				tmp.Copy(equenow[j])
				pro = append(pro, tmp)
			}
		}
		for j := 0; j < n; j++ {
			equepre[j].Copy(equenow[j])
			equenow[j].Copy(equetmp[j])
		}
	}

	for j := 0; j < n; j++ {
		pro = append(pro, equenow[j])
	}

	return check(pro, n)
}

func check(pro []*BitSet, n int) bool {
	for i := 0; i < n; i++ {
		for j := i + 1; j < len(pro); j++ {
			if !pro[j].IsSet(i) {
				continue
			}
			if !pro[i].IsSet(i) {
				pro[i], pro[j] = pro[j], pro[i]
			} else {
				pro[j].Xor(pro[i])
			}
		}
	}

	for i := 0; i < len(pro); i++ {
		allZero := true
		for j := 0; j < n; j++ {
			if pro[i].IsSet(j) {
				allZero = false
				break
			}
		}
		if allZero && pro[i].IsSet(n) {
			return false
		}
	}
	return true
}

type BitSet struct {
	arr []uint64
	sz  int
}

func NewBitSet(sz int) *BitSet {
	n := sz / 64
	arr := make([]uint64, n)
	return &BitSet{arr, sz}
}

func (this *BitSet) Set(p int) {
	x, y := p/64, p%64
	this.arr[x] |= (1 << uint64(y))
}

func (this *BitSet) Flip(p int) {
	x, y := p/64, p%64
	this.arr[x] ^= (1 << uint64(y))
}

func (this *BitSet) IsSet(p int) bool {
	x, y := p/64, p%64
	return (this.arr[x]>>uint64(y))&1 == 1
}

func (this *BitSet) Xor(that *BitSet) *BitSet {
	for i := 0; i < len(this.arr); i++ {
		this.arr[i] ^= that.arr[i]
	}
	return this
}

func (this *BitSet) Copy(that *BitSet) {
	copy(this.arr, that.arr)
}

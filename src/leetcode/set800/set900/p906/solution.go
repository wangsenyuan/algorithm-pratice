package p906

import (
	"sort"
	"strconv"
	"strings"
)

var ps []uint64

func init() {
	s := "0 1 4 9 121 484 10201 12321 14641 40804 44944 1002001 1234321 4008004 100020001 102030201 104060401 121242121 123454321 125686521 400080004 404090404 10000200001 10221412201 12102420121 12345654321 40000800004 1000002000001 1002003002001 1004006004001 1020304030201 1022325232201 1024348434201 1210024200121 1212225222121 1214428244121 1232346432321 1234567654321 4000008000004 4004009004004 100000020000001 100220141022001 102012040210201 102234363432201 121000242000121 121242363242121 123212464212321 123456787654321 400000080000004 10000000200000001 10002000300020001 10004000600040001 10020210401202001 10022212521222001 10024214841242001 10201020402010201 10203040504030201 10205060806050201 10221432623412201 10223454745432201 12100002420000121 12102202520220121 12104402820440121 12122232623222121 12124434743442121 12321024642012321 12323244744232321 12343456865434321 12345678987654321 40000000800000004 40004000900040004 1000000002000000001 1000220014100220001 1002003004003002001 1002223236323222001 1020100204020010201 1020322416142230201 1022123226223212201 1022345658565432201 1210000024200000121 1210242036302420121 1212203226223022121 1212445458545442121 1232100246420012321 1232344458544432321 1234323468643234321 4000000008000000004"
	nums := strings.Split(s, " ")
	ps = make([]uint64, len(nums))
	for i := 0; i < len(nums); i++ {
		ps[i] = toInt64([]byte(nums[i]))
	}
	/*ps = make([]uint64, 0, 1000)
	ps = append(ps, 0)

	for i := 0; i <= 1000000000; i++ {
		if (isPalindrome(uint64(i))) {
			num := uint64(i)
			j := num * num
			if isPalindrome(j) {
				ps = append(ps, j)
			}
		}

	}

	for i := 1; i <= 100000; i++ {
		ii := mirror(i)
		j := ii * ii
		if isPalindrome(j) {
			ps = append(ps, j)
		}
	}

	sort.Sort(Int64Slice(ps))

	var j int
	for i := 1; i <= len(ps); i++ {
		if i < len(ps) && ps[i] == ps[i-1] {
			i++
		}
		ps[j] = ps[i-1]
		j++
	}
	ps = ps[:j]
	fmt.Println(os.Stderr, ps)*/
}

type Int64Slice []uint64

func (this Int64Slice) Len() int {
	return len(this)
}

func (this Int64Slice) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this Int64Slice) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func isPalindrome(num uint64) bool {
	var rev uint64
	tmp := num
	for tmp != 0 {
		rev = rev*10 + tmp%10
		tmp /= 10
	}
	return rev == num
}

func mirror(i int) uint64 {
	s := []byte(strconv.Itoa(i))
	ss := make([]byte, len(s)*2)
	copy(ss, s)
	for i, j := len(s)-1, len(s); i >= 0; i, j = i-1, j+1 {
		ss[j] = s[i]
	}
	return toInt64(ss)
}

func toInt64(s []byte) uint64 {
	var res uint64
	for i := 0; i < len(s); i++ {
		res = res*10 + uint64(s[i]-'0')
	}
	return res
}

func superpalindromesInRange(L string, R string) int {
	a := toInt64([]byte(L))
	b := toInt64([]byte(R))

	i := sort.Search(len(ps), func(i int) bool {
		return ps[i] >= a
	})

	j := sort.Search(len(ps), func(j int) bool {
		return ps[j] > b
	})
	return j - i
}

package p3043

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	tr := new(Trie)

	for _, num := range arr1 {
		arr := getDigits(num)
		tr.Add(arr, 0)
	}

	var res int

	for _, num := range arr2 {
		arr := getDigits(num)

		tmp := tr
		for i := 0; i < len(arr); i++ {
			x := arr[i]
			if tmp.children[x] == nil {
				res = max(res, i)
				tmp = nil
				break
			}
			tmp = tmp.children[x]
		}
		if tmp != nil {
			res = max(res, len(arr))
		}
	}

	return res
}

func getDigits(num int) []int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	reverse(arr)
	return arr
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type Trie struct {
	children [10]*Trie
}

func (t *Trie) Add(nums []int, i int) {
	if i == len(nums) {
		return
	}
	x := nums[i]
	if t.children[x] == nil {
		t.children[x] = new(Trie)
	}
	t.children[x].Add(nums, i+1)
}

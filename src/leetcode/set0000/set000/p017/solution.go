package p017

var board = make(map[byte]string)

func init() {
	board['2'] = "abc"
	board['3'] = "def"
	board['4'] = "ghi"
	board['5'] = "jkl"
	board['6'] = "mno"
	board['7'] = "pqrs"
	board['8'] = "tuv"
	board['9'] = "wxyz"
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	result := make([]string, 0, 10)

	var dfs func(int, string)

	dfs = func(i int, path string) {
		if i == len(digits) {
			result = append(result, path)
			return
		}

		digit := digits[i]
		if v, found := board[digit]; found {
			for _, x := range v {
				dfs(i+1, path+string(x))
			}
		}
	}
	dfs(0, "")
	return result
}

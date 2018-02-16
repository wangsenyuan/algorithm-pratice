package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] == ' ' {
		i++
	}
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readOneNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
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

	t := readOneNum(scanner)

	for t > 0 {
		n := readOneNum(scanner)
		nums := readNNums(scanner, n)

		fmt.Println(solve(n, nums))

		t--
	}
}

const INF = 1 << 30

func solve(n int, nums []int) int {
	a := solve1(n, nums)
	b := solve2(n, nums)

	if a < 0 {
		return b
	}
	if b < 0 {
		return a
	}
	return min(a, b)
}

func solve1(n int, nums []int) int {
	if n%2 == 1 {
		x := nums[n/2]
		if n/2 > 0 {
			if x >= nums[n/2-1] || x >= nums[n/2+1] {
				// center must be min value
				return -1
			}
		}
	}
	// a0 > a1 < a2 > a3 < a4
	a, b := 0, 1
	// a means no change in previous position
	// b means change in previous position
	for i, j := 1, n-2; i < j; i, j = i+1, j-1 {
		c, d := INF, INF
		if a < INF {
			if i%2 == 1 {
				if j%2 == 1 {
					if nums[i] < nums[i-1] && nums[j] < nums[j+1] {
						c = a
					}
					if nums[i] < nums[j+1] && nums[j] < nums[i] {
						// change is valid
						d = a + 1
					}
				} else {
					if nums[i] < nums[i-1] && nums[j] > nums[j+1] {
						c = a
					}
					if nums[i] > nums[j+1] && nums[j] < nums[i-1] {
						d = a + 1
					}
				}
			} else {
				// i is at max position
				if j%2 == 0 {
					// j is at max too
					if nums[i] > nums[i-1] && nums[j] > nums[j+1] {
						c = a
					}
					if nums[i] > nums[j+1] && nums[j] > nums[i-1] {
						d = a + 1
					}
				} else {
					// j is at min
					if nums[i] > nums[i-1] && nums[j] < nums[j+1] {
						c = a
					}
					if nums[j] > nums[i-1] && nums[i] < nums[j+1] {
						d = a + a
					}
				}
			}
		}

		if b < INF {
			if i%2 == 1 {
				if j%2 == 1 {
					if nums[i] < nums[n-i] && nums[j] < nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] < nums[n-j-2] && nums[j] < nums[n-i] {
						// change is valid
						d = min(d, b+1)
					}
				} else {
					if nums[i] < nums[n-i] && nums[j] > nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] < nums[n-j-2] && nums[j] > nums[n-i] {
						d = min(d, b+1)
					}
				}
			} else {
				// i is at max position
				if j%2 == 0 {
					// j is at max too
					if nums[i] > nums[n-i] && nums[j] > nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] < nums[n-j-2] && nums[j] < nums[n-i] {
						d = min(d, b+1)
					}
				} else {
					// j is at min
					if nums[i] > nums[n-i] && nums[j] < nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] < nums[n-j-2] && nums[j] < nums[n-i] {
						d = min(d, b+1)
					}
				}
			}
		}
		if c == INF && d == INF {
			return -1
		}
		a, b = c, d
	}

	return min(a, b)
}

func solve2(n int, nums []int) int {
	if n%2 == 1 {
		x := nums[n/2]
		if n/2 > 0 {
			if x <= nums[n/2-1] || x <= nums[n/2+1] {
				// center must be max value
				return -1
			}
		}
	}
	// a0 < a1 > a2 < a3 > a4
	a, b := 0, 1
	// a means no change in previous position
	// b means change in previous position
	for i, j := 1, n-2; i < j; i, j = i+1, j-1 {
		c, d := INF, INF
		if a < INF {
			if i%2 == 1 {
				if j%2 == 1 {
					if nums[i] > nums[i-1] && nums[j] > nums[j+1] {
						c = a
					}
					if nums[i] > nums[j+1] && nums[j] > nums[i] {
						// change is valid
						d = a + 1
					}
				} else {
					if nums[i] > nums[i-1] && nums[j] < nums[j+1] {
						c = a
					}
					if nums[i] < nums[j+1] && nums[j] > nums[i-1] {
						d = a + 1
					}
				}
			} else {
				// i is at max position
				if j%2 == 0 {
					// j is at max too
					if nums[i] < nums[i-1] && nums[j] < nums[j+1] {
						c = a
					}
					if nums[i] < nums[j+1] && nums[j] < nums[i-1] {
						d = a + 1
					}
				} else {
					// j is at min
					if nums[i] < nums[i-1] && nums[j] > nums[j+1] {
						c = a
					}
					if nums[j] < nums[i-1] && nums[i] > nums[j+1] {
						d = a + a
					}
				}
			}
		}

		if b < INF {
			if i%2 == 1 {
				if j%2 == 1 {
					if nums[i] > nums[n-i] && nums[j] > nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] > nums[n-j-2] && nums[j] > nums[n-i] {
						// change is valid
						d = min(d, b+1)
					}
				} else {
					if nums[i] > nums[n-i] && nums[j] < nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] > nums[n-j-2] && nums[j] < nums[n-i] {
						d = min(d, b+1)
					}
				}
			} else {
				// i is at max position
				if j%2 == 0 {
					// j is at max too
					if nums[i] < nums[n-i] && nums[j] < nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] > nums[n-j-2] && nums[j] > nums[n-i] {
						d = min(d, b+1)
					}
				} else {
					// j is at min
					if nums[i] < nums[n-i] && nums[j] > nums[n-j-2] {
						c = min(c, b)
					}
					if nums[i] > nums[n-j-2] && nums[j] > nums[n-i] {
						d = min(d, b+1)
					}
				}
			}
		}
		if c == INF && d == INF {
			return -1
		}
		a, b = c, d
	}
	return min(a, b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

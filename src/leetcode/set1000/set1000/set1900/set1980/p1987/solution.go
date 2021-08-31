package p1987

/**
You are given a binary string binary. A subsequence of binary is considered good if it is not empty and has no leading zeros (with the exception of "0").

Find the number of unique good subsequences of binary.

For example, if binary = "001", then all the good subsequences are ["0", "0", "1"], so the unique good subsequences are "0" and "1". Note that subsequences "00", "01", and "001" are not good because they have leading zeros.
Return the number of unique good subsequences of binary. Since the answer may be very large, return it modulo 109 + 7.

A subsequence is a sequence that can be derived from another sequence by deleting some or no elements without changing the order of the remaining elements.

1 <= binary.length <= 105
binary consists of only '0's and '1's.
*/

const MOD = 1000000007

func numberOfUniqueGoodSubsequences(binary string) int {
	//如果当前为1，不考虑unique的情况下，good subseq 的数量是 pow(2, l)
	var odd, even int
	var hasZero bool
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			hasZero = true
			even = (odd + even) % MOD
		} else {
			// f[1] = f[0] + f[1]
			odd = (odd + even + 1) % MOD
		}
	}
	res := (odd + even) % MOD
	if hasZero {
		res = (res + 1) % MOD
	}
	return res
}

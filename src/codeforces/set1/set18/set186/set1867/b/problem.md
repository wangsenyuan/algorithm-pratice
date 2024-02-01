You are given a binary string 𝑠
of length 𝑛
(a string that consists only of 0
and 1
). A number 𝑥
is good if there exists a binary string 𝑙
of length 𝑛
, containing 𝑥
ones, such that if each symbol 𝑠𝑖
is replaced by 𝑠𝑖⊕𝑙𝑖
(where ⊕
denotes the bitwise XOR operation), then the string 𝑠
becomes a palindrome.

You need to output a binary string 𝑡
of length 𝑛+1
, where 𝑡𝑖
(0≤𝑖≤𝑛
) is equal to 1
if number 𝑖
is good, and 0
otherwise.

A palindrome is a string that reads the same from left to right as from right to left. For example, 01010, 1111, 0110
are palindromes.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
(1≤𝑡≤105
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
(1≤𝑛≤105
).

The second line of each test case contains a binary string 𝑠
of length 𝑛
.

It is guaranteed that the sum of 𝑛
over all test cases does not exceed 105
.

### thoughts

1. 考虑s ^ l 后是一个回文串
2. 如果 s[i] == s[n - 1 - i], 那么 l[i] == l[n-1-i] (都为0或者1)
3. 如果 s[i] != s[n-1-i], 那么 l[i] != l[n-1-i] (一个0，一个1)
4. 换句话说，那些不同位置的地方，它们至少贡献一个1，其他的至少两个1
5. 
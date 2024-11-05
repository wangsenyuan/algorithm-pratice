Kirk has a binary string 𝑠
 (a string which consists of zeroes and ones) of length 𝑛
 and he is asking you to find a binary string 𝑡
 of the same length which satisfies the following conditions:

For any 𝑙
 and 𝑟
 (1≤𝑙≤𝑟≤𝑛
) the length of the longest non-decreasing subsequence of the substring 𝑠𝑙𝑠𝑙+1…𝑠𝑟
 is equal to the length of the longest non-decreasing subsequence of the substring 𝑡𝑙𝑡𝑙+1…𝑡𝑟
;
The number of zeroes in 𝑡
 is the maximum possible.
A non-decreasing subsequence of a string 𝑝
 is a sequence of indices 𝑖1,𝑖2,…,𝑖𝑘
 such that 𝑖1<𝑖2<…<𝑖𝑘
 and 𝑝𝑖1≤𝑝𝑖2≤…≤𝑝𝑖𝑘
. The length of the subsequence is 𝑘
.

If there are multiple substrings which satisfy the conditions, output any.

### ideas
1. 最大化t中0的数量
2. 然后在给定的一段l...r中间， 
3. s[l...r]的非递减的子序列的长度和t[l...r]的非递减的子序列的长度相等（是相等，不是大于等于）
4. s[l...r]的非递减的子序列的长度(可以预先计算出来)
5. f(s[l...r]) >= max(cnt[0], cnt[1])
6. 所以，理论上，可以把前半部分f(...) 设置为1，后半部分全部设置为0，
7. 011， 010 =》 100
8. 1010 =〉1100 (这个不对， 因为中间部分是在s中是2)
9. 也就是说10, 不能替换, 还必须是10，
10. 但是 11 =》 00， 01 =》 00
11. 1110 =》 0010
12. 01110 =》 00010
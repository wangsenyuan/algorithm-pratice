We call a string good, if after merging all the consecutive equal characters, the resulting string is palindrome. For
example, "aabba" is good, because after the merging step it will become "aba".

Given a string, you have to find two values:

the number of good substrings of even length;
the number of good substrings of odd length.

### ideas

1. A substring s[l, r] (1 ≤ l ≤ r ≤ n) of string s = s1s2... sn is string slsl + 1... sr.
2. 字符串有a/b组成 （0/1)
3. 如果s[i] = s[i-1]， 那么到i-1的结果，可以直接用在i上，i会和i-1进行合并 （但是长度奇偶性要变化)
4. 如果s[i] != s[i-1] 那么就需要知道到i-1为止，(假设s[i-1] = 0) 前面是1，且以0开始的（1没有被匹配的）的回文字串的结果
5. 肯定能确定的一个点是，到目前为止，字符串的pattern必然是 0101 或者1010
6. 所以dp[0]表示 01, dp[1]表示10, dp[2] 表示101, dp[3]表示010
7. 然后遇到一个新的0 01 =》 010， 10 =》 10 （但是奇偶性要变）
8. 101 =》 010 （奇偶性有点麻烦了，如果是偶数个1，和奇数个1）
9. 010 =》 010 （奇偶性要变)
9. 不过至少进了一步
10. 01111 1000111
11. 如果不考虑merge的情况， 1010101010，其中任意两个0/1之间都是满足回文的条件的
12. 所以，对于当前0，只需要知道前面和自己相同parity位的0有多少个，加上去，就是奇数长度的计数
13. 100001
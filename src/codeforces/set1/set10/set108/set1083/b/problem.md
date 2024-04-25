Recently, the Fair Nut has written 𝑘
strings of length 𝑛
, consisting of letters "a" and "b". He calculated 𝑐
— the number of strings that are prefixes of at least one of the written strings. Every string was counted only one
time.

Then, he lost his sheet with strings. He remembers that all written strings were lexicographically not smaller than
string 𝑠
and not bigger than string 𝑡
. He is interested: what is the maximum value of 𝑐
that he could get.

A string 𝑎
is lexicographically smaller than a string 𝑏
if and only if one of the following holds:

𝑎
is a prefix of 𝑏
, but 𝑎≠𝑏
;
in the first position where 𝑎
and 𝑏
differ, the string 𝑎
has a letter that appears earlier in the alphabet than the corresponding letter in 𝑏

### ideas

1. 这个题目有点看不懂～～～
2. k个字符的长度是一样的, len(s) = len(t)
3. 然后这 k个字符 >= s and <= t
4. c是至少是这个其中一个字符串的pref的计数
5. 比如 aa bb 如果不考虑k,那么中间应该至少有 bb - aa(二进制表示的数量多的字符串)
6. 但是这里限制只有k个字符串
7. 如果bb和aa中间有k个字符串，那么c 至少 = k
8. 我们假设不考虑k的情况，如何计算c
9. 假设到目前为止 取的字符串在保证 t < b， t > a的情况下，理论上，后面t可以取剩下所有的可能性
10. 假设剩下的长度是w, 那么可能性就是 min(pow(2, w), k)
11. 但是对c的贡献是多少？
12. c = k + ？
13. 从这里可以看出来，要尽可能的长
14. 从后往前， dp[i]表示在保证s[i..] <= t[i...]的情况下
15. 假设 i是 a[i] < b[i]的第一个位置， 那么如果剩下的长度x <= 32 可以计算中间有多少个可能，
16. 如果 >= k, 那么就是k中可能 w = n
17. 如果 <= k

.
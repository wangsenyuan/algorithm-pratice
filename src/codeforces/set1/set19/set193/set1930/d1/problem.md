For a binary†
 pattern 𝑝
 and a binary string 𝑞
, both of length 𝑚
, 𝑞
 is called 𝑝
-good if for every 𝑖
 (1≤𝑖≤𝑚
), there exist indices 𝑙
 and 𝑟
 such that:

1≤𝑙≤𝑖≤𝑟≤𝑚
, and
𝑝𝑖
 is a mode‡
 of the string 𝑞𝑙𝑞𝑙+1…𝑞𝑟
.
For a pattern 𝑝
, let 𝑓(𝑝)
 be the minimum possible number of 𝟷
s in a 𝑝
-good binary string (of the same length as the pattern).

You are given a binary string 𝑠
 of size 𝑛
. Find
∑𝑖=1𝑛∑𝑗=𝑖𝑛𝑓(𝑠𝑖𝑠𝑖+1…𝑠𝑗).
In other words, you need to sum the values of 𝑓
 over all 𝑛(𝑛+1)2
 substrings of 𝑠
.

†
 A binary pattern is a string that only consists of characters 𝟶
 and 𝟷
.

‡
 Character 𝑐
 is a mode of string 𝑡
 of length 𝑚
 if the number of occurrences of 𝑐
 in 𝑡
 is at least ⌈𝑚2⌉
. For example, 𝟶
 is a mode of 𝟶𝟷𝟶
, 𝟷
 is not a mode of 𝟶𝟷𝟶
, and both 𝟶
 and 𝟷
 are modes of 𝟶𝟷𝟷𝟶𝟷𝟶
.

### ideas
1. q是p-good的字符串，表示
2. 对于下标i，存在l, r, p[i] = mode(q[l...r])
3. 比如对于字符串 p = 01
4. 那么q = 01 就是 p-good 的
5. p = 001, q = 011 或者 q = 001 都是 p-good 的
6. f(p) = 在p-good的字符串中最少的1的个数
7. 对于 p = 001, f(p) = 1 (q = 001)
8. 对于字符串s, 要球所有的f(s[i...j]) 所有子串的f(p)的和
9. 对于给定的p, q = p显然是它的一个符合条件的串
10. 有没有可能减少其中1的个数？
11. 如果p[i] = 1, 那么如果q[i-1] = 1， 那么 q[i]可以为0
12. 如果q[i-1] = 0, 但是q[i]必须为1吗？其实也不定，可以如果存在q[i+1], 可以让它为1
13. 当p[i] = 0时，q[i] = 0, 当p[i] = 1时，如果q[i-1] = 1, 则 q[i] = 0
14. 如果q[i-1] = 0, 但是存在 i+1, q[i] = 0, q[i+1] = 1 (即使p[i+1] = 0, 因为它可以和q组成一组，所以也是安全的)
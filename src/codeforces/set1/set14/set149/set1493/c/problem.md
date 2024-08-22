You are given a string 𝑠
 consisting of lowercase English letters and a number 𝑘
. Let's call a string consisting of lowercase English letters beautiful if the number of occurrences of each letter in that string is divisible by 𝑘
. You are asked to find the lexicographically smallest beautiful string of length 𝑛
, which is lexicographically greater or equal to string 𝑠
. If such a string does not exist, output −1
.

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
.

### ideas
1. 显然 n % k = 0, 必须成立
2. 可以假设在位置i处， x[i] > s[i]
3. 那么在这之前的字符的数量是确定的，然后剩余的，在保证freq[?] % k = 0 的情况下，尽量使用a即可
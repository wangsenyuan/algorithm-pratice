Prefix function of string 𝑡=𝑡1𝑡2…𝑡𝑛
 and position 𝑖
 in it is defined as the length 𝑘
 of the longest proper (not equal to the whole substring) prefix of substring 𝑡1𝑡2…𝑡𝑖
 which is also a suffix of the same substring.

For example, for string 𝑡=
 abacaba the values of the prefix function in positions 1,2,…,7
 are equal to [0,0,1,0,1,2,3]
.

Let 𝑓(𝑡)
 be equal to the maximum value of the prefix function of string 𝑡
 over all its positions. For example, 𝑓(
abacaba)=3
.

You are given a string 𝑠
. Reorder its characters arbitrarily to get a string 𝑡
 (the number of occurrences of any character in strings 𝑠
 and 𝑡
 must be equal). The value of 𝑓(𝑡)
 must be minimized. Out of all options to minimize 𝑓(𝑡)
, choose the one where string 𝑡
 is the lexicographically smallest.


 ### ideas
 1. prefix function就是lps(最长proper后缀)，用kmp算法就可以算出来
 2. 但是 f(s)是最大的值 lps[i]
 3. 这里需要f(s)最小，且在最小的情况下，最小的string
 4. 如果有个字符只出现了一次，貌似可以得到f(s) = 0的情况
 5. 如果出现了多次，f(s)至少是1
 6. 如果所有的都一样，没有选择
 7. 假设首字母是a,第二个字母是b,如果后面不出现ab连在一起的,那么答案就是f(t) <= 1 (如果没有a, 那么就是 f(t) = 0)
 8. 而且是不是f(t) = 1, 始终是可以达到的呢?（每个字符至少两次)
 9. 不一定
 10. 如果cnt(a) + 1 <= 出现的字符的unique的次数， 那么 f(t) = 1
 11. 如果不是，比如 aaaaabbbcccc 这样的一个字符串 
 12. babbbcaaaacc (f(t) = 1) 
 1.  如果 cnt(a) - 1 > 出现的次数 bbbaaaa
 2.  aaabaaba
 3.  
You are given a string 𝑠
. For a fixed 𝑘
, consider a division of 𝑠
 into exactly 𝑘
 continuous substrings 𝑤1,…,𝑤𝑘
. Let 𝑓𝑘
 be the maximal possible 𝐿𝐶𝑃(𝑤1,…,𝑤𝑘)
 among all divisions.

𝐿𝐶𝑃(𝑤1,…,𝑤𝑚)
 is the length of the Longest Common Prefix of the strings 𝑤1,…,𝑤𝑚
.

For example, if 𝑠=𝑎𝑏𝑎𝑏𝑎𝑏𝑐𝑎𝑏
 and 𝑘=4
, a possible division is 𝑎𝑏𝑎𝑏𝑎𝑏𝑐𝑎𝑏
. The 𝐿𝐶𝑃(𝑎𝑏,𝑎𝑏,𝑎𝑏𝑐,𝑎𝑏)
 is 2
, since 𝑎𝑏
 is the Longest Common Prefix of those four strings. Note that each substring consists of a continuous segment of characters and each character belongs to exactly one substring.

Your task is to find 𝑓𝑙,𝑓𝑙+1,…,𝑓𝑟
. In this version 𝑙=𝑟
.

### ideas
1. 可以二分查询吗？确实是可以的
2. 因为当长度为x的lcp存在的时候，长度为x-1的(x > 0)肯定也是存在的，
3. 且这个长度x就是前缀s[1:x]
4. 辅以kmp， 就可以很快的判断出来
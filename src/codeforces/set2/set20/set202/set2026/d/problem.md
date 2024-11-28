You are given a sequence of integers [𝑎1,𝑎2,…,𝑎𝑛]
. Let 𝑠(𝑙,𝑟)
 be the sum of elements from 𝑎𝑙
 to 𝑎𝑟
 (i. e. 𝑠(𝑙,𝑟)=∑𝑖=𝑙𝑟𝑎𝑖
).

Let's construct another sequence 𝑏
 of size 𝑛(𝑛+1)2
 as follows: 𝑏=[𝑠(1,1),𝑠(1,2),…,𝑠(1,𝑛),𝑠(2,2),𝑠(2,3),…,𝑠(2,𝑛),𝑠(3,3),…,𝑠(𝑛,𝑛)]
.

For example, if 𝑎=[1,2,5,10]
, then 𝑏=[1,3,8,18,2,7,17,5,15,10]
.

You are given 𝑞
 queries. During the 𝑖
-th query, you are given two integers 𝑙𝑖
 and 𝑟𝑖
, and you have to calculate ∑𝑗=𝑙𝑖𝑟𝑖𝑏𝑗
.

### ideas
1. 假设把s排列成一个矩阵
2. s[0,0], s[0,1], .... s[0.n-1]
3. . s[1,1]......
4. .. .... s[2,2]...
5. .................... s[n-1,n-1]
6. 一个询问 b[l] + ... + b[r]
7. 必然是某行的后缀+m整行+某行的前缀
8. 如果能够快速的计算出某行的和/前缀/后缀
9. 那么问题就解决了
10. s[i,i] + s[i,i+1] + ... s[i...n-1] = 
11.  pref[i+1] - pref[i] + pref[i+2] - pref[i] + ... pref[n] - pref[i]
12. ( pref[i+1] + pref[n]) * (n - i) / 2 - pref[i] * (n - i)
13. 
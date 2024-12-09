The Winter holiday will be here soon. Mr. Chanek wants to decorate his house's wall with ornaments. The wall can be
represented as a binary string 𝑎
of length 𝑛
. His favorite nephew has another binary string 𝑏
of length 𝑚
(𝑚≤𝑛
).

Mr. Chanek's nephew loves the non-negative integer 𝑘
. His nephew wants exactly 𝑘
occurrences of 𝑏
as substrings in 𝑎
.

However, Mr. Chanek does not know the value of 𝑘
. So, for each 𝑘
(0≤𝑘≤𝑛−𝑚+1
), find the minimum number of elements in 𝑎
that have to be changed such that there are exactly 𝑘
occurrences of 𝑏
in 𝑎
.

A string 𝑠
occurs exactly 𝑘
times in 𝑡
if there are exactly 𝑘
different pairs (𝑝,𝑞)
such that we can obtain 𝑠
by deleting 𝑝
characters from the beginning and 𝑞
characters from the end of 𝑡
.

### ideas

1. 考虑一个简单的问题
2. 在给定k的情况下，如何计算最小的修改量
3. 如果t = 1001， s = 1000001001
4. dp[i][j][u] 表示到i为止，有j个子串为t，且最后一段和u匹配的距离是u
5. 如果 s[i+1] == t[u], dp[i][j][u] => dp[i+1][j][u+1]
6. 如果 s[i+1] != t[u], dp[i][j][u] +1 => dp[i+1][j][u+1]
7. 如果 u+1 = m, 状态再转换到 0
8. 如果 s[i+1] != t[u], 还有一个转换是不改变，从而得到 next[u-1]
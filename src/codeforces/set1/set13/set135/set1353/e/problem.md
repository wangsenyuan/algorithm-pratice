You are given a garland consisting of 𝑛
 lamps. States of the lamps are represented by the string 𝑠
 of length 𝑛
. The 𝑖
-th character of the string 𝑠𝑖
 equals '0' if the 𝑖
-th lamp is turned off or '1' if the 𝑖
-th lamp is turned on. You are also given a positive integer 𝑘
.

In one move, you can choose one lamp and change its state (i.e. turn it on if it is turned off and vice versa).

The garland is called 𝑘
-periodic if the distance between each pair of adjacent turned on lamps is exactly 𝑘
. Consider the case 𝑘=3
. Then garlands "00010010", "1001001", "00010" and "0" are good but garlands "00101001", "1000001" and "01001100" are not. Note that the garland is not cyclic, i.e. the first turned on lamp is not going after the last turned on lamp and vice versa.

Your task is to find the minimum number of moves you need to make to obtain 𝑘
-periodic garland from the given one.

You have to answer 𝑡
 independent test cases.

 ### ideas
 1. dp[i][1] 表示第i个点亮时的最小操作数
 2. dp[i][1] = f(i - 1) 前面都被关掉时的操作数
 3.          = dp[i-k+1][1] + g(i-k+2, i-1) 前i-k+1被点亮，且中间都关闭时的cost
 4. 还需要从后往前也计算一遍
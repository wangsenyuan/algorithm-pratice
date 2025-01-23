Polycarp likes numbers that are divisible by 3.

He has a huge number 𝑠
. Polycarp wants to cut from it the maximum number of numbers that are divisible by 3
. To do this, he makes an arbitrary number of vertical cuts between pairs of adjacent digits. As a result, after 𝑚
 such cuts, there will be 𝑚+1
 parts in total. Polycarp analyzes each of the obtained numbers and finds the number of those that are divisible by 3
.

For example, if the original number is 𝑠=3121
, then Polycarp can cut it into three parts with two cuts: 3|1|21
. As a result, he will get two numbers that are divisible by 3
.

Polycarp can make an arbitrary number of vertical cuts, where each cut is made between a pair of adjacent digits. The resulting numbers cannot contain extra leading zeroes (that is, the number can begin with 0 if and only if this number is exactly one character '0'). For example, 007, 01 and 00099 are not valid numbers, but 90, 0 and 10001 are valid.

What is the maximum number of numbers divisible by 3
 that Polycarp can obtain?

### ideas
1. dp[i][x] = 表示到i为止，且余数是0时的最大值
2. dp[i][x] = dp[i-1][(x + s[i]) % 3] (将s[i] 和前面的数进行合并) 
3. dp[i][s[i] % 3] = max(dp[i-1][?]) + 1 (重新开启一段新的区间)
4. 但是0要特殊处理
You have a fence consisting of 𝑛
 vertical boards. The width of each board is 1
. The height of the 𝑖
-th board is 𝑎𝑖
. You think that the fence is great if there is no pair of adjacent boards having the same height. More formally, the fence is great if and only if for all indices from 2
 to 𝑛
, the condition 𝑎𝑖−1≠𝑎𝑖
 holds.

Unfortunately, it is possible that now your fence is not great. But you can change it! You can increase the length of the 𝑖
-th board by 1
, but you have to pay 𝑏𝑖
 rubles for it. The length of each board can be increased any number of times (possibly, zero).

Calculate the minimum number of rubles you have to spend to make the fence great again!

You have to answer 𝑞
 independent queries.

Input
The first line contains one integer 𝑞
 (1≤𝑞≤3⋅105
) — the number of queries.

The first line of each query contains one integers 𝑛
 (1≤𝑛≤3⋅105
) — the number of boards in the fence.

The following 𝑛
 lines of each query contain the descriptions of the boards. The 𝑖
-th line contains two integers 𝑎𝑖
 and 𝑏𝑖
 (1≤𝑎𝑖,𝑏𝑖≤109
) — the length of the 𝑖
-th board and the price for increasing it by 1
, respectively.

It is guaranteed that sum of all 𝑛
 over all queries not exceed 3⋅105
.

### ideas
1. 有点懵呐。
2. 如果 a[i] + 1 = a[i+1], 那么增加 a[i],要么也增加 a[i+1], 要们增加a[i] +2
3. 但是存在增加3的情况吗？
4. a[1]肯定不能增加3， 要么它和后面的不一样（不用增加），要么它和后面的一样，但是增加它更优，只需要+1
5. 假设到i的时候，都不存在增加3的情况
6. 然后考虑i+1，同样的，如果单独考虑i,i+1, i最多增加1，但是因为有可能它已经增加了1（因为前面）所以它可以到2
7. dp[i][0] 表示a[i] 不变时的最小值(当然要满足条件) 1/2表示增高1/2时的最优解
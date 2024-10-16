You are planning to buy an apartment in a 𝑛
-floor building. The floors are numbered from 1
 to 𝑛
 from the bottom to the top. At first for each floor you want to know the minimum total time to reach it from the first (the bottom) floor.

Let:

𝑎𝑖
 for all 𝑖
 from 1
 to 𝑛−1
 be the time required to go from the 𝑖
-th floor to the (𝑖+1)
-th one (and from the (𝑖+1)
-th to the 𝑖
-th as well) using the stairs;
𝑏𝑖
 for all 𝑖
 from 1
 to 𝑛−1
 be the time required to go from the 𝑖
-th floor to the (𝑖+1)
-th one (and from the (𝑖+1)
-th to the 𝑖
-th as well) using the elevator, also there is a value 𝑐
 — time overhead for elevator usage (you need to wait for it, the elevator doors are too slow!).
In one move, you can go from the floor you are staying at 𝑥
 to any floor 𝑦
 (𝑥≠𝑦
) in two different ways:

If you are using the stairs, just sum up the corresponding values of 𝑎𝑖
. Formally, it will take ∑𝑖=𝑚𝑖𝑛(𝑥,𝑦)𝑚𝑎𝑥(𝑥,𝑦)−1𝑎𝑖
 time units.
If you are using the elevator, just sum up 𝑐
 and the corresponding values of 𝑏𝑖
. Formally, it will take 𝑐+∑𝑖=𝑚𝑖𝑛(𝑥,𝑦)𝑚𝑎𝑥(𝑥,𝑦)−1𝑏𝑖
 time units.
You can perform as many moves as you want (possibly zero).

So your task is for each 𝑖
 to determine the minimum total time it takes to reach the 𝑖
-th floor from the 1
-st (bottom) floor.

### ideas
1. dp[i][0] 表示到达第i层时，使用的是stairs
2. dp[i][1] 表示到达第i层时，使用的是evelator
3. dp[i][0] = min(dp[j][0], dp[j][1]) + sum(a[j...i])
4. dp[i][1] = dp[j][1] + sum(b[j...i]) or dp[j][0] + c + sum(b[j...i])
5. dp[i][0] = pref_a(i) - pref_a[j] + dp[j] 也就是 dp[j] - pref_a[j] 最小的那个j
6. dp[i][1] = pref_b(i) - pref_b(j) + min(dp[j][1], dp[j][0] + c) 
7.          = pref_b(i) + min(... - pref_b(j))
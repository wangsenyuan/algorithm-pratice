Welcome to Innopolis city. Throughout the whole year, Innopolis citizens suffer from everlasting city construction.

From the window in your room, you see the sequence of n hills, where i-th of them has height ai. The Innopolis administration wants to build some houses on the hills. However, for the sake of city appearance, a house can be only built on the hill, which is strictly higher than neighbouring hills (if they are present). For example, if the sequence of heights is 5, 4, 6, 2, then houses could be built on hills with heights 5 and 6 only.

The Innopolis administration has an excavator, that can decrease the height of an arbitrary hill by one in one hour. The excavator can only work on one hill at a time. It is allowed to decrease hills up to zero height, or even to negative values. Increasing height of any hill is impossible. The city administration wants to build k houses, so there must be at least k hills that satisfy the condition above. What is the minimum time required to adjust the hills to achieve the administration's plan?

However, the exact value of k is not yet determined, so could you please calculate answers for all k in range ? Here  denotes n divided by two, rounded up.

### ideas
1. dp[i][j] 表示到目前为止已经构建了j个房子时的最小值
2. dp[i][j] = dp[i-1][j] (当前的hill上不建房子)
3.         dp[i-2][j-1] + ..
4. 这里有个问题，就是在当前hill上面建房子，那么在i+1这里是不能建房子的
5. 还有一个问题，就是在i这里建房子的时候，如果i+1太高（或者不高，要降低高度，其实i+2对它也是有影响的）
6. dp[i][j][0] 表示到达i处有j个房子，但在i处没有房子时的状态
7. dp[i][j][1] ...i处有房子时的状态
8. dp[i][j][0] = dp[i-1][j][1] + (如果a[i-1] > a[i] 0, else a[i] - (a[i-1] - 1))
9.             = dp[i-1][j][0] (如果在i-1处也没有房子，对i是没有影响的)
10. dp[i][j][1] = dp[i-1][j][0] + 如果 a[i] > a[i-1] 0, else a[i] - .. 但是这里的问题是 dp[i-1][j][0]其实包含了已经减去的高度
11. 必须考虑两个hill
12. dp[i][j][0] 表示 i-2, i都没有房子时, 
13.  1 只有i处有房子, 2 i-2处有房子， 3 i-2, i处都有房子时的状态
14. dp[i][j][0] = min(dp[i-2][j][0], dp[i-2][j][2])
15. dp[i][j][1] = min(dp[i-2][j-1][0], dp[i-2][j-1][2]) + a[i] - (a[i-1] - 1)
16. dp[i][j][2] = min(dp[i-2][j][3], dp[i-2][j][1]) + a[i-2] - (a[i-1] - 1)
17. dp[i][j][3] = min(dp[i-2][j-1][3], dp[i-2][j-1][1]) + max(a[i], a[i-2]) - (a[i-1] - 1)
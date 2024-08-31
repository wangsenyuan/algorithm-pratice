You are playing one famous sandbox game with the three-dimensional world. The map of the world can be represented as a matrix of size 𝑛×𝑚
, where the height of the cell (𝑖,𝑗)
 is 𝑎𝑖,𝑗
.

You are in the cell (1,1)
 right now and want to get in the cell (𝑛,𝑚)
. You can move only down (from the cell (𝑖,𝑗)
 to the cell (𝑖+1,𝑗)
) or right (from the cell (𝑖,𝑗)
 to the cell (𝑖,𝑗+1)
). There is an additional restriction: if the height of the current cell is 𝑥
 then you can move only to the cell with height 𝑥+1
.

Before the first move you can perform several operations. During one operation, you can decrease the height of any cell by one. I.e. you choose some cell (𝑖,𝑗)
 and assign (set) 𝑎𝑖,𝑗:=𝑎𝑖,𝑗−1
. Note that you can make heights less than or equal to zero. Also note that you can decrease the height of the cell (1,1)
.

Your task is to find the minimum number of operations you have to perform to obtain at least one suitable path from the cell (1,1)
 to the cell (𝑛,𝑚)
. It is guaranteed that the answer exists.

You have to answer 𝑡
 independent test cases.

Input
The first line of the input contains one integer 𝑡
 (1≤𝑡≤100
) — the number of test cases. Then 𝑡
 test cases follow.

The first line of the test case contains two integers 𝑛
 and 𝑚
 (1≤𝑛,𝑚≤100
) — the number of rows and the number of columns in the map of the world. The next 𝑛
 lines contain 𝑚
 integers each, where the 𝑗
-th integer in the 𝑖
-th line is 𝑎𝑖,𝑗
 (1≤𝑎𝑖,𝑗≤1015
) — the height of the cell (𝑖,𝑗)
.

It is guaranteed that the sum of 𝑛
 (as well as the sum of 𝑚
) over all test cases does not exceed 100
 (∑𝑛≤100;∑𝑚≤100
).

Output
For each test case, print the answer — the minimum number of operations you have to perform to obtain at least one suitable path from the cell (1,1)
 to the cell (𝑛,𝑚)
. It is guaranteed that the answer exists.

### ideas
1. dp[i][j][x] 表示达到(i, j)， 且a[i][j] = x时的最小值
2. 如果能这么处理，就简单了。但是问题时a[i][j]可以有1e15, 显然时没法这样定义状态的
3. 但是反过来讲 dp[i][j][x] x = a[0][0] + (i + j) 似乎就简单多了
4. 而且，这里x似乎也不需要记录（因为 a[i][j]必须等于 a[0][0] + i + j
5. 但是这里有个问题是，有可能 a[i][j] < a[0][0] + i + j, 那么此时就必须减少整个路径上的值，以适应a[i][j]
6. 这里有个观察，对于位置(i, j)来说，需要它降低高度，对它有影响的区域，就是它的右下角范围
7. 而这个范围内，需要它减低的数值，不会超过 n * m 个
8. 换句话说，在这个路径上，肯定有一个点是不用减小的
9. 假设这个条件不成立，也就是说，在所有的路径上都必须至少减少1，但是显然大家都少减少1，可以得到一个更优的结果
10. 所以，我们可以固定一个位置(r, c)，它是不能减少的，在它不能减少的情况，向前移动，向后因动，从而获得最小值
11. 如果每个位置必须影响到(r, c)， 那么这个变化就是无效的
12. 如果这个位置是在（r, c)后面， 那么就是 a[i][j] < a[r][c] + i - r + j - c
13. 那么这样子，可以 n * m * n * m = 1e8; 似乎可行
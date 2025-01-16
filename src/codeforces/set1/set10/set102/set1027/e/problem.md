You are given a square board, consisting of 𝑛
 rows and 𝑛
 columns. Each tile in it should be colored either white or black.

Let's call some coloring beautiful if each pair of adjacent rows are either the same or different in every position. The same condition should be held for the columns as well.

Let's call some coloring suitable if it is beautiful and there is no rectangle of the single color, consisting of at least 𝑘
 tiles.

Your task is to count the number of suitable colorings of the board of the given size.

Since the answer can be very large, print it modulo 998244353
.

### ideas
1. 如果 k <= n * m / 2 => 0 (无论怎么样，肯定有一半的会是，要么黑色，要么白色)
2. 第一行确定后，其他的所有的行，要么和第一行相同，要么和第一行（完全）不一样
3. 问题是列怎么搞？
4. 101000
5. ...
6. 第一列和第三列，肯定要一样(似乎用上面的策略，对于列也是必然满足的)
7. 考虑 10001010 (第一行)
8. 后面能够使用的行 10001010 或者 01110101
9. 如果是这样，那么就考虑k的影响
10. 一行中有x个1，n-x的0，这样的C(n, x) (n个位置中，选择x个1)
11. 然后考虑有a行和一致的，n-b行不一致的，a * x <= k and (n - b) * (n - x) <= k成立的时候，把结果加进去
12. 
Dima the hamster enjoys nibbling different things: cages, sticks, bad problemsetters and even trees!

Recently he found a binary search tree and instinctively nibbled all of its edges, hence messing up the vertices. Dima knows that if Andrew, who has been thoroughly assembling the tree for a long time, comes home and sees his creation demolished, he'll get extremely upset.

To not let that happen, Dima has to recover the binary search tree. Luckily, he noticed that any two vertices connected by a direct edge had their greatest common divisor value exceed 1
.

Help Dima construct such a binary search tree or determine that it's impossible. The definition and properties of a binary search tree can be found here.

### ideas
1. sort a[i] first
2. 如果 a[i]是root， 那么 dp[i...i-1] d[i+1...n]必须是ok的
3. 如果要构造l....r, 那么选择 i为root 
4. dp[l..i-1], dp[i+1...r] and gcd(a[i], ?) > 1
5. 也就是说需要知道 dp[l...r][i] 这样子 700 * 700 * 700 = 300 * 1e6
6. 似乎有点悬呐, 如果选择了a[i], 那么左右节点的root， 肯定是和gcd(a[i], ?) > 1的那些
7. 那么a[i]的质数因子，应该要包括进去，所以，可以使用 dp[i][j][x] (这里x是root的质数因子)
8. 这样的数不会超过30个， 700 * 700 * 30 = 47 * 3 * 1e5
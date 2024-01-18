You are given an integer 𝑛
and 𝑘
intervals. The 𝑖
-th interval is [𝑙𝑖,𝑟𝑖]
where 1≤𝑙𝑖≤𝑟𝑖≤𝑛
.

Let us call a regular bracket sequence†,‡
of length 𝑛
hyperregular if for each 𝑖
such that 1≤𝑖≤𝑘
, the substring 𝑠𝑙𝑖𝑠𝑙𝑖+1…𝑠𝑟𝑖⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯
is also a regular bracket sequence.

Your task is to count the number of hyperregular bracket sequences. Since this number can be really large, you are only
required to find it modulo 998244353
.

†
A bracket sequence is a string containing only the characters "(" and ")".

‡
A bracket sequence is called regular if one can turn it into a valid math expression by adding characters + and 1. For
example, sequences (())(), (), (()(())) and the empty string are regular, while )(, ((), and (()))( are not.

### thoughts

1. 显然n和区间的长度都必须是偶数
2. 对于区间[l..r]来说，要求它也是一个regular的，如果单独处理它的话，似乎也挺难的，先假设有办法处理
3. 如果两个区间没有重叠，那就单独处理后，再相乘
4. 如果有重叠，有两种情况，完全包含，那么相当于分成两块，内部和剩余的部分分别要regular
5. 有部分重叠呢？分成3段吗？有没有可能中间的不是regular，但是合并在一起后regular了？
6. 不可能，假设中加的非regular，左括号比右括号多x个，那么对于以它为suffix的那个来说，它前面部分，必须多x个右括号，这个将造成它不是一个regular
7. 所以，分成了3部分
8. 最后的结果就是分成了非重叠的和最后剩余的一个部分
9. 现在计算对于给定长度的n，如何保证它是regular的计数？
10. dp[n] = ？
11. dp[0] = 1, dp[2] = 1, dp[4] = 2 ()() (())
12. dp[n + 2] = 先放左边的括号 这里一共有n + 1个位置, 然后放右括号的位置, 需要将右括号放置在左括号的后面
13. 不对。这样子要求 dp[n]必须是regular的，但是对于n+2来说，n个长度不一定是regular的
14. 左括号+1, 右括号-1, 那么就是要 pref[n] = 0
15. dp[i][x] = 在第i位时sum = x的计数
16. dp[i+1][x+1] = dp[i][x] + dp[i][x+2]
17. 这样子是 n * n 的， pass
18. 把左括号看成向右移动，右括号看做向下移动，不考虑其他的限制，就是从(1, 1) 移动到 (n, n)的距离
19. 这个答案是C(n + n, n)
20. 但是有个限制是必须只能在上半区域移动，这个公式我看到过 Catalan number
21. C_n = (2n)! / ((n + 1)! * n!)
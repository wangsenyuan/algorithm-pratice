You are given a binary string 𝑠
(a binary string is a string consisting of characters 0 and/or 1).

Let's call a binary string balanced if the number of subsequences 01 (the number of indices 𝑖
and 𝑗
such that 1≤𝑖<𝑗≤𝑛
, 𝑠𝑖=0
and 𝑠𝑗=1
) equals to the number of subsequences 10 (the number of indices 𝑘
and 𝑙
such that 1≤𝑘<𝑙≤𝑛
, 𝑠𝑘=1
and 𝑠𝑙=0
) in it.

For example, the string 1000110 is balanced, because both the number of subsequences 01 and the number of subsequences
10 are equal to 6
. On the other hand, 11010 is not balanced, because the number of subsequences 01 is 1
, but the number of subsequences 10 is 5
.

You can perform the following operation any number of times: choose two characters in 𝑠
and swap them. Your task is to calculate the minimum number of operations to make the string 𝑠
balanced.

### thoughts

1. 交换(0, 0) (1, 1)没有意义，
2. 只能交换 (0, 1) 或者 (1, 0)
3. 假设交换了(i, j) 且 s[i] = 0, s[j] = 1, 那么总的balance变化为 -suf[i][1] - pref[j][0] + suf[i][0] + pref[j][1]
4. suf[i][0] - suf[i][1] + pref[j][1] - pref[j][0]
5. 假设处理(i, j) pair 似乎从外向那处理更合理，在这种情况下，后缀里面1和0的个数，以及前缀中1, 0的个数，都是可以确定袭来的
6. 在前后缀里面总的1个数量是定了，cnt1, 如果后缀中有 s1个，那么前缀中p1 = cnt1 - s1, 进而p0 = i - p1, s0 = j - s1
7. 那么看起来是 n * n * n 的复杂性
8. dp[i][j][x] 表示处理到[i,j]时后缀有x个1时的最小交换数
9. 但怎么判断dp[i][j][x]是平衡的呢？
10. 不可能要求每一步都是平衡的吧？因为有些步骤得到的结果就是不平衡的
11. dp[i][j][x] = Pair{平衡度, 最小的交换次数}
12. 这里还有个信息被隐藏了。
13. 在确定j的后面1的数量的情况下， 交换(i, j) 能够带来的balance的变化是确定的
14. 假设一开始的imbalance是x
15. 那是不是找到一个序列 (i0, j0) (i1, j1) ... (ik, jk) 最小的k使得 这个sum of d[i][j][?] = -x ?

### solution

Let's calculate the following dynamic programming: 𝑑𝑝𝑖,𝑐𝑛𝑡0,𝑐𝑛𝑡01
— the minimum number of changes in string 𝑠
if we consider only 𝑖
first characters of it, the number of characters 0 on that prefix is 𝑐𝑛𝑡0
, and the number of subsequences 01 on that prefix is equal to 𝑐𝑛𝑡01
.

The transitions are pretty simple. Let's look at the transitions according to the character we are trying to place at
the next position:

if it is 0, then there is transition from the state (𝑖,𝑐𝑛𝑡0,𝑐𝑛𝑡01)
to the state (𝑖+1,𝑐𝑛𝑡0+1,𝑐𝑛𝑡01)
and the value of 𝑑𝑝
depends on 𝑠𝑖
(the value stays the same if 𝑠𝑖=0
, and increases by 1
otherwise);
if it is 1, then there is transition from the state (𝑖,𝑐𝑛𝑡0,𝑐𝑛𝑡01)
to the state (𝑖+1,𝑐𝑛𝑡0,𝑐𝑛𝑡01+𝑐𝑛𝑡0)
, and the value of 𝑑𝑝
depends on 𝑠𝑖
(the value stays the same if 𝑠𝑖=1
and increases by 1
otherwise).
So, this dynamic programming works in 𝑂(𝑛4)
.

It remains us to get the answer to the problem from that dynamic programming. It is stored in 𝑑𝑝𝑛,𝑐𝑛𝑡0,𝑛𝑒𝑒𝑑
, where 𝑐𝑛𝑡0
is equal to the number of characters 0 in the string 𝑠
, and 𝑛𝑒𝑒𝑑=𝑐𝑛𝑡0⋅𝑐𝑛𝑡12
(because the number of subsequences 01 should be equal to the number of subsequences 10). But our dynamic programming
stores the number of changes in the string, and the problems asks for the minimum number of swaps. However, we can
easily get it from the 𝑑𝑝
value. Since the amounts of zeroes and ones are fixed in the string, then the number of changes from 0 to 1 equals to
the number of changes from 1 to 0 and we can pair them up. So, the answer to the problem is the half of the 𝑑𝑝
value.
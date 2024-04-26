You are given an array 𝑎
of length 2𝑛
. You should process 𝑞
queries on it. Each query has one of the following 4
types:

𝑅𝑒𝑝𝑙𝑎𝑐𝑒(𝑥,𝑘)
— change 𝑎𝑥
to 𝑘
;
𝑅𝑒𝑣𝑒𝑟𝑠𝑒(𝑘)
— reverse each subarray [(𝑖−1)⋅2𝑘+1,𝑖⋅2𝑘]
for all 𝑖
(𝑖≥1
);
𝑆𝑤𝑎𝑝(𝑘)
— swap subarrays [(2𝑖−2)⋅2𝑘+1,(2𝑖−1)⋅2𝑘]
and [(2𝑖−1)⋅2𝑘+1,2𝑖⋅2𝑘]
for all 𝑖
(𝑖≥1
);
𝑆𝑢𝑚(𝑙,𝑟)
— print the sum of the elements of subarray [𝑙,𝑟]
.
Write a program that can quickly process given queries.

### ideas

### solution

如何翻转一个数组？
对于如下数组，让我们换一个视角：
12345678
首先，交换相邻数字：
21436587
然后，两个数一组，交换：
43218765
最后，四个数一组，交换：
87654321

注：用这个思路可以解决

190. 颠倒二进制位

所以操作 2 可以用操作 3 表达：
调用 "2 k"，相当于调用 "3 0", "3 1", "3 2",..., "3 k-1"。

对于每个 k，记录 "3 k" 的调用次数。偶数次相当于不变。
注意，只记录次数，不修改数组。

现在剩下操作 1 4，单点修改区间查询，用线段树解决。
注意操作 2 和操作 3 的 k 越大，每个区间越长，所以 k 越大越靠近线段树的根节点。

怎么把操作 3 加进去？或者说，操作 3 对操作 1 和操作 4 有何影响？
如果线段树第 d 层执行了奇数次操作 3，那么原来要递归左子树的逻辑，就要变成递归右子树了；原来要递归右子树的逻辑，就要变成递归左子树了。

代码
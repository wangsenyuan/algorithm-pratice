You are exploring the stunning region of Natlan! This region consists of 𝑛
 cities, and each city is rated with an attractiveness 𝑎𝑖
. A directed edge exists from City 𝑖
 to City 𝑗
 if and only if 𝑖<𝑗
 and gcd(𝑎𝑖,𝑎𝑗)≠1
, where gcd(𝑥,𝑦)
 denotes the greatest common divisor (GCD) of integers 𝑥
 and 𝑦
.

Starting from City 1
, your task is to determine the total number of distinct paths you can take to reach City 𝑛
, modulo 998244353
. Two paths are different if and only if the set of cities visited is different.


### ideas
1. 从前往后考虑， 到达i的时候， 如果 i < j 且 gcd(a[i], a[j]) > 1, 那么j要增加i的ways
2. 当然没法一个个检查j
3. 假设a[i]的素数因子有h个（不超过30个）
4. 然后它应该去更新i+1范围的所有的值
5. 这个不对，考虑2个节点1，2, 且它们的值相同为6
6. 答案显然为1，但是如果用我之前的想法，那么答案就变成2了
7. 错在哪里呢？
8. 反过来考虑从n到i；假设处理i，它的值是a[i]， 此时如果知道到所有值为x的计数
9. 那么dp[i] = sum(fp[x]) where gcd(x, a[i]) > 1
10. 但是迭代x(1...最大值，肯定不行)
11. 这个感觉要用容斥原理。
12. 
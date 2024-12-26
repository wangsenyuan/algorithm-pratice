You have a long fence which consists of 𝑛
 sections. Unfortunately, it is not painted, so you decided to hire 𝑞
 painters to paint it. 𝑖
-th painter will paint all sections 𝑥
 such that 𝑙𝑖≤𝑥≤𝑟𝑖
.

Unfortunately, you are on a tight budget, so you may hire only 𝑞−2
 painters. Obviously, only painters you hire will do their work.

You want to maximize the number of painted sections if you choose 𝑞−2
 painters optimally. A section is considered painted if at least one painter paints it.

Input
The first line contains two integers 𝑛
 and 𝑞
 (3≤𝑛,𝑞≤5000
) — the number of sections and the number of painters availible for hire, respectively.

Then 𝑞
 lines follow, each describing one of the painters: 𝑖
-th line contains two integers 𝑙𝑖
 and 𝑟𝑖
 (1≤𝑙𝑖≤𝑟𝑖≤𝑛
).

Output
Print one integer — maximum number of painted sections if you hire 𝑞−2
 painters.


 ### ideas
 1. q * q * lg(n) 的算法是有的。有没有更好的？
 2. 如果存在这样的情况，就是 a 完全包含b的区间，那么把b去掉，肯定不会更坏 
 3. 如果fire了至少2人，直接计算剩下的覆盖区间即可
 4. 如果fire了1人，那么可以用 q * lg(n)的方式
 5. 如果fire了0人,现在剩下的就是，要么没有重叠的，要么在边界重叠的
 6. 假设要计算只是删除了一个区间的情况
 7. dp[i][0] 很好算
 8. dp[i][1] = dp[i-1][0] max dp[j][0] + 当前的范围（如果j的右端点在当前范围的左端点前面)
 9.        或者 dp[j][0] + r[i] - r[j] (那么可以使用dp[j][0] - r[j] 维护一个seg tree)
10. dp[i][2] = dp[i-1][1] max dp[j][1] + .... 
11. 貌似可以搞～

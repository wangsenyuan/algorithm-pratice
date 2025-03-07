Given a tree with 𝑛
 vertices rooted at vertex 1
. While walking through it with her cat Chefir, Sakurako got distracted, and Chefir ran away.

To help Sakurako, Kosuke recorded his 𝑞
 guesses. In the 𝑖
-th guess, he assumes that Chefir got lost at vertex 𝑣𝑖
 and had 𝑘𝑖
 stamina.

Also, for each guess, Kosuke assumes that Chefir could move along the edges an arbitrary number of times:

from vertex 𝑎
 to vertex 𝑏
, if 𝑎
 is an ancestor∗
 of 𝑏
, the stamina will not change;
from vertex 𝑎
 to vertex 𝑏
, if 𝑎
 is not an ancestor of 𝑏
, then Chefir's stamina decreases by 1
.
If Chefir's stamina is 0
, he cannot make a move of the second type.

For each assumption, your task is to find the distance to the farthest vertex that Chefir could reach from vertex 𝑣𝑖
, having 𝑘𝑖
 stamina.

 ### ideas
 1. 这个条件有点难理解，就是b可以是任何一个点吗？还是只能是邻居的？
 2. 感觉应该只能是邻居，否则只需要花费最多1个能量，就移动到任何位置去了
 3. 对于v来说，如果离他最远的点在它的子树中，那么花费是0
 4. 否则的话，就要往上移动
 5. 所以，对于任何一个点，需要保留两个（子树中）最远的点的信息
 6. 如果到达了点u，那么就可以算出到达点u后，能够到达的最远的点（如果是从第一个子树中来的，就是第二个最远点，否则就是第一个最远点）
 7. 然后就是判断，在给定v，和k的情况下，需要到达的u是哪里
 8. 不考虑v中的情况，假设另外一个点是w, dep[v] - dep[u] (u是v, w的lca) <= k
 9. dep[v] - dp[u] + dp[w] - dp[u] = dp[v] + dp[w] - 2 * dp[u]
 10. 在v确定的情况下，找到最大的 dp[w] - 2 * dp[u]即可
 11. 但是，如何排除掉v和v相同子树中的情况呢？
 12. 
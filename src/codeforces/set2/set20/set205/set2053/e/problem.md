There is a tree consisting of 𝑛
 vertices. Let a caterpillar be denoted by an integer pair (𝑝,𝑞)
 (1≤𝑝,𝑞≤𝑛
, 𝑝≠𝑞
): its head is at vertex 𝑝
, its tail is at vertex 𝑞
, and it dominates all the vertices on the simple path from 𝑝
 to 𝑞
 (including 𝑝
 and 𝑞
). The caterpillar sequence of (𝑝,𝑞)
 is defined as the sequence consisting only of the vertices on the simple path, sorted in the ascending order of the distance to 𝑝
.

Nora and Aron are taking turns moving the caterpillar, with Nora going first. Both players will be using his or her own optimal strategy:

They will play to make himself or herself win;
However, if it is impossible, they will play to prevent the other person from winning (thus, the game will end in a tie).
In Nora's turn, she must choose a vertex 𝑢
 adjacent to vertex 𝑝
, which is not dominated by the caterpillar, and move all the vertices in it by one edge towards vertex 𝑢
∗
. In Aron's turn, he must choose a vertex 𝑣
 adjacent to vertex 𝑞
, which is not dominated by the caterpillar, and move all the vertices in it by one edge towards vertex 𝑣
. Note that the moves allowed to the two players are different.

Whenever 𝑝
 is a leaf†
, Nora wins‡
. Whenever 𝑞
 is a leaf, Aron wins. If either initially both 𝑝
 and 𝑞
 are leaves, or after 10100
 turns the game has not ended, the result is a tie.

Please count the number of integer pairs (𝑝,𝑞)
 with 1≤𝑝,𝑞≤𝑛
 and 𝑝≠𝑞
 such that, if the caterpillar is initially (𝑝,𝑞)
, Aron wins the game.

### ideas
1. 任何两个内部节点，距离leaf超过2的（也就是 dist[u] > 1, dist[v] > 1) 的组合，是不是肯定是tie？
2. 因为Nora往前拖一下，Aron就可以往后拖一步，从而回到原来的状态
3. 如果是这样，那么只有一种情况Aron可以赢，就是q点本来是一个leaf。
4. 如果是这样，岂不是太简单了？
5. 所有距离 >= 2 的内部节点作为p点，所有叶子节点作为q
6. 上面的分析是不对的
7. 除了这些，还有一类是，q是叶子节点，但是p不是的
8. 这里漏了一种case，就是如果p, q都在内部节点，但是p移动一次后，q 却接近了一个叶子节点
This is an interactive problem.

Little Dormi was faced with an awkward problem at the carnival: he has to guess the edges of an unweighted tree of 𝑛
 nodes! The nodes of the tree are numbered from 1
 to 𝑛
.

The game master only allows him to ask one type of question:

Little Dormi picks a node 𝑟
 (1≤𝑟≤𝑛
), and the game master will reply with an array 𝑑1,𝑑2,…,𝑑𝑛
, where 𝑑𝑖
 is the length of the shortest path from node 𝑟
 to 𝑖
, for all 1≤𝑖≤𝑛
.
Additionally, to make the game unfair challenge Little Dormi the game master will allow at most ⌈𝑛2⌉
 questions, where ⌈𝑥⌉
 denotes the smallest integer greater than or equal to 𝑥
.

Faced with the stomach-churning possibility of not being able to guess the tree, Little Dormi needs your help to devise a winning strategy!

Note that the game master creates the tree before the game starts, and does not change it during the game.


### ideas
1. 通过2次，可以找到，直径的两个端点
2. 然后，可以找到这个直径 dp[a][i] + dp[b][i] = dia
3. 然后咋办呢？这个直径，有可能占去所有的点，也可能只占3个点
4. 所以不能保证一半的上限
5. 好像有办法了
6. 一开始查 f(1) => d[1], d[2], d[i]
7. 那么可以肯定的是，所以d[i] = 1 的是，连接到1的，
8. 然后查询d[i] = 2的x，那么可以肯定的是，那么x的parent，必然是某个d[j] = 1的，且这次也是1的
9. 其他的d[k] = 1 的，那些是 x的子代，且这次 d[i] = 2 的，仍然是判断不出来的
10. 继续查找那些，不确定的，但是d[i] = 2 的部分(并记录它的parent = 1)的那些
11. 因为，每次都跳过了， d[i] = 1 的部分， 所以，只有一半的会被查询到 
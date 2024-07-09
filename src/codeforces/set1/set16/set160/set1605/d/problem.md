Eikooc and Sushi play a game.

The game is played on a tree having 𝑛
 nodes numbered 1
 to 𝑛
. Recall that a tree having 𝑛
 nodes is an undirected, connected graph with 𝑛−1
 edges.

They take turns alternately moving a token on the tree. Eikooc makes the first move, placing the token on any node of her choice. Sushi makes the next move, followed by Eikooc, followed by Sushi, and so on. In each turn after the first, a player must move the token to a node 𝑢
 such that

𝑢
 is adjacent to the node 𝑣
 the token is currently on
𝑢
 has not been visited before
𝑢⊕𝑣≤𝑚𝑖𝑛(𝑢,𝑣)
Here 𝑥⊕𝑦
 denotes the bitwise XOR operation on integers 𝑥
 and 𝑦
.

Both the players play optimally. The player who is unable to make a move loses.

The following are examples which demonstrate the rules of the game.

Suppose Eikooc starts the game by placing the token at node 4
. Sushi then moves the token to node 6
, which is unvisited and adjacent to 4
. It also holds that 6⊕4=2≤𝑚𝑖𝑛(6,4)
. In the next turn, Eikooc moves the token to node 5
, which is unvisited and adjacent to 6
. It holds that 5⊕6=3≤𝑚𝑖𝑛(5,6)
. Sushi has no more moves to play, so she loses.	
Suppose Eikooc starts the game by placing the token at node 3
. Sushi moves the token to node 2
, which is unvisited and adjacent to 3
. It also holds that 3⊕2=1≤𝑚𝑖𝑛(3,2)
. Eikooc cannot move the token to node 6
 since 6⊕2=4≰𝑚𝑖𝑛(6,2)
. Since Eikooc has no moves to play, she loses.
Before the game begins, Eikooc decides to sneakily relabel the nodes of the tree in her favour. Formally, a relabeling is a permutation 𝑝
 of length 𝑛
 (sequence of 𝑛
 integers wherein each integer from 1
 to 𝑛
 occurs exactly once) where 𝑝𝑖
 denotes the new numbering of node 𝑖
.

She wants to maximize the number of nodes she can choose in the first turn which will guarantee her a win. Help Eikooc find any relabeling which will help her do so.

Input
The first line contains a single integer 𝑡 (1≤𝑡≤105)
  — the number of test cases. The description of each test case is as follows.

The first line of each test case contains an integer 𝑛 (1≤𝑛≤2⋅105)
  — the number of nodes in the tree.

The next 𝑛−1
 lines contain two integers 𝑢
 and 𝑣
 (1≤𝑢,𝑣≤𝑛;𝑢≠𝑣)
  — denoting an edge between nodes 𝑢
 and 𝑣
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 2⋅105
.

Output
For each test case print any suitable relabeling  — a permutation of length 𝑛
 which maximizes the number of nodes that can be chosen in the first turn that guarantee a win for Eikooc. If there are multiple such relabelings, you may print any of them.


 ### ideas
 1. 考虑什么情况下， 当前player可以移动token？
 2. 假设(u, v)是两个节点的编号，那么u ^ v <= min(u, v) 就可以移动，否则不能移动
 3. 不妨假设u < v, 那么就是 u ^ v < u
 4. 如果v是2的指数次，1, 2, 4, 8... 那么 u < v 的情况下，不存在 u ^ v < u
 5. 如果 u > v, 那么u最高位，必须和v是相同的，这样子，u ^ v < v 是成立的
 6. 我们考虑u和v的最高位不同，那么 u ^ v 的结果中，更高的那个的最高位，会被保留下来
 7. 所以，这个结果肯定会大于另外一个较小的值 => u和v的最高位必须是一致的
 8. 那么这样子 (u ^ v) 最高位没有了, u ^ v < min(u, v)
 9. 所以，可以得出结论是，任何相邻的两个节点，u,v的最高位一致，就可以移动
 10. 那么最高位一致的，个数分别是 1, 2, 4, 8....
 11. 是为了最大化Eikooc能够获胜的节点数
 12. 考虑2个节点，(1, 2, 3), 如果把1放在中间，那么Eikooc就可以在所有节点获胜
 13. 如果是7个节点， 1, 2, 4, 且它们就是一条直线，4, 2 5, 3, 6, 1, 7 也可以全部获胜
 14.  也就是说，对于任何一个节点，如果和它相临的节点的最高位不一样，它就可以获胜
 15.  假设有任何一个，那么eikooc只有这条路径一半的节点，可以获胜
 16.  如果没有节点的度数 >= 最多的那个（比如4）那么eikooc就可以在全部节点获胜
 17.  如果存在这样一个点，eikooc肯定会失去一些点，那么就不妨把最多的那个给分配掉，那么只会失去一个点
 18.  如果这个节点的度数 < 4, 那么就把这一圈都赋值为4
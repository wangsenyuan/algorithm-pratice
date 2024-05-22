Vladislav has a son who really wanted to go to MIT. The college dormitory at MIT (Moldova Institute of Technology) can be represented as a tree with 𝑛
 vertices, each vertex being a room with exactly one student. A tree is a connected undirected graph with 𝑛
 vertices and 𝑛−1
 edges.

Tonight, there are three types of students:

students who want to party and play music (marked with 𝙿
),
students who wish to sleep and enjoy silence (marked with 𝚂
), and
students who don't care (marked with 𝙲
).
Initially, all the edges are thin walls which allow music to pass through, so when a partying student puts music on, it will be heard in every room. However, we can place some thick walls on any edges — thick walls don't allow music to pass through them.

The university wants to install some thick walls so that every partying student can play music, and no sleepy student can hear it.

Because the university lost a lot of money in a naming rights lawsuit, they ask you to find the minimum number of thick walls they will need to use.

###
1. normally it is a dp problem
2. C的是不是可以直接删除掉？
3. 如果一个是C，那么就把它的所有的child，连接到它的parent上面去 （选择一个非C的root）
4. 这样是不对的
5. 因为如果存在一个C，通过建一个wall，就可以保护C种的那些slice的人，不被这个子树外面的P所影响，
6. 但是通过上面的压缩，这个方式就行不通了
7. 那这里就存在这样一种情况
8. 如果对于子树x，如果它里面存在人喜欢silience，但是子树外面有人在party，那么通过在x的edge上隔绝，就可以保护这颗子树
9. dp[u][0] 表示在子树u的入边不隔绝，
10. dp[u][1]表示隔绝u后，u的情况
11. dp[u][0] = 这个不大对
12. 从外围往里面处理，假设现在遇到了一个P，如果直到遇到一个S
13. 感觉是要组团？
14. P - C - P
15. S - C - S
16. 好像还是得回到上面的方式，但是连接的时候，要增加一个额外的属性 x表明是从哪个子节点连接的
17. 现在整个图里面只剩下 S和P了，那就在它们中间加搁板就可以了
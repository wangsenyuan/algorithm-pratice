This is an interactive problem.

Consider an undirected connected graph consisting of 𝑛
 vertices and 𝑚
 edges. Each vertex can be colored with one of three colors: 1
, 2
, or 3
. Initially, all vertices are uncolored.

Alice and Bob are playing a game consisting of 𝑛
 rounds. In each round, the following two-step process happens:

Alice chooses two different colors.
Bob chooses an uncolored vertex and colors it with one of the two colors chosen by Alice.
Alice wins if there exists an edge connecting two vertices of the same color. Otherwise, Bob wins.

You are given the graph. Your task is to decide which player you wish to play as and win the game.

### ideas
1. 判断alice有没有必胜的策略（或者在什么情况下必胜）
2. 如果存在一条线，连接了两个相同颜色的节点，那么alice win；else bob win
3. 如果某个节点，它有至少3个邻居，那么alice必胜（因为另外3个节点，最多有3种着色，中间这个节点肯定可以找到一个同色的邻居）
4. 如果是二部图，那么bob win（只需要两个颜色即可）
5. 
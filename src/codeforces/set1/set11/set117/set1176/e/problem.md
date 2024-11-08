You are given an undirected unweighted connected graph consisting of 𝑛
 vertices and 𝑚
 edges. It is guaranteed that there are no self-loops or multiple edges in the given graph.

Your task is to choose at most ⌊𝑛2⌋
 vertices in this graph so each unchosen vertex is adjacent (in other words, connected by an edge) to at least one of chosen vertices.

It is guaranteed that the answer exists. If there are multiple answers, you can print any.

You will be given multiple independent queries to answer.

### ideas
1. 对这个图，进行二分染色
2. 如果没有那种必须染成2种颜色的，已经ok了
3. 如果存在某个节点u，它被同时染成了红色、绿色， 这些节点也排除掉即可
4. 因为，它能变成杂色，意味着它同时有红色、绿色的邻居，红色的已经被选中了，但是绿色，肯定也有红色的邻居
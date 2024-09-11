Gildong was hiking a mountain, walking by millions of trees. Inspired by them, he suddenly came up with an interesting idea for trees in data structures: What if we add another edge in a tree?

Then he found that such tree-like graphs are called 1-trees. Since Gildong was bored of solving too many tree problems, he wanted to see if similar techniques in trees can be used in 1-trees as well. Instead of solving it by himself, he's going to test you by providing queries on 1-trees.

First, he'll provide you a tree (not 1-tree) with 𝑛
 vertices, then he will ask you 𝑞
 queries. Each query contains 5
 integers: 𝑥
, 𝑦
, 𝑎
, 𝑏
, and 𝑘
. This means you're asked to determine if there exists a path from vertex 𝑎
 to 𝑏
 that contains exactly 𝑘
 edges after adding a bidirectional edge between vertices 𝑥
 and 𝑦
. A path can contain the same vertices and same edges multiple times. All queries are independent of each other; i.e. the added edge in a query is removed in the next query.

Input
The first line contains an integer 𝑛
 (3≤𝑛≤105
), the number of vertices of the tree.

Next 𝑛−1
 lines contain two integers 𝑢
 and 𝑣
 (1≤𝑢,𝑣≤𝑛
, 𝑢≠𝑣
) each, which means there is an edge between vertex 𝑢
 and 𝑣
. All edges are bidirectional and distinct.

Next line contains an integer 𝑞
 (1≤𝑞≤105
), the number of queries Gildong wants to ask.

Next 𝑞
 lines contain five integers 𝑥
, 𝑦
, 𝑎
, 𝑏
, and 𝑘
 each (1≤𝑥,𝑦,𝑎,𝑏≤𝑛
, 𝑥≠𝑦
, 1≤𝑘≤109
) – the integers explained in the description. It is guaranteed that the edge between 𝑥
 and 𝑦
 does not exist in the original tree.

Output
For each query, print "YES" if there exists a path that contains exactly 𝑘
 edges from vertex 𝑎
 to 𝑏
 after adding an edge between vertices 𝑥
 and 𝑦
. Otherwise, print "NO".

You can print each letter in any case (upper or lower).
Jamie has recently found undirected weighted graphs with the following properties very interesting:

The graph is connected and contains exactly n vertices and m edges.
All edge weights are integers and are in range [1, 109] inclusive.
The length of shortest path from 1 to n is a prime number.
The sum of edges' weights in the minimum spanning tree (MST) of the graph is a prime number.
The graph contains no loops or multi-edges.
If you are not familiar with some terms from the statement you can find definitions of them in notes section.

Help Jamie construct any graph with given number of vertices and edges that is interesting!

###          

1. 如果 n - 1 是 prime number
2. 那么mst的每条边都可以是1， 然后再增加额外的2的边即可
3. 就把n个节点组成一条线，然后找出第一个 >= n 的质数
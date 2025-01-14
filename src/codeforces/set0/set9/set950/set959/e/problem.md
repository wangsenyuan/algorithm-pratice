Ehab is interested in the bitwise-xor operation and the special graphs. Mahmoud gave him a problem that combines both. He has a complete graph consisting of n vertices numbered from 0 to n - 1. For all 0 ≤ u < v < n, vertex u and vertex v are connected with an undirected edge that has weight  (where  is the bitwise-xor operation). Can you find the weight of the minimum spanning tree of that graph?

You can read about complete graphs in https://en.wikipedia.org/wiki/Complete_graph

You can read about the minimum spanning tree in https://en.wikipedia.org/wiki/Minimum_spanning_tree

The weight of the minimum spanning tree is the sum of the weights on the edges included in it.


### ideas
1. hard ~~~
2. (2, 3), (4, 5), (6, 7), ... (2 * n, 2 * n + 1)它们都可以通过 1的边连接在一起
3. (2 ^ 4 = 3 ^ 5 = 6)
4. 感觉是最低位为1的，通过1的cost挂到偶数上，然后最低位为2的(6, 10, 14)等，通过2挂到最近的(4, 8, 12)等上面
5. 1100   1110
PROBLEM
Given that there exists an undirected weighted connected graph with NN vertices, find the total weight of its Minimum Spanning Tree.

You can perform 2 kinds of queries:

Provide 2 non-empty, non-intersectin sets of vertices AA and BB as input to the judge and get the minimum weighted edge between any pair of nodes u \subseteq Au⊆A and v \subseteq Bv⊆B. If there is no such edge, the judge returns -1.
Send the weight of the MST to the judge.
The cost of performing a query is |A|∣A∣. The max cost that can be incurred is 10^410 
4
 .

Constraints:
2 \le N \le 10002≤N≤1000

1 \le weight_{edge} \le 10^51≤weight 
edge
​	
 ≤10 
5
 

EXPLANATION
Will Kruskal’s/Prim’s Algorithm work?
Primitive Kruskal / Prim will not work. Kruskal’s wont work because we will perform N^2N 
2
  queries - once for each pair of nodes, where the cost for each query will be 11.

For Prim’s the set of nodes you know the answer for, will keep on increasing. Hence, the cost for the queries would be something like: 1, 2, 3, ... , N/2, N/2, (N/2 - 1), ..., 3, 2, 11,2,3,...,N/2,N/2,(N/2−1),...,3,2,1. Hence, the cost would be of the order O(N^2)O(N 
2
 ).

Approach
Find the nearest node of each node uu by using a query of the format: [u][u] vs. all except uu.

This gives us connected components of size 22.
In increasing order of component size, we can continue to perform queries like: [component] vs. [all except component].

You can observe that after every query, the component’s size gets doubled.


inpMH34.png
1138×584
Since the size of a component always gets doubled, we can say that every vertex occurs as a vertex in AA at most log(N)log(N) times. Components can be maintained using a DSU.

Each vertex contributes 11 to the net cost. Hence, the overall cost is bounded by O(N * log(N))O(N∗log(N)).

Since the value of NN is at most 10001000, our bound ensures that the max cost of 10^410 
4
  is never exceeded.


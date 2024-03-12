You are currently researching a graph traversal algorithm called the Breadth First Search (BFS). Suppose you have an
input graph of 𝑁
nodes (numbered from 1
to 𝑁
). The graph is represented by an adjacency matrix 𝑀
, for which node 𝑢
can traverse to node 𝑣
if 𝑀𝑢,𝑣
is 1
, otherwise it is 0
. Your algorithm will output the order the nodes are visited in the BFS. The pseudocode of the algorithm is presented as
follows.

    BFS(M[1..N][1..N]):
        let A be an empty array
        let Q be an empty queue

        append 1 to A
        push 1 to Q

        while Q is not empty:
            pop the front element of Q into u
            for v = 1 to N:
                if M[u][v] == 1 and v is not in A:
                    append v to A
                    push v to Q

        return A

During your research, you are interested in the following problem. Given an array 𝐴
such that 𝐴
is a permutation of 1
to 𝑁
and 𝐴1=1
. How many simple undirected graph with 𝑁
nodes and adjacency matrix 𝑀
such that BFS(𝑀)=𝐴
? Since the answer can be very large, calculate the answer modulo 998244353
.

A simple graph has no self-loop (𝑀𝑖,𝑖=0
for 1≤𝑖≤𝑁
) and there is at most one edge that connects a pair of nodes. In an undirected graph, if node 𝑢
is adjacent to node 𝑣
, then node 𝑣
is also adjacent to node 𝑢
; formally, 𝑀𝑢,𝑣=𝑀𝑣,𝑢
for 1≤𝑢<𝑣≤𝑁
.

Two graphs are considered different if there is an edge that exists in one graph but not the other. In other words, two
graphs are considered different if their adjacency matrices are different.

### ideas

1. has to dp
2. but how?
3. 这个bfs的过程是，level by level的过程
4. 第一层是1，然后第二层，然后第三层，...
5. 假设当前层是d, 当前是i，那么i可以连接到任意d-1层的某个节点上
6. 但也不能完全说是任意，如果它是第一个，那么它确实可以连接到任意一个节点上
7. 如果它是第2个，它只能连接到第一连接的节点的后面的节点上
8. 完全没想法～

### solution

We will try to emulate the behaviour of BFS. We use the dynamic programming DP(i, inQueue, isF irst)
— how many ways to add edges into the graph if:
• the BFS has pushed A[1..i − 1] inside the queue;
• inQueue number of A[1..i − 1] are still inside the queue;
• the front element of the queue has children (if isF irst is true) or not (if isF irst is f alse);
• we only consider adding edges to A[1..i − 1].
The DP transitions are as follows.
• Pop the queue and transition to (i, inQueue − 1, true).
• Enqueue A[i] by connecting it with the vertex that is in front of the queue. This can be done
when the isF irst is true or A[i−1] < A[i]. Additionally, A[i] can be connected with any of the
other vertices in the queue. Hence, there are 2
(inQueue−1) ways multiplied by the transition to
(i + 1, inQueue + 1, f alse).

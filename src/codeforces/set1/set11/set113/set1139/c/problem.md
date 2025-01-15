You are given a tree (a connected undirected graph without cycles) of 𝑛
 vertices. Each of the 𝑛−1
 edges of the tree is colored in either black or red.

You are also given an integer 𝑘
. Consider sequences of 𝑘
 vertices. Let's call a sequence [𝑎1,𝑎2,…,𝑎𝑘]
 good if it satisfies the following criterion:

We will walk a path (possibly visiting same edge/vertex multiple times) on the tree, starting from 𝑎1
 and ending at 𝑎𝑘
.
Start at 𝑎1
, then go to 𝑎2
 using the shortest path between 𝑎1
 and 𝑎2
, then go to 𝑎3
 in a similar way, and so on, until you travel the shortest path between 𝑎𝑘−1
 and 𝑎𝑘
.
If you walked over at least one black edge during this process, then the sequence is good.
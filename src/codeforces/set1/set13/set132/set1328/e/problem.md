You are given a rooted tree consisting of 𝑛
 vertices numbered from 1
 to 𝑛
. The root of the tree is a vertex number 1
.

A tree is a connected undirected graph with 𝑛−1
 edges.

You are given 𝑚
 queries. The 𝑖
-th query consists of the set of 𝑘𝑖
 distinct vertices 𝑣𝑖[1],𝑣𝑖[2],…,𝑣𝑖[𝑘𝑖]
. Your task is to say if there is a path from the root to some vertex 𝑢
 such that each of the given 𝑘
 vertices is either belongs to this path or has the distance 1
 to some vertex of this path.

 ### ideas
 1. u可以选为k个节点中，最深的那个点，那么其他的点，就可以判断，是否在这条path上，或者parent在这个path上
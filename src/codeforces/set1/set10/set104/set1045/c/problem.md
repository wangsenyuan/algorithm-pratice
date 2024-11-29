In an unspecified solar system, there are 𝑁
 planets. A space government company has recently hired space contractors to build 𝑀
 bidirectional Hyperspace™ highways, each connecting two different planets. The primary objective, which was to make sure that every planet can be reached from any other planet taking only Hyperspace™ highways, has been completely fulfilled. Unfortunately, lots of space contractors had friends and cousins in the Space Board of Directors of the company, so the company decided to do much more than just connecting all planets.

In order to make spending enormous amounts of space money for Hyperspace™ highways look neccessary, they decided to enforce a strict rule on the Hyperspace™ highway network: whenever there is a way to travel through some planets and return to the starting point without travelling through any planet twice, every pair of planets on the itinerary should be directly connected by a Hyperspace™ highway. In other words, the set of planets in every simple cycle induces a complete subgraph.

You are designing a Hyperspace™ navigational app, and the key technical problem you are facing is finding the minimal number of Hyperspace™ highways one needs to use to travel from planet 𝐴
 to planet 𝐵
. As this problem is too easy for Bubble Cup, here is a harder task: your program needs to do it for 𝑄
 pairs of planets.

Input
The first line contains three positive integers 𝑁
 (1≤𝑁≤100000
), 𝑀
 (1≤𝑀≤500000
) and 𝑄
 (1≤𝑄≤200000
), denoting the number of planets, the number of Hyperspace™ highways, and the number of queries, respectively.

Each of the following M lines contains a highway: highway 𝑖
 is given by two integers 𝑢𝑖
 and 𝑣𝑖
 (1≤𝑢𝑖<𝑣𝑖≤𝑁
), meaning the planets 𝑢𝑖
 and 𝑣𝑖
 are connected by a Hyperspace™ highway. It is guaranteed that the network of planets and Hyperspace™ highways forms a simple connected graph.

Each of the following 𝑄
 lines contains a query: query 𝑗
 is given by two integers 𝑎𝑗
 and 𝑏𝑗
 (1≤𝑎𝑗<𝑏𝑗≤𝑁)
, meaning we are interested in the minimal number of Hyperspace™ highways one needs to take to travel from planet 𝑎𝑗
 to planet 𝑏𝑗
.

### ideas
1. 那些在环上的节点，进入到离开，距离为2
2. 假设从u -> v, 那么必然经过了好几个环，然后进入某个环（把单个节点也作为环）
3. 那么就是求lca 了
4. 除了lca，还需要知道从哪个点进入的这个ssc
5. 不大对。 比如例子中节点3，它可以在两个component中
6. 但是有个想法就是，所有的节点，大体组成了一条line
7. 那么只要知道某个节点在这条line上的什么位置，然后就可以计算了
8. 然后就是怎么找到这条line的问题
9. 从1出发，找到最远的点，然后再从那个地方出发，找到一条最长的线
10. 这条线，就是轴线
11. 每个部分最多只有两个点再这个轴线上
12. 这个假设是错误的
13. 它们应该构成了一棵树，而不是构成一条线
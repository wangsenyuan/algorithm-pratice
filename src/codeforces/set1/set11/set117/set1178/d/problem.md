Every person likes prime numbers. Alice is a person, thus she also shares the love for them. Bob wanted to give her an affectionate gift but couldn't think of anything inventive. Hence, he will be giving her a graph. How original, Bob! Alice will surely be thrilled!

When building the graph, he needs four conditions to be satisfied:

It must be a simple undirected graph, i.e. without multiple (parallel) edges and self-loops.
The number of vertices must be exactly 𝑛
 — a number he selected. This number is not necessarily prime.
The total number of edges must be prime.
The degree (i.e. the number of edges connected to the vertex) of each vertex must be prime.
Below is an example for 𝑛=4
. The first graph (left one) is invalid as the degree of vertex 2
 (and 4
) equals to 1
, which is not prime. The second graph (middle one) is invalid as the total number of edges is 4
, which is not a prime number. The third graph (right one) is a valid answer for 𝑛=4
.

### ideas
1. 如果n不是质数，下一个质数 < 2 * n
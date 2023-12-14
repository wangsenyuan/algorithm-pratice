You are given an integer 𝑛
, as well as 𝑚
pairs of integers (𝑎𝑖,𝑏𝑖)
, where 1≤𝑎𝑖,𝑏𝑖≤𝑛
, 𝑎𝑖≠𝑏𝑖
.

You want to construct a sequence satisfying the following requirements:

1. All elements in the sequence are integers between 1
   and 𝑛
   .
2. There is exactly one element with value 1
3. in the sequence.
   For each 𝑖
   (1≤𝑖≤𝑚
   ), between any two elements (on different positions) in the sequence with value 𝑎𝑖
   , there is at least one element with value 𝑏𝑖
   .
4. The sequence constructed has the maximum length among all possible sequences satisfying the above properties.

Sometimes, it is possible that such a sequence can be arbitrarily long, in which case you should output "INFINITE".
Otherwise, you should output "FINITE" and the sequence itself. If there are multiple possible constructions that
yield the maximum length, output any.

### thoughts

### solution

Construct a graph with 𝑛
vertices and add a directed edge 𝑎→𝑏
if between every two 𝑎
there must be a 𝑏
.

Let 𝑣𝑎
be the number of occurrences of 𝑎
. The key observation is that if 𝑎→𝑏
, then 𝑣𝑎≤𝑣𝑏+1
.

Suppose 𝑎𝑘→𝑎𝑘−1→⋯𝑎1
is a directed path, where 𝑎1=1
. Then since 𝑣1=1
, we must have 𝑣𝑎𝑖≤𝑖
. In other words, 𝑣𝑠≤𝑑𝑠
. where 𝑑𝑠
is one plus the length of the shortest directed path from 𝑠
to 1
.

Therefore, the total array length does not exceed ∑𝑛𝑖=1𝑑𝑖
. We claim that we can achieve this.

It is easy to calculate the 𝑑𝑠
by a BFS. Let 𝑇𝑖
consists of vertices 𝑥
such that 𝑣𝑥=𝑠
. Let 𝑀
the largest value of 𝑑𝑖
among all 𝑖∈1,2⋯𝑛
. Consider

[𝑇𝑀],[𝑇𝑀−1][𝑇𝑀],[𝑇𝑀−2][𝑇𝑀−1][𝑇𝑀],⋯[𝑇1][𝑇2][𝑇3]⋯[𝑇𝑚]
where for each 𝑖
, vertices in various occurrences of 𝑇𝑖
must be arranged in the same order.

It is easy to check that this construction satisfies all the constraints and achieve the upper bound ∑𝑛𝑖=1𝑑𝑖
. Thus, this output is correct.

The sequence can be arbitrarily long if and only if there is some 𝑣
that does not have a path directed to 1
. To see this, let 𝑆
be the set of vertices that do not have path directed to 1
, then the following construction gives an arbitrarily long output that satisfy all constraints:

1[𝑆][𝑆][𝑆]⋯
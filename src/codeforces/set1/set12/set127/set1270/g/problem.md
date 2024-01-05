You are given 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
, such that for each 1≤𝑖≤𝑛
holds 𝑖−𝑛≤𝑎𝑖≤𝑖−1
.

Find some nonempty subset of these integers, whose sum is equal to 0
. It can be shown that such a subset exists under given constraints. If there are several possible subsets with
zero-sum, you can find any of them.

Input
Each test contains multiple test cases. The first line contains the number of test cases 𝑡
(1≤𝑡≤106
). The description of the test cases follows.

The first line of each test case contains a single integer 𝑛
(1≤𝑛≤106
).

The second line of each test case contains 𝑛
integers 𝑎1,𝑎2,…,𝑎𝑛
(𝑖−𝑛≤𝑎𝑖≤𝑖−1
).

It is guaranteed that the sum of 𝑛
over all test cases does not exceed 106
.

Output
For each test case, output two lines.

In the first line, output 𝑠
(1≤𝑠≤𝑛
) — the number of elements in your subset.

In the second line, output 𝑠
integers 𝑖1,𝑖2,…,𝑖𝑠
(1≤𝑖𝑘≤𝑛
). All integers have to be pairwise different, and 𝑎𝑖1+𝑎𝑖2+⋯+𝑎𝑖𝑠
has to be equal to 0
. If there are several possible subsets with zero-sum, you can find any of them.

Note that the condition 𝑖−𝑛≤𝑎𝑖≤𝑖−1
is equivalent to 1≤𝑖−𝑎𝑖≤𝑛
.

Let's build an oriented graph 𝐺
on 𝑛
nodes with the following principle: for each 𝑖
from 1
to 𝑛
, draw an oriented edge from vertex 𝑖
to vertex 𝑖−𝑎𝑖
. In this graph, there is an outgoing edge from every vertex, so it has an oriented cycle. Let vertices of this cycle
be — 𝑖1,𝑖2,…,𝑖𝑘
.

Then:

𝑖1−𝑎𝑖1=𝑖2
𝑖2−𝑎𝑖2=𝑖3
⋮
𝑖𝑛−𝑎𝑖𝑛=𝑖1
After adding all these equalities, we get 𝑎𝑖1+𝑎𝑖2+⋯+𝑎𝑖𝑘=0
We can find some oriented cycle in 𝑂(𝑛)
(just go by an available edge until you get to previously visited vertex).
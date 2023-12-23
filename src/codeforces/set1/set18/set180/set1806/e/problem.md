You are given a tree with 𝑛
weighted vertices labeled from 1
to 𝑛
rooted at vertex 1
. The parent of vertex 𝑖
is 𝑝𝑖
and the weight of vertex 𝑖
is 𝑎𝑖
. For convenience, define 𝑝1=0
.

For two vertices 𝑥
and 𝑦
of the same depth†
, define 𝑓(𝑥,𝑦)
as follows:

Initialize ans=0
.
While both 𝑥
and 𝑦
are not 0
:
ans←ans+𝑎𝑥⋅𝑎𝑦
;
𝑥←𝑝𝑥
;
𝑦←𝑝𝑦
.
𝑓(𝑥,𝑦)
is the value of ans
.
You will process 𝑞
queries. In the 𝑖
-th query, you are given two integers 𝑥𝑖
and 𝑦𝑖
and you need to calculate 𝑓(𝑥𝑖,𝑦𝑖)
.

†
The depth of vertex 𝑣
is the number of edges on the unique simple path from the root of the tree to vertex 𝑣
.
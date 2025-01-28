You are given two simple undirected graphs 𝐹
and 𝐺
with 𝑛
vertices. 𝐹
has 𝑚1
edges while 𝐺
has 𝑚2
edges. You may perform one of the following two types of operations any number of times:

Select two integers 𝑢
and 𝑣
(1≤𝑢,𝑣≤𝑛
) such that there is an edge between 𝑢
and 𝑣
in 𝐹
. Then, remove that edge from 𝐹
.
Select two integers 𝑢
and 𝑣
(1≤𝑢,𝑣≤𝑛
) such that there is no edge between 𝑢
and 𝑣
in 𝐹
. Then, add an edge between 𝑢
and 𝑣
in 𝐹
.
Determine the minimum number of operations required such that for all integers 𝑢
and 𝑣
(1≤𝑢,𝑣≤𝑛
), there is a path from 𝑢
to 𝑣
in 𝐹
if and only if there is a path from 𝑢
to 𝑣
in 𝐺
.

### ideas

1. 对于G中的一个comp，F中的也要和它一致，但却不用完全一样，比如G中有1-2-3， 那么 1-3-2也是ok的
2. 如果 u-v不在一个comp，必须删除（+1）
3. 如果 u-v没有边，还需要想办法连起来
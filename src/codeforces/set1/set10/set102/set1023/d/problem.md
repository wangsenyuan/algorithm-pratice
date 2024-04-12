Initially there was an array 𝑎
consisting of 𝑛
integers. Positions in it are numbered from 1
to 𝑛
.

Exactly 𝑞
queries were performed on the array. During the 𝑖
-th query some segment (𝑙𝑖,𝑟𝑖)
(1≤𝑙𝑖≤𝑟𝑖≤𝑛)
was selected and values of elements on positions from 𝑙𝑖
to 𝑟𝑖
inclusive got changed to 𝑖
. The order of the queries couldn't be changed and all 𝑞
queries were applied. It is also known that every position from 1
to 𝑛
got covered by at least one segment.

We could have offered you the problem about checking if some given array (consisting of 𝑛
integers with values from 1
to 𝑞
) can be obtained by the aforementioned queries. However, we decided that it will come too easy for you.

So the enhancement we introduced to it is the following. Some set of positions (possibly empty) in this array is
selected and values of elements on these positions are set to 0
.

Your task is to check if this array can be obtained by the aforementioned queries. Also if it can be obtained then
restore this array.

If there are multiple possible arrays then print any of them.

### ideas

1. 如果没有出现0，那么怎么处理呢？
2. 首先要找到q的区间，这个区间必须是连续的，然后把它删除掉，再找到q-1的区间，再删除掉
3. 如果出现0呢？忽略就可以了吗？
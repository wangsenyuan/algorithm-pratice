You are given an array of integers 𝑎1,𝑎2,…,𝑎𝑛
, as well as a binary string†
𝑠
consisting of 𝑛
characters.

Augustin is a big fan of data structures. Therefore, he asked you to implement a data structure that can answer 𝑞
queries. There are two types of queries:

"1 𝑙
𝑟
" (1≤𝑙≤𝑟≤𝑛
) — replace each character 𝑠𝑖
for 𝑙≤𝑖≤𝑟
with its opposite. That is, replace all 𝟶
with 𝟷
and all 𝟷
with 𝟶
.
"2 𝑔
" (𝑔∈{0,1}
) — calculate the value of the bitwise XOR of the numbers 𝑎𝑖
for all indices 𝑖
such that 𝑠𝑖=𝑔
. Note that the XOR
of an empty set of numbers is considered to be equal to 0
.
Please help Augustin to answer all the queries!

For example, if 𝑛=4
, 𝑎=[1,2,3,6]
, 𝑠=𝟷𝟶𝟶𝟷
, consider the following series of queries:

"2 0
" — we are interested in the indices 𝑖
for which 𝑠𝑖=𝟶
, since 𝑠=𝟷𝟶𝟶𝟷
, these are the indices 2
and 3
, so the answer to the query will be 𝑎2⊕𝑎3=2⊕3=1
.
"1 1
3
" — we need to replace the characters 𝑠1,𝑠2,𝑠3
with their opposites, so before the query 𝑠=𝟷𝟶𝟶𝟷
, and after the query: 𝑠=𝟶𝟷𝟷𝟷
.
"2 1
" — we are interested in the indices 𝑖
for which 𝑠𝑖=𝟷
, since 𝑠=𝟶𝟷𝟷𝟷
, these are the indices 2
, 3
, and 4
, so the answer to the query will be 𝑎2⊕𝑎3⊕𝑎4=2⊕3⊕6=7
.
"1 2
4
" — 𝑠=𝟶𝟷𝟷𝟷
→
𝑠=𝟶𝟶𝟶𝟶
.
"2 1
" — 𝑠=𝟶𝟶𝟶𝟶
, there are no indices with 𝑠𝑖=𝟷
, so since the XOR
of an empty set of numbers is considered to be equal to 0
, the answer to this query is 0
.

### thoughts

1. 操作1比较关键，
2. 考虑有这样一种数据结构，它表示连续相同0/1的段，那么操作1的结果是，对中间的段没有影响，影响的只是l,r之间的连接
3. 有可能产生新的段，也可能将两段连接在一起，
4. 这里维护两个队列，0/1的队列，当切换的时候，就是把对两个队列属于[l...r]中间的部分，交换位置，同时更新这两个队列的xor值
5. 但是要整体的移动一个区间，似乎也不是很高效，
6. 需要用到那个特殊的数据结构了～

### solution

Of course this problem has solutions that use data structures. For example, you can use a segment tree with range
updates to solve it in 𝑂((𝑛+𝑞)log𝑛)
time, or you can use a square root decomposition to solve it in 𝑂((𝑛+𝑞)𝑛√)
time.

However, of course, we do not expect participants in Div3 to have knowledge of these advanced data structures, so there
is a simpler solution.

We will store 2 variables: 𝑋0,𝑋1
, which represent the XOR of all numbers from group 0
and group 1
, respectively. When answering a query of type 2, we will simply output either 𝑋0
or 𝑋1
. Now we need to understand how to update 𝑋0
and 𝑋1
after receiving a query of type 1.

Let's first solve a simplified version: suppose that in type 1 queries, only a single character of the string 𝑠
is inverted, i.e., 𝑙=𝑟
in all type 1 queries.

Let's see how 𝑋0
and 𝑋1
change after this query. If 𝑠𝑖
was 0
and became 1
, then the number 𝑎𝑖
will be removed from group 0
. So, we need to invert "XOR 𝑎𝑖
" from 𝑋0
. Since XOR is its own inverse operation (𝑤⊕𝑤=0
), we can do this with 𝑋0=𝑋0⊕𝑎𝑖
. And in 𝑋1
, we need to add the number 𝑎𝑖
, since now 𝑠𝑖=1
. And we can do this with 𝑋1=𝑋1⊕𝑎𝑖
.

The same thing happens if 𝑠𝑖
was 1
.

This is the key observation: when we invert 𝑠𝑖
, 𝑋0
and 𝑋1
change in the same way, regardless of whether this inversion was from 0
to 1
or from 1
to 0
.

Therefore, to update 𝑋0
and 𝑋1
after a query of type 1 with parameters 𝑙,𝑟
, we need to do this: 𝑋0=𝑋0⊕(𝑎𝑙⊕…⊕𝑎𝑟)
, and the same for 𝑋1
.

To quickly find the XOR value on a subsegment of the array 𝑎
, we can use a prefix XOR array. If 𝑝𝑖=𝑎1⊕…𝑎𝑖
, then: 𝑎𝑙⊕…⊕𝑎𝑟=𝑝𝑟⊕𝑝𝑙−1
.
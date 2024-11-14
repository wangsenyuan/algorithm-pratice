Iris has a tree rooted at vertex 1
. Each vertex has a value of 𝟶
 or 𝟷
.

Let's consider a leaf of the tree (the vertex 1
 is never considered a leaf) and define its weight. Construct a string formed by the values of the vertices on the path starting at the root and ending in this leaf. Then the weight of the leaf is the difference between the number of occurrences of 𝟷𝟶
 and 𝟶𝟷
 substrings in it.

Take the following tree as an example. Green vertices have a value of 𝟷
 while white vertices have a value of 𝟶
.


Let's calculate the weight of the leaf 5
: the formed string is 𝟷𝟶𝟷𝟷𝟶
. The number of occurrences of substring 𝟷𝟶
 is 2
, the number of occurrences of substring 𝟶𝟷
 is 1
, so the difference is 2−1=1
.
Let's calculate the weight of the leaf 6
: the formed string is 𝟷𝟶𝟷
. The number of occurrences of substring 𝟷𝟶
 is 1
, the number of occurrences of substring 𝟶𝟷
 is 1
, so the difference is 1−1=0
.


### ideas
1. 考虑 0?0, 无论？是0，还是1，对结果没有影响，如果是1,-1然后+1
2. 考虑 0?1， 也没有影响
3. 所以，唯一有影响的是头和尾
4. 先考虑头不是？的情况，这个比较简单，iris看那些需要被换成1的，bob找那些换成0的就可以了
5. 000？这个换成0，对iris没有坏处，但是换成1，会扣分
6. 111？对于bob来说，也是一样的
7. 就看有几个，依次处理就可以了。如果切好有偶数个叶子节点要处理，那么这偶数个叶子节点会抵消掉
8. 如果多一个，那么iris得一分
9. 现在考虑s[0] = '?'的情况，
10. iris似乎必须先把s[0]定下来。
11. 其实一串的得分就看头尾
12. 比如 000， 101， 010 = 0
13. 001 011, ... -1
14. 100 110  1010, 1
15. 似乎有简单一些了
16. 头部确定的情况下，看尾部的数量（0/1确定的计算）？的看个数（iris占优）
17. 头部不确定的情况下，只要0/1节点的数量 > 0, iris都必须先确定头部节点(都应该放置1)
18. 如果全部都是？，iris还是要先放置1
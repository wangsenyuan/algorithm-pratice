Qingshan has a string 𝑠
which only contains 𝟶
and 𝟷
.

A string 𝑎
of length 𝑘
is good if and only if

𝑎𝑖≠𝑎𝑘−𝑖+1
for all 𝑖=1,2,…,𝑘
.
For Div. 2 contestants, note that this condition is different from the condition in problem B.

For example, 𝟷𝟶
, 𝟷𝟶𝟷𝟶
, 𝟷𝟷𝟷𝟶𝟶𝟶
are good, while 𝟷𝟷
, 𝟷𝟶𝟷
, 𝟶𝟶𝟷
, 𝟶𝟶𝟷𝟷𝟶𝟶
are not good.

Qingshan wants to make 𝑠
good. To do this, she can do the following operation at most 300
times (possibly, zero):

insert 𝟶𝟷
to any position of 𝑠
(getting a new 𝑠
).
Please tell Qingshan if it is possible to make 𝑠
good. If it is possible, print a sequence of operations that makes 𝑠
good.

### thoughts

1. 如果是奇数长度，肯定不行，因为中间的无法满足条件
2. 偶数位，感觉还是有戏的
3. 考虑 (i, j) (i = k + 1 - j) 它们是相同的
4. 这时候,如果是在它们中间或者外部添加元素，对它们没有影响
5. 000110，那只能在边界这里添加，如果 s[i] = 0,那么在j的后面添加后， 和s[i]匹配的正好是1
6. 如果s[i] = 1, 那就在i的前面添加，和j匹配的，正好是0
7. 但是如何保证在300次内是ok的？
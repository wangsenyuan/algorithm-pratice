You are given an array, consisting of 𝑛
integers. Initially, all elements are red.

You can apply the following operation to the array multiple times. During the 𝑖
-th operation, you select an element of the array; then:

if the element is red, it increases by 𝑖
and becomes blue;
if the element is blue, it decreases by 𝑖
and becomes red.
The operations are numbered from 1
, i. e. during the first operation some element is changed by 1
and so on.

You are asked 𝑞
queries of the following form:

given an integer 𝑘
, what can the largest minimum in the array be if you apply exactly 𝑘
operations to it?
Note that the operations don't affect the array between queries, all queries are asked on the initial array 𝑎
.

### thoughts

1. 进行正好k次操作，使得a种的最小值最大
2. 对于a[j]如果对它只进行了一次操作，那么可以让它净增加
3. 如果进行了2次操作，那么就时净减少
4. 是否可以二分呢？如果能够到达x+1， 是否一定能够到达x呢？
5. 对任何一个数来说，它一开始是red的
6. +i，然后变成蓝色，再操作一次后变成了 -j
7. 如果 k <= n, 那么最优的方案，是给所有的数都操作一次
8. let k -= n， 如果k是偶数，那么对于任何一个数，都可以进行 偶数次， 且剩余的偶数次，都要作用在blue的数上面
9. -i + j (且j > i)
10. 如果k是奇数，最好的处理是在第2次的时候，将其作用在第一个操作的数上面
11. 剩下的数，就变成了偶数次
12. 所以这里一个想法是，选择x个数，使得k-x 是偶数，然后判断，是否可以让这x数，变成比expect大，已经剩余的数比expect也要大
13. 
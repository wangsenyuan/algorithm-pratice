Let D(x) be the number of positive divisors of a positive integer x. For example, D(2)= 2 (2 is divisible by 1 and 2),
D(6)= 4 (6 is divisible by 1, 2, 3 and 6).

You are given an array a of n integers. You have to process two types of queries:

REPLACE l r — for every replace ai with D(ai);
SUM l r — calculate .
Print the answer for each SUM query.

### ideas

1. let f(x) = the number of positive divisors of x
2. 这个个数貌似是 sqrt(x) 的，
3. 所以直觉上，如果某个数(x >= 2)经过足够次数replace后，它就变成了2，且不再变化了
4. 1 （可能需要特殊处理）， 先放着
5. 对一个区间进行replace，发生了什么呢？
6. 貌似最多操作10次， 就变成了2
7. 1000000 -> 2000 -> 100 -> 10
8. 所以每个节点上维护一个0次操作时的计数 -> 1次操作时的计数... -> x次
9. it is 6
10. 直接用segment tree维护每段被更新了几次，这个方案是不行的
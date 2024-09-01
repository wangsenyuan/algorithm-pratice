You have to restore the wall. The wall consists of 𝑁
 pillars of bricks, the height of the 𝑖
-th pillar is initially equal to ℎ𝑖
, the height is measured in number of bricks. After the restoration all the 𝑁
 pillars should have equal heights.

You are allowed the following operations:

put a brick on top of one pillar, the cost of this operation is 𝐴
;
remove a brick from the top of one non-empty pillar, the cost of this operation is 𝑅
;
move a brick from the top of one non-empty pillar to the top of another pillar, the cost of this operation is 𝑀
.
You cannot create additional pillars or ignore some of pre-existing pillars even if their height becomes 0
.

What is the minimal total cost of restoration, in other words, what is the minimal total cost to make all the pillars of equal height?

### ideas
1. 假设最后的高度是x
2. 那么如果 x * n <= sum, 那么多余的这些，只能通过R的cost移除掉
3. 剩余的部分，可以通过move，或者remove + add 的操作来处理
4. 如果 R + A < M, 那么可以让 M = R + A， 且只考虑移动就可以了
5. 如果 x * n > sum, 那么缺少的部分，只能通过A的cost增加
6. 按照h排序后，删除肯定是后端的，增加肯定是前端的
7. 但是这里的问题是，给定x好计算，但是却没法一个个算过去。且不是x越大越好（此时需要add很多）也不是x越小越好（remove太多）
8. 貌似x是处于一个中间的值，且这个是一个一次函数，所以是线性的
9. 所以，可以使用3路搜索
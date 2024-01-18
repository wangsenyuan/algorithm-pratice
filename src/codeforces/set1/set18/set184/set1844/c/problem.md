You have discovered 𝑛
mysterious particles on a line with integer charges of 𝑐1,…,𝑐𝑛
. You have a device that allows you to perform the following operation:

Choose a particle and remove it from the line. The remaining particles will shift to fill in the gap that is created. If
there were particles with charges 𝑥
and 𝑦
directly to the left and right of the removed particle, they combine into a single particle of charge 𝑥+𝑦
.
For example, if the line of particles had charges of [−3,1,4,−1,5,−9]
, performing the operation on the 4
th particle will transform the line into [−3,1,9,−9]
.

### thoughts

1. 两端的负值，最好直接删除掉，否则会是结果更差
2. 假设在为止i处，进行操作，但是它左边得到的最好结果仍然小于0
3. 那么更好的操作是把它和左边同时舍弃掉
4. 对于两个连续的正数，是没有办法同时保留的
5. 对于为止i和j，如果它们中间有奇数个数，那么它们才有机会合并
6. a, b, c, d, e, f
7. a + c + e 似乎始终能够得到
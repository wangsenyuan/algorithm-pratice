You are given a sequence 𝑎1,𝑎2,…,𝑎𝑛
 consisting of 𝑛
 integers.

You may perform the following operation on this sequence: choose any element and either increase or decrease it by one.

Calculate the minimum possible difference between the maximum element and the minimum element in the sequence, if you can perform the aforementioned operation no more than 𝑘
 times.



 ### ideas
 1. binary search on expectation
 2. 最大值和最小值的差不超过expectation
 3. 有两种情况，一种是某些数不变，
 4. 还有一种是都变，小的数变大，大的数变小
 5. 第二种情况存在吗？
 6. 如果前后的数一样多，如果前面的数增加一个dx，后面的数，就可以减少dx，也就是可以在边界上获取到最大值，而不需要在中间
 7. 假设前面的数多，后面的数少，那么显然往多的那方移动，会减少操作次数，且不会影响到结果（如果原来ok的话）
 8. 所以，其实只有一种情况
 9. 边界始终在某个数那里（它可以是下边界，也可以是上边界）
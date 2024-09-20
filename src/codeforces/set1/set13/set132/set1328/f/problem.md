You are given the array 𝑎
 consisting of 𝑛
 elements and the integer 𝑘≤𝑛
.

You want to obtain at least 𝑘
 equal elements in the array 𝑎
. In one move, you can make one of the following two operations:

Take one of the minimum elements of the array and increase its value by one (more formally, if the minimum value of 𝑎
 is 𝑚𝑛
 then you choose such index 𝑖
 that 𝑎𝑖=𝑚𝑛
 and set 𝑎𝑖:=𝑎𝑖+1
);
take one of the maximum elements of the array and decrease its value by one (more formally, if the maximum value of 𝑎
 is 𝑚𝑥
 then you choose such index 𝑖
 that 𝑎𝑖=𝑚𝑥
 and set 𝑎𝑖:=𝑎𝑖−1
).
Your task is to calculate the minimum number of moves required to obtain at least 𝑘
 equal elements in the array.

### ideas
1. 假设最后的最多相等的数字是x, 那么x要么是最大值，要么是最小值， （除非一开始就有k个x）
2. 如果x不是最大值，也不是最小值，那么就不可能被选中，并-1/+1操作
3. 如果它是最小值，那么就需要把所有小于x的最小值都变成x，并保留大于x的不变
4. 这里忽略了一个严重的情况，就是假设最后的值是x，其实是允许存在x-1的
5. 假设最后是k个x，那么一开始有a个x，有b个从最小值增加上来，还有c个是从最大值减小而来
6. 那么 a + b + c = k
7. b个增加上来 = 所有前面的增加到x-1(这个是确定的)，其中b个增加上来（那么这个cost = b)
8. 对于减少下来的也是一样的， 就是全部的变成x+1, 其中c个减少下来, cost = c
9. 只要 b + c = k - a， 至于怎么变化的，没有关系
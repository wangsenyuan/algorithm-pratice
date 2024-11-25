You are given an array 𝑎
 consisting of 𝑛
 integers. In one move you can choose any 𝑎𝑖
 and divide it by 2
 rounding down (in other words, in one move you can set 𝑎𝑖:=⌊𝑎𝑖2⌋
).

You can perform such an operation any (possibly, zero) number of times with any 𝑎𝑖
.

Your task is to calculate the minimum possible number of operations required to obtain at least 𝑘
 equal numbers in the array.

Don't forget that it is possible to have 𝑎𝑖=0
 after some operations, thus the answer always exists.



### ideas
1. 构造一个图，根据 x/2 = y 建造出来
2. 然后假设最小的数是x, 然后一直去查找就可以了
3. 这个树的高度，不会超过30
4. 边的长度 = 操作的次数
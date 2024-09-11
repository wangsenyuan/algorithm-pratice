You have a bag of size 𝑛
. Also you have 𝑚
 boxes. The size of 𝑖
-th box is 𝑎𝑖
, where each 𝑎𝑖
 is an integer non-negative power of two.

You can divide boxes into two parts of equal size. Your goal is to fill the bag completely.

For example, if 𝑛=10
 and 𝑎=[1,1,32]
 then you have to divide the box of size 32
 into two parts of size 16
, and then divide the box of size 16
. So you can fill the bag with boxes of size 1
, 1
 and 8
.

Calculate the minimum number of divisions required to fill the bag of size 𝑛
.

### ideas
1. 如果sum(a) < n => -1
2. 否则的话，应该始终是有答案的
3. 因为，通过不断的二分，可以至少减少1
4. 然后看n的每一位，从低位开始；
5. 如果存在a[i] = n[i] => continue （过程中需要合并一些东西上来)
6. 否则的话，就需要从最近的分裂出来
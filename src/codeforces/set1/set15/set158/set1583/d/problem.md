It turns out that the meaning of life is a permutation 𝑝1,𝑝2,…,𝑝𝑛
 of the integers 1,2,…,𝑛
 (2≤𝑛≤100
). Omkar, having created all life, knows this permutation, and will allow you to figure it out using some queries.

A query consists of an array 𝑎1,𝑎2,…,𝑎𝑛
 of integers between 1
 and 𝑛
. 𝑎
 is not required to be a permutation. Omkar will first compute the pairwise sum of 𝑎
 and 𝑝
, meaning that he will compute an array 𝑠
 where 𝑠𝑗=𝑝𝑗+𝑎𝑗
 for all 𝑗=1,2,…,𝑛
. Then, he will find the smallest index 𝑘
 such that 𝑠𝑘
 occurs more than once in 𝑠
, and answer with 𝑘
. If there is no such index 𝑘
, then he will answer with 0
.

You can perform at most 2𝑛
 queries. Figure out the meaning of life 𝑝
.


In the sample, the hidden permutation 𝑝
 is [3,2,1,5,4]
. Three queries were made.

The first query is 𝑎=[4,4,2,3,2]
. This yields 𝑠=[3+4,2+4,1+2,5+3,4+2]=[7,6,3,8,6]
. 6
 is the only number that appears more than once, and it appears first at index 2
, making the answer to the query 2
.

The second query is 𝑎=[3,5,1,5,5]
. This yields 𝑠=[3+3,2+5,1+1,5+5,4+5]=[6,7,2,10,9]
. There are no numbers that appear more than once here, so the answer to the query is 0
.

The third query is 𝑎=[5,2,4,3,1]
. This yields 𝑠=[3+5,2+2,1+4,5+3,4+1]=[8,4,5,8,5]
. 5
 and 8
 both occur more than once here. 5
 first appears at index 3
, while 8
 first appears at index 1
, and 1<3
, making the answer to the query 1
.

Note that the sample is only meant to provide an example of how the interaction works; it is not guaranteed that the above queries represent a correct strategy with which to determine the answer.

### ideas
1. p是一个permuation [1...n]
2. 选定一个a,返回的是出现次数超过1次的，其中最小的下标
3. 如果a = [1, 1, 1, 1] => 所有的数加1，大家还是不不同，所以不行
4. 但是如果是 [1, 1, 2, 2] 也就是一半数+1， 一半数+2
5. 如果这个结果是0， =>前半段 < 后半段 [1, 2], [3, 4]
6. 然后交换 [2, 2, 1, 1] => 那么这样子肯定是有结果的， 假设这个数i，那么p[i] = 2（因为前面1.。。2，后面是3.。。4)所以，只有2能够和后面相同
7. 前进了一步
8. 如果 【1， 1， 2， 2】 = i, 那么能够确定的是 p[i] > n / 2 (比如p[i] = 3, 或者4)
9. [4, 1, 3, 2]， [2, 3, 4, 1]
10. [1, 1, 1, 1, 2] => 如果得到0 => p[n] = n else p[i] = p[n] + 1
11. [1, 1, (n)(i), 1, 3] => 0 => p[n] = n - 1吗？好像是可以的，else p[j] = p[n] + 2
12. 通过这种方式，可以知道 p[n]... n 的序列
13. 然后把剩下的再按照上面的方式进行
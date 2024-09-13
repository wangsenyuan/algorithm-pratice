Polycarp is a frequent user of the very popular messenger. He's chatting with his friends all the time. He has 𝑛
 friends, numbered from 1
 to 𝑛
.

Recall that a permutation of size 𝑛
 is an array of size 𝑛
 such that each integer from 1
 to 𝑛
 occurs exactly once in this array.

So his recent chat list can be represented with a permutation 𝑝
 of size 𝑛
. 𝑝1
 is the most recent friend Polycarp talked to, 𝑝2
 is the second most recent and so on.

Initially, Polycarp's recent chat list 𝑝
 looks like 1,2,…,𝑛
 (in other words, it is an identity permutation).

After that he receives 𝑚
 messages, the 𝑗
-th message comes from the friend 𝑎𝑗
. And that causes friend 𝑎𝑗
 to move to the first position in a permutation, shifting everyone between the first position and the current position of 𝑎𝑗
 by 1
. Note that if the friend 𝑎𝑗
 is in the first position already then nothing happens.

For example, let the recent chat list be 𝑝=[4,1,5,3,2]
:

if he gets messaged by friend 3
, then 𝑝
 becomes [3,4,1,5,2]
;
if he gets messaged by friend 4
, then 𝑝
 doesn't change [4,1,5,3,2]
;
if he gets messaged by friend 2
, then 𝑝
 becomes [2,4,1,5,3]
.
For each friend consider all position he has been at in the beginning and after receiving each message. Polycarp wants to know what were the minimum and the maximum positions.

### ideas
1. 模拟？
2. 需要根据a[i]把对应的friend的位置移动到第一位，同时，1...j (a[i]现在所处位置要增加1)
3. 感觉不大对。用双端队列?
4. 如果一个数字没有出现在a中，那么它最终的位置，看它后面的数字，有多少个出现过了（它们都会被移动到前面去）
5. 如果一个数字出现在了a中，那么看它两次出现的次数中间的，unique的那些数字的个数+1
6. 那怎么计算呢？用persistence tree 吗？
7. 假设按照节点出现的顺序进行排序，那么就是出现顺序在上次出现后的那些数字的和
8. 
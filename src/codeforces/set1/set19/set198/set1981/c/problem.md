Turtle was playing with a sequence 𝑎1,𝑎2,…,𝑎𝑛
 consisting of positive integers. Unfortunately, some of the integers went missing while playing.

Now the sequence becomes incomplete. There may exist an arbitrary number of indices 𝑖
 such that 𝑎𝑖
 becomes −1
. Let the new sequence be 𝑎′
.

Turtle is sad. But Turtle remembers that for every integer 𝑖
 from 1
 to 𝑛−1
, either 𝑎𝑖=⌊𝑎𝑖+12⌋
 or 𝑎𝑖+1=⌊𝑎𝑖2⌋
 holds for the original sequence 𝑎
.

Turtle wants you to help him complete the sequence. But sometimes Turtle makes mistakes, so you need to tell him if you can't complete the sequence.

Formally, you need to find another sequence 𝑏1,𝑏2,…,𝑏𝑛
 consisting of positive integers such that:

For every integer 𝑖
 from 1
 to 𝑛
, if 𝑎′𝑖≠−1
, then 𝑏𝑖=𝑎′𝑖
.
For every integer 𝑖
 from 1
 to 𝑛−1
, either 𝑏𝑖=⌊𝑏𝑖+12⌋
 or 𝑏𝑖+1=⌊𝑏𝑖2⌋
 holds.
For every integer 𝑖
 from 1
 to 𝑛
, 1≤𝑏𝑖≤109
.
If there is no sequence 𝑏1,𝑏2,…,𝑏𝑛
 that satisfies all of the conditions above, you need to report −1
.


### ideas
1. 分段考虑，假设有一段连续的-1， 它两端的数字分别是x,y
2. 如果中间只有一个-1， 那么可以很容易判断， 要们是2 * x, 要么是 x/2
3. x/2最多到1， 2 * x 最多到1 << 30
4. 如果 x = y, 那么中间有奇数长度的1，那可以用 *2, /2 的序列填充
5. 当然如果x最大的情况下，且不是偶数，那么是处理不了的
6. 如果 x < y
7. 上面的理解有点问题， x = [y/2], 或者 y = [x/2]
8. 所以，x是奇数时，y也可以是奇数
9. 所以，对于给定的起点x，必须知道在经过一段距离的-1后，它能覆盖的范围是什么
10. 比如1可以覆盖 [2, 3], 但是不能是1
11. 2 => [4, 5] union [1] （居然不是连续的）
12.  3 =》 (6, 7) union [1]
13.  4 => [8, 9] union [2]
14.  所以，从1开始，到达l后，如果是a[i] = [a[i+1] / 2] 的变化的话，可以得到[1 << l, 1 << (l+1) - 1]
15.  这个范围的数字
16.  或者 [1 << (l - 2), 1 << (l-1) - 1]范围的 (1 -> 2 -> 1 )少了2层
17.  也就是中间那层[1 << (l-1), 1 << l - 1]这部分是没有的
18.  这是从1开始变化的， 如果是从x开始呢？再乘以x?
19.  似乎是可以的
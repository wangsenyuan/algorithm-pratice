With a new body, our idol Aroma White (or should we call her Kaori Minamiya?) begins to uncover her lost past through the OS space.

The space can be considered a 2D plane, with an infinite number of data nodes, indexed from 0
, with their coordinates defined as follows:

The coordinates of the 0
-th node is (𝑥0,𝑦0)
For 𝑖>0
, the coordinates of 𝑖
-th node is (𝑎𝑥⋅𝑥𝑖−1+𝑏𝑥,𝑎𝑦⋅𝑦𝑖−1+𝑏𝑦)
Initially Aroma stands at the point (𝑥𝑠,𝑦𝑠)
. She can stay in OS space for at most 𝑡
 seconds, because after this time she has to warp back to the real world. She doesn't need to return to the entry point (𝑥𝑠,𝑦𝑠)
 to warp home.

While within the OS space, Aroma can do the following actions:

From the point (𝑥,𝑦)
, Aroma can move to one of the following points: (𝑥−1,𝑦)
, (𝑥+1,𝑦)
, (𝑥,𝑦−1)
 or (𝑥,𝑦+1)
. This action requires 1
 second.
If there is a data node at where Aroma is staying, she can collect it. We can assume this action costs 0
 seconds. Of course, each data node can be collected at most once.
Aroma wants to collect as many data as possible before warping back. Can you help her in calculating the maximum number of data nodes she could collect within 𝑡
 seconds?

### ideas
1. 这个题目没有告诉这个plane的尺寸，所以不大可能计算出它的大小，所以，直接计算出 (xs, ys)的值，也非常难
2. 比较特别的就是 ax, ay的取值
3. 在bx, by = 0时，考虑第i个数据节点， 它的下标是就是 (pow(ax, i), pow(ay, i))
4. 也就是说在i超过60时，要访问到第61个node，从（任何一个前面的datanode出发），都远远超过1e16
5. 所以，可以先算出前i个节点的位置，然后从(xs, ys)模拟即可
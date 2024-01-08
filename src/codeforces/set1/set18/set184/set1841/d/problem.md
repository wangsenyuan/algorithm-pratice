Two segments [𝑙1,𝑟1]
and [𝑙2,𝑟2]
intersect if there exists at least one 𝑥
such that 𝑙1≤𝑥≤𝑟1
and 𝑙2≤𝑥≤𝑟2
.

An array of segments [[𝑙1,𝑟1],[𝑙2,𝑟2],…,[𝑙𝑘,𝑟𝑘]]
is called beautiful if 𝑘
is even, and is possible to split the elements of this array into 𝑘2
pairs in such a way that:

every element of the array belongs to exactly one of the pairs;
segments in each pair intersect with each other;
segments in different pairs do not intersect.
For example, the array [[2,4],[9,12],[2,4],[7,7],[10,13],[6,8]]
is beautiful, since it is possible to form 3
pairs as follows:

the first element of the array (segment [2,4]
) and the third element of the array (segment [2,4]
);
the second element of the array (segment [9,12]
) and the fifth element of the array (segment [10,13]
);
the fourth element of the array (segment [7,7]
) and the sixth element of the array (segment [6,8]
).
As you can see, the segments in each pair intersect, and no segments from different pairs intersect.

You are given an array of 𝑛
segments [[𝑙1,𝑟1],[𝑙2,𝑟2],…,[𝑙𝑛,𝑟𝑛]]
. You have to remove the minimum possible number of elements from this array so that the resulting array is beautiful.

### thoughts

1. 对于两个区间，它们如果有重叠，在它们中间拉条线
2. 感觉可以用最优匹配
3. 但还有一个问题，就是如果 (a, b), (c, d) 那么 b和c要无关
4. 所以还是要排个序（考虑按照右端点排序）
5. 对于当前[l,r]，迭代前面可以和它重叠的分段
6. 那么找到在l前面的最优解
7. 要进行坐标压缩
You are developing a rating system for an online game. Every time a player participates in a match, the player's rating
changes depending on the results.

Initially, the player's rating is 0
. There are 𝑛
matches; after the 𝑖
-th match, the rating change is equal to 𝑎𝑖
(the rating increases by 𝑎𝑖
if 𝑎𝑖
is positive, or decreases by |𝑎𝑖|
if it's negative. There are no zeros in the sequence 𝑎
).

The system has an additional rule: for a fixed integer 𝑘
, if a player's rating has reached the value 𝑘
, it will never fall below it. Formally, if a player's rating at least 𝑘
, and a rating change would make it less than 𝑘
, then the rating will decrease to exactly 𝑘
.

Your task is to determine the value 𝑘
in such a way that the player's rating after all 𝑛
matches is the maximum possible (among all integer values of 𝑘
). If there are multiple possible answers, you can print any of them.

### thoughts

1. 如果知道k, 确实很容易计算最后的最大值是多少
2. 这个k显然不能太大，它大的话，可能永远都不能起作用，但也不能太小，太小可能得不到最优解
3. 这个k应该属于前面的值能够正常情况下能够达到的数值
4. 这个可以很方便的计算出来，
5. 假设是是x，那么需要知道第一次达到x的位置，假设是i，那么需要知道i之后，从x能够达到的最大值是什么？
6. 如果能知道x后面最大的数和最小的数，x + diff(max - min)
7. k选定后，它就是最后能够达到的最大值 （在a一直是正数的情况下，k可以无限大)
8. 如果不是，可以假设还能达到更大的k，k1, 那么可以在达到k1的时候，在设定为k，保证结果不比k1小
9. 那么就出现了两个情况，第一等于k， 和最后一次等于k, 把中间的结果给删除掉，剩下部分相加的sum就是k
10.

### solution

Let's fix some 𝑘
and look at the first and the last moment when the rating should fall below 𝑘
, but doesn't. After such moments, the rating is equal to 𝑘
. So we can "delete" all changes (array elements) between those moments. And the remaining changes (to the left from the
first moment and to the right from the last moment) can be considered without any additional constraints (the value of
𝑘
doesn't affect those changes).

Using this fact, we can see that the total rating is not greater than the sum of the whole array minus some continuous
segment. So the maximum possible final rating doesn't exceed the sum of the array minus the minimum sum segment (this
segment can be empty, if all elements are positive). In fact, there is always such 𝑘
that provides such rating. Let the minimum sum segment be [𝑙;𝑟]
, then 𝑘
is equal to the prefix sum from 1
-st to (𝑙−1)
-th positions of the array.

The only remaining question is: why the rating after 𝑟
-th match is equal to 𝑘
? It can't fall below 𝑘
(since the rating is at least 𝑘
already); and if it is greater than 𝑘
, then there is a positive suffix of the segment [𝑙;𝑟]
, so we can remove it, and the sum of the segment would decrease. Which means the segment [𝑙;𝑟]
is not minimum sum segment, which contradicts the previous statement. So the rating after 𝑟
-th match is equal to 𝑘
.
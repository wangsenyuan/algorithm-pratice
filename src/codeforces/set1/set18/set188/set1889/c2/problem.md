### solution

We consider a brute force solution first.

At the beginning, we calculate number of intervals that cover position 𝑖
for each 𝑖=1,2,…,𝑛
by prefix sum. Now we can know the number uncovered positions. Let it be 𝐴
. Then we need to calculate the number of new uncovered position after removing two intervals. Let it be 𝐵
. So the answer in the end is 𝐴+max𝐵
.

For calculating 𝐵
, let's enumerate two intervals 𝐼1,𝐼2
.

If they have no intersection, 𝐵
is equal to the number of positions that are covered exactly once in interval 𝐼1
and 𝐼2
;
If they have intersection. Let the intersection be 𝐽
(It is an interval). 𝐵
is equal to the number of positions that are covered exactly once in interval 𝐼1
and 𝐼2
plus the number of positions that are covered exactly twice in interval 𝐽
;
The algorithm can be done in 𝑂(𝑛+𝑚2)
time by prefix sum. For optimization, we should notice that:

In the "no intersection" case, we can just simply pick two best intervals.

In the "intersection" case, there are at most 𝑛
useful interval pairs. The proof and algorithm goes: for each position 𝑖
, if it is covered by exactly 2
intervals, then this interval pair is useful and may update the answer.

Great, but for code implementation, how do we find those interval pairs? For each interval [𝑙,𝑟]
, we consider it as two events (like difference and prefix sum): it appears at position 𝑙
, and disappears at position 𝑟+1
. That way, set or just array is able to handle.

Time complexity 𝑂(𝑛+𝑚)
.

𝑘
is bigger in this version. However, it is still small, which leads us to a DP approach.

Let 𝑑𝑝𝑖,𝑗
be the number of uncovered positions in [1,𝑖]
, and the last uncovered position is 𝑖
, and the number of deleted intervals is 𝑗
.

For transition, we need to get all the intervals that covers position 𝑖
. We've mentioned this in the editorial of the easy version.

Let's iterate the last uncovered position 𝑡
, and calculate the number of intervals that need to be deleted in this transition. Let it be 𝑑𝑡
and it is the number of interval [𝑙,𝑟]
such that 𝑡<𝑙≤𝑖≤𝑟
. Check this for example:

And the transition goes:

𝑑𝑝𝑖,𝑗←1+max𝑡=0𝑖−1𝑑𝑝𝑡,𝑗−𝑑𝑡
The time complexity is 𝑂(𝑛2𝑘)
now. We need to speed the transition up.

Note that 𝑑𝑡
is increasing while 𝑡
is decreasing, and 𝑑𝑡
is at most 𝑘
. It actually splits the transition interval [0,𝑖−1]
into at most 𝑘+1
intervals [𝑠0,𝑠1],[𝑠1+1,𝑠2],⋯,[𝑠𝑙−1+1,𝑠𝑙]
such that 0=𝑠0≤𝑠1<𝑠2<⋯<𝑠𝑙=𝑖−1
and for 𝑡
in the same interval, 𝑑𝑡
is the same.

So the sparse table can be used for transition.

Time complexity 𝑂(𝑛𝑘2)
. Memory complexity 𝑂(𝑘𝑛log𝑛)
.
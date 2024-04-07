Luba needs your help again! Luba has n TV sets. She knows that i-th TV set will be working from moment of time li till
moment ri, inclusive.

Luba wants to switch off one of TV sets in order to free the socket. Let's call some TV set redundant if after switching
it off the number of integer moments of time when at least one of TV sets is working won't decrease. Luba will be very
upset if she has to switch off a non-redundant TV set.

Help Luba by telling her the index of some redundant TV set. If there is no any, print -1.

Input
The first line contains one integer number n (1 ≤ n ≤ 2·105) — the number of TV sets.

Then n lines follow, each of them containing two integer numbers li, ri (0 ≤ li ≤ ri ≤ 109) denoting the working time of
i-th TV set.

Output
If there is no any redundant TV set, print -1. Otherwise print the index of any redundant TV set (TV sets are indexed
from 1 to n).

If there are multiple answers, print any of them.

### observation

1. 似乎有点简单呐
2. 对于区间[l..r] 如果这个区间被cover了，那么它就是可以被删除的
3. segment tree + lazy range update/query
4. 好像还有个简单点的办法，求区间 [l...r]中间的最小值，如果 >= 2, 那么这个区间就是可以被删除的
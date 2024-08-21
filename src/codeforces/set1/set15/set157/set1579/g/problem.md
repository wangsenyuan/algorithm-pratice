You are given n
 lengths of segments that need to be placed on an infinite axis with coordinates.

The first segment is placed on the axis so that one of its endpoints lies at the point with coordinate 0
. Let's call this endpoint the "start" of the first segment and let's call its "end" as that endpoint that is not the start.

The "start" of each following segment must coincide with the "end" of the previous one. Thus, if the length of the next segment is d
 and the "end" of the previous one has the coordinate x
, the segment can be placed either on the coordinates [x−d,x]
, and then the coordinate of its "end" is x−d
, or on the coordinates [x,x+d]
, in which case its "end" coordinate is x+d
.

The total coverage of the axis by these segments is defined as their overall union which is basically the set of points covered by at least one of the segments. It's easy to show that the coverage will also be a segment on the axis. Determine the minimal possible length of the coverage that can be obtained by placing all the segments on the axis without changing their order.

Input
The first line contains an integer t
 (1≤t≤1000
) — the number of test cases.

The next 2t
 lines contain descriptions of the test cases.

The first line of each test case description contains an integer n
 (1≤n≤104
) — the number of segments. The second line of the description contains n
 space-separated integers ai
 (1≤ai≤1000
) — lengths of the segments in the same order they should be placed on the axis.

### ideas
1. 这个区间，唯一的被它的左端点、右端点所确定{l..r} => [l, r] (如果 a[i]的长度小于r -l)（保持不变）
2. 感觉有个很简单的点，咋就想不到呢。。。
3. 假设最优的结果是 [l...r]
4. 那么从一开始排的时候，如果 a[i]往右放进去，没有超过r的情况下，就放到右边去？这个是一个合理的方案吗？
5. 假设suffix能够cover的最小的宽度是w
6. 处理i的时候，如果w比a[i]短，那么新的w就是a[i] (这是因为，可以让后缀掉转方向的情况，仍然得到w的宽度)
7. 知道w还不够，必须知道在(r, w)这样的对
8. 对于所有的(r, w)进行处理，以得到新的 (nr, nw)
9. 这样的r不会超过(n - i)个，所以整体的复杂性是 n * n的 
10. 试试看
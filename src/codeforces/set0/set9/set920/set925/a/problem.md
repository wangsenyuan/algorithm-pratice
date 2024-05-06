In the year of 30𝑋𝑋
participants of some world programming championship live in a single large hotel. The hotel has 𝑛
floors. Each floor has 𝑚
sections with a single corridor connecting all of them. The sections are enumerated from 1
to 𝑚
along the corridor, and all sections with equal numbers on different floors are located exactly one above the other.
Thus, the hotel can be represented as a rectangle of height 𝑛
and width 𝑚
. We can denote sections with pairs of integers (𝑖,𝑗)
, where 𝑖
is the floor, and 𝑗
is the section number on the floor.

The guests can walk along the corridor on each floor, use stairs and elevators. Each stairs or elevator occupies all
sections (1,𝑥)
, (2,𝑥)
, …
, (𝑛,𝑥)
for some 𝑥
between 1
and 𝑚
. All sections not occupied with stairs or elevators contain guest rooms. It takes one time unit to move between
neighboring sections on the same floor or to move one floor up or down using stairs. It takes one time unit to move up
to 𝑣
floors in any direction using an elevator. You can assume you don't have to wait for an elevator, and the time needed to
enter or exit an elevator is negligible.

You are to process 𝑞
queries. Each query is a question "what is the minimum time needed to go from a room in section (𝑥1,𝑦1)
to a room in section (𝑥2,𝑦2)
?"

Input
The first line contains five integers 𝑛,𝑚,𝑐𝑙,𝑐𝑒,𝑣
(2≤𝑛,𝑚≤108
, 0≤𝑐𝑙,𝑐𝑒≤105
, 1≤𝑐𝑙+𝑐𝑒≤𝑚−1
, 1≤𝑣≤𝑛−1
) — the number of floors and section on each floor, the number of stairs, the number of elevators and the maximum speed
of an elevator, respectively.

The second line contains 𝑐𝑙
integers 𝑙1,…,𝑙𝑐𝑙
in increasing order (1≤𝑙𝑖≤𝑚
), denoting the positions of the stairs. If 𝑐𝑙=0
, the second line is empty.

The third line contains 𝑐𝑒
integers 𝑒1,…,𝑒𝑐𝑒
in increasing order, denoting the elevators positions in the same format. It is guaranteed that all integers 𝑙𝑖
and 𝑒𝑖
are distinct.

The fourth line contains a single integer 𝑞
(1≤𝑞≤105
) — the number of queries.

The next 𝑞
lines describe queries. Each of these lines contains four integers 𝑥1,𝑦1,𝑥2,𝑦2
(1≤𝑥1,𝑥2≤𝑛
, 1≤𝑦1,𝑦2≤𝑚
) — the coordinates of starting and finishing sections for the query. It is guaranteed that the starting and finishing
sections are distinct. It is also guaranteed that these sections contain guest rooms, i. e. 𝑦1
and 𝑦2
are not among 𝑙𝑖
and 𝑒𝑖
.

### ideas

1. 好复杂的题目
2. 如果在同一行里，只需要计算距离，没有必要走楼梯，或者乘坐电梯
3. 如果在多行中，有一个观察是，如果能乘坐电梯，就乘坐电梯
4. 还有一个观察是，尽量使用同一个方向的电梯，比如 y1 < y2
5. 那么应该选择楼梯和电梯在y1到y2之间的，尽量减少回头的距离
6. 如果已经选择了一个电梯，就直接选用这个电梯即可
7. 楼梯只有在需要使用电梯，但是楼梯更近的情况下，使用才有作用
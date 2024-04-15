Vasya has got a robot which is situated on an infinite Cartesian plane, initially in the cell (0,0)
. Robot can perform the following four kinds of operations:

U — move from (𝑥,𝑦)
to (𝑥,𝑦+1)
;
D — move from (𝑥,𝑦)
to (𝑥,𝑦−1)
;
L — move from (𝑥,𝑦)
to (𝑥−1,𝑦)
;
R — move from (𝑥,𝑦)
to (𝑥+1,𝑦)
.
Vasya also has got a sequence of 𝑛
operations. Vasya wants to modify this sequence so after performing it the robot will end up in (𝑥,𝑦)
.

Vasya wants to change the sequence so the length of changed subsegment is minimum possible. This length can be
calculated as follows: 𝑚𝑎𝑥𝐼𝐷−𝑚𝑖𝑛𝐼𝐷+1
, where 𝑚𝑎𝑥𝐼𝐷
is the maximum index of a changed operation, and 𝑚𝑖𝑛𝐼𝐷
is the minimum index of a changed operation. For example, if Vasya changes RRRRRRR to RLRRLRL, then the operations with
indices 2
, 5
and 7
are changed, so the length of changed subsegment is 7−2+1=6
. Another example: if Vasya changes DDDD to DDRD, then the length of changed subsegment is 1
.

If there are no changes, then the length of changed subsegment is 0
. Changing an operation means replacing it with some operation (possibly the same); Vasya can't insert new operations
into the sequence or remove them.

Help Vasya! Tell him the minimum length of subsegment that he needs to change so that the robot will go from (0,0)
to (𝑥,𝑦)
, or tell him that it's impossible.

### ideas

1. 从 (0, 0) 到 (x, y) 需要横向移动x步，纵向移动y步
2. 如果 abs(x) + abs(y) > len(seq) => no
3. 否则肯定是有答案的
4. 固定r，然后找到l，使得改变(l...r)中间的操作后，能够到达目的地
5. 在l之前的dx, dy, 已经r后面的 dx, dy是不变的
6. 所以有 dx[l-1] + dx[r+1] + u == x,
7. dy[l-1] + dy[r+1] + v == y
8. 且 abs(u) + abs(v) = r - l + 1
9. 考虑r-l+1要满足什么条件，才能让6、7成立
10. 如果 r - l + 1 > abs(u) + abs(v) 也是可以的，这时候，只要简单的让前缀和原来的字符相同就可以了
11. u = x - dx[r] - dx[l]
12. v = y - dy[r] - dy[l]
13. 如果 u > 0 && v > 0
14. x - (...) + y - (...) = r - l - 1
15. 如果 u > 0 && v < 0
16. x - dx[r] - dx[l] - (y - dy[r] - dy[l]) = r - l - 1
17. x - dx[r] + dy[r] - (y - dx[l] - dy[l]) = r - l + 1
18. x - y - (dx[r] - dy[r]) - r - 1 = -(dx[l] - dy[l]) - l
19. 
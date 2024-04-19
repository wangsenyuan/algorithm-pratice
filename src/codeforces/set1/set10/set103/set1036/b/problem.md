Mikhail walks on a Cartesian plane. He starts at the point (0,0)
, and in one move he can go to any of eight adjacent points. For example, if Mikhail is currently at the point (0,0)
, he can go to any of the following points in one move:

(1,0)
;
(1,1)
;
(0,1)
;
(−1,1)
;
(−1,0)
;
(−1,−1)
;
(0,−1)
;
(1,−1)
.
If Mikhail goes from the point (𝑥1,𝑦1)
to the point (𝑥2,𝑦2)
in one move, and 𝑥1≠𝑥2
and 𝑦1≠𝑦2
, then such a move is called a diagonal move.

Mikhail has 𝑞
queries. For the 𝑖
-th query Mikhail's target is to go to the point (𝑛𝑖,𝑚𝑖)
from the point (0,0)
in exactly 𝑘𝑖
moves. Among all possible movements he want to choose one with the maximum number of diagonal moves. Your task is to
find the maximum number of diagonal moves or find that it is impossible to go from the point (0,0)
to the point (𝑛𝑖,𝑚𝑖)
in 𝑘𝑖
moves.

Note that Mikhail can visit any point any number of times (even the destination point!).

### ideas

1. when max(x, y) > k => -1 (至少需要max(x, y)步才能到达目的地)
2. 如果 k 和 max(x, y) 的 parity相同， 那么可以用 min(x, y)进行斜边运动，然后到达目的地，然后剩下的部分使用斜边运动即可
3. parity不同时，必须使用额外的两次直线运动，替换一次斜边运动
4. 分情况讨论 如果 min(x, y) > 0, 那么 -1， 然后k - 1， 然后其他的用斜边运动
5. 如果 min(x, y) =0, 但是 k - max(x, y) > 1 (k - max(x, y)) - 2
6. else 1
7. 如果 x, y = 0, 0， 但是 k = 1 => -1
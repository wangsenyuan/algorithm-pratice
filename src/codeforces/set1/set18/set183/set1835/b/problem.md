𝑛
people indexed with integers from 1
to 𝑛
came to take part in a lottery. Each received a ticket with an integer from 0
to 𝑚
.

In a lottery, one integer called target is drawn uniformly from 0
to 𝑚
. 𝑘
tickets (or less, if there are not enough participants) with the closest numbers to the target are declared the winners.
In case of a draw, a ticket belonging to the person with a smaller index is declared a winner.

Bytek decided to take part in the lottery. He knows the values on the tickets of all previous participants. He can pick
whatever value he wants on his ticket, but unfortunately, as he is the last one to receive it, he is indexed with an
integer 𝑛+1
.

Bytek wants to win the lottery. Thus, he wants to know what he should pick to maximize the chance of winning. He wants
to know the smallest integer in case there are many such integers. Your task is to find it and calculate his chance of
winning.

### thoughts

1. 题目都看不懂～缓缓再思考
2. 首先，bytek不能选择已经出现的数字，因为无论target是什么，他都不能win
3. 这些假设a已经排好序了
4. 对于一个target = x 属于[0...m]中间，判断bytek能够选择的数的范围
5. 对于x，如何找到范围[l...r]呢？这个应该可以使用尺取法的
6. bytek可以取的数，理论上是 [0...m]中的所有数，但只需要考虑其中不超过2 * n个数
7. 分别是每个数的+1和-?, 足够小，能够有

### solution

Let's assume that Bytek has selected a certain position 𝑐
. Let the closest occupied position to the left be 𝑑
, and the closest occupied position to the right be 𝑒
. Let's denote the position of the 𝑘
-th person to the left as 𝑎
and the 𝑘
-th person to the right as 𝑏
(on the picture 𝑘=3
).

Note that for Bytek to win, the target position should be closer to him than 𝑎
and closer to him than 𝑏
. So his winning range is in the interval (𝑎+𝑐2,𝑏+𝑐2)
. It will either have a length of ⌊𝑏−𝑎2−1⌋
or ⌊𝑏−𝑎2−2⌋
, which depends only on whether he chooses an even or odd position relative to the people on positions 𝑎
and 𝑏
.

So the solution to the task was to consider each pair of people standing next to each other and see what happens if
Bytek stands between them. To do this, we find the person 𝑘
positions to the left and 𝑘
positions to the right for Bytek and then check what the result will be if Bytek stands on the leftmost position inside
this interval and what if Bytek stands on the second position from the left inside this interval. The other positions in
this interval would give Bytek the same results but wouldn't be the leftmost.

In addition, we should look at what would happen if Bytek stood in a position where someone is already standing (this
may help if there is not enough space between consecutive people). There are also two more edge cases – from the left
and the right. One of these cases is to look at what would happen if Bytek stands one or two positions in front of the
𝑘
-th person from the left. This position would give Bytek the biggest winning range containing 0
. The other case is analogous from the right.

The final complexity is (𝑛)
or (𝑛log𝑛)
based on implementation.
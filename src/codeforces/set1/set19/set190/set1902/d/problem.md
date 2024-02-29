There is an infinite 2
-dimensional grid. Initially, a robot stands in the point (0,0)
. The robot can execute four commands:

U — move from point (𝑥,𝑦)
to (𝑥,𝑦+1)
;
D — move from point (𝑥,𝑦)
to (𝑥,𝑦−1)
;
L — move from point (𝑥,𝑦)
to (𝑥−1,𝑦)
;
R — move from point (𝑥,𝑦)
to (𝑥+1,𝑦)
.
You are given a sequence of commands 𝑠
of length 𝑛
. Your task is to answer 𝑞
independent queries: given four integers 𝑥
, 𝑦
, 𝑙
and 𝑟
; determine whether the robot visits the point (𝑥,𝑦)
, while executing a sequence 𝑠
, but the substring from 𝑙
to 𝑟
is reversed (i. e. the robot performs commands in order 𝑠1𝑠2𝑠3…𝑠𝑙−1𝑠𝑟𝑠𝑟−1𝑠𝑟−2…𝑠𝑙𝑠𝑟+1𝑠𝑟+2…𝑠𝑛
).

### solution

Let's divide the path of the robot into three parts:

points before the 𝑙
-th operation;
points from the 𝑙
-th to the 𝑟
-th operations;
points after the 𝑟
-th operation;
The first and the third parts are pretty simple to check because the reverse of the substring (𝑙,𝑟)
does not affect them. So we can precompute a dictionary 𝑚𝑝𝑥,𝑦
that stores all indices of operations in the initial sequence 𝑠
after which the robot stands at the point (𝑥,𝑦)
. Then point (𝑥,𝑦)
lies on the first part if 𝑚𝑝𝑥,𝑦
contains at least one index among indices from 0
to 𝑙−1
. Similarly, for the third part, but we have to check indices from 𝑟
to 𝑛
.

It remains to understand how the second part works. Reversing of the order of commands from 𝑙
to 𝑟
means that the robot first performs all commands from 1
to 𝑙−1
in the original string, then the command 𝑠𝑟
, then 𝑠𝑟−1
, and so on, until the command 𝑠𝑙
. So, it means that the points that the robot visits can be represented as follows: for an integer 𝑘
such that 𝑙≤𝑘≤𝑟
, perform all commands from 1
to 𝑟
, except for the commands from 𝑙
to 𝑘
.

Let 𝑝𝑜𝑠𝑖
be the point where the robot stands after the 𝑖
-th operation, and Δ𝑖,𝑗
be the total movement using operations from 𝑖
to 𝑗
(note that Δ𝑖,𝑗=𝑝𝑜𝑠𝑗−𝑝𝑜𝑠𝑖−1
). Then we have to check whether there is such 𝑖
that 𝑙≤𝑖≤𝑟
and 𝑝𝑜𝑠𝑙−1+Δ𝑖,𝑟=(𝑥,𝑦)
. Using the aforementioned equality for Δ
, we can rewrite this as 𝑝𝑜𝑠𝑙−1+𝑝𝑜𝑠𝑟−𝑝𝑜𝑠𝑖=(𝑥,𝑦)
. As a result, our task is to check whether there is such 𝑖
that 𝑙≤𝑖≤𝑟
and 𝑝𝑜𝑠𝑖=𝑝𝑜𝑠𝑙−1+𝑝𝑜𝑠𝑟−(𝑥,𝑦)
. And we already know how to check that, using the dictionary 𝑚𝑝
(you will need some sort of binary search or a function similar to lower_bound to find whether there is a moment from 𝑙
to 𝑟
when a point is visited).

If the point belongs to at least one of the parts of the path, then the answer is YES, otherwise the answer is NO.
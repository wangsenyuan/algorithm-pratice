salyg1n gave Alice a set 𝑆
of 𝑛
distinct integers 𝑠1,𝑠2,…,𝑠𝑛
(0≤𝑠𝑖≤109
). Alice decided to play a game with this set against Bob. The rules of the game are as follows:

Players take turns, with Alice going first.
In one move, Alice adds one number 𝑥
(0≤𝑥≤109
) to the set 𝑆
. The set 𝑆
must not contain the number 𝑥
at the time of the move.
In one move, Bob removes one number 𝑦
from the set 𝑆
. The set 𝑆
must contain the number 𝑦
at the time of the move. Additionally, the number 𝑦
must be strictly smaller than the last number added by Alice.
The game ends when Bob cannot make a move or after 2⋅𝑛+1
moves (in which case Alice's move will be the last one).
The result of the game is MEX†(𝑆)
(𝑆
at the end of the game).
Alice aims to maximize the result, while Bob aims to minimize it.
Let 𝑅
be the result when both players play optimally. In this problem, you play as Alice against the jury program playing as
Bob. Your task is to implement a strategy for Alice such that the result of the game is always at least 𝑅
.

†
MEX
of a set of integers 𝑐1,𝑐2,…,𝑐𝑘
is defined as the smallest non-negative integer 𝑥
which does not occur in the set 𝑐
. For example, MEX({0,1,2,4})
=
3
.

### thoughts

1. alice 的目标是让mex最大，而bob是为了最小
2. bob每次只能移除比x小的数，所以一定程度上alice可以控制bob的行为
3. 如果当前set的mex是R，alice可以保证结果至少是R。alice先摆放R，bob移除一个小于R的数，alice再放回被移除的数
4. 直到alice放置x = 0
5. R不可能比n大
6. 如果不存在0，那么alice可以在最后一步补上0，来终止游戏，但是这时bob能够移除的是
7. 而且这个必须是第一步，否则mex = 0. 这是因为，如果alice放置了任何一个x，如果没有比x小的数，游戏终止
8. 此时的mex = 0,或者存在某个数y < x, 那么bob移除这个最小的数即可，如果alice下次把y加入了
9. 因为没有更小的数（0也不存在），此时的mex = 0
10. 如果alice把0加入，此时的mex 最大也就是y
11. 如果数组那存在0，x = mex(set)
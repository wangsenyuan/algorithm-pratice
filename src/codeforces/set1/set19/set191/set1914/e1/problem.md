The easy and hard versions of this problem differ only in the constraints on the number of test cases and 𝑛
. In the easy version, the number of test cases does not exceed 103
, and 𝑛
 does not exceed 6
.

Recently, Alice and Bob were given marbles of 𝑛
 different colors by their parents. Alice has received 𝑎1
 marbles of color 1
, 𝑎2
 marbles of color 2
,..., 𝑎𝑛
 marbles of color 𝑛
. Bob has received 𝑏1
 marbles of color 1
, 𝑏2
 marbles of color 2
, ..., 𝑏𝑛
 marbles of color 𝑛
. All 𝑎𝑖
 and 𝑏𝑖
 are between 1
 and 109
.

After some discussion, Alice and Bob came up with the following game: players take turns, starting with Alice. On their turn, a player chooses a color 𝑖
 such that both players have at least one marble of that color. The player then discards one marble of color 𝑖
, and their opponent discards all marbles of color 𝑖
. The game ends when there is no color 𝑖
 such that both players have at least one marble of that color.

The score in the game is the difference between the number of remaining marbles that Alice has and the number of remaining marbles that Bob has at the end of the game. In other words, the score in the game is equal to (𝐴−𝐵)
, where 𝐴
 is the number of marbles Alice has and 𝐵
 is the number of marbles Bob has at the end of the game. Alice wants to maximize the score, while Bob wants to minimize it.

Calculate the score at the end of the game if both players play optimally.


### ideas
1. 考虑几种情况
2. i, j  a[i] > a[j], b[i] > b[j]
3. 那么对于alice来说，优先选择i是更好的，他可以获得a[i]-1分，且阻止对方获得b[i]-1分
4. 如果 a[i] > a[j], 但是 b[i] < b[j]呢
5. 如果alice选择i, 那么得分是 a[i] - 1, 但是bob得分是b[j] - 1
6. 反过来就是 a[j] - 1, b[i] - 1
7. 所以，从alice的角度来看如果 a[i] - b[j] > a[j] - b[i], 那么就是划算的
8. a[i] + b[i] > a[j] + b[j]
9. 对bob也是一样的
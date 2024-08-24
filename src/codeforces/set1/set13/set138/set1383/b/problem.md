Koa the Koala and her best friend want to play a game.

The game starts with an array 𝑎
 of length 𝑛
 consisting of non-negative integers. Koa and her best friend move in turns and each have initially a score equal to 0
. Koa starts.

Let's describe a move in the game:

During his move, a player chooses any element of the array and removes it from this array, xor-ing it with the current score of the player.
More formally: if the current score of the player is 𝑥
 and the chosen element is 𝑦
, his new score will be 𝑥⊕𝑦
. Here ⊕
 denotes bitwise XOR operation.

Note that after a move element 𝑦
 is removed from 𝑎
.

The game ends when the array is empty.
At the end of the game the winner is the player with the maximum score. If both players have the same score then it's a draw.

If both players play optimally find out whether Koa will win, lose or draw the game.

Input
Each test contains multiple test cases. The first line contains 𝑡
 (1≤𝑡≤104
) — the number of test cases. Description of the test cases follows.

The first line of each test case contains the integer 𝑛
 (1≤𝑛≤105
) — the length of 𝑎
.

The second line of each test case contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (0≤𝑎𝑖≤109
) — elements of 𝑎
.

It is guaranteed that the sum of 𝑛
 over all test cases does not exceed 105
.

### ideas
1. 又有点懵逼
2. 假设最后的结果是 x, y, 那么 x ^ y = xor(a[1], ... a[n]) = s
3. 看s的最高位bit，如果koa 能够先得到这个bit，是不是肯定能够获胜？
4. 也不是的，比如 [2, 2, 3]
5. koa 可以拿到3（以保证获得s[1]), 但是koala获得2，koa只能获得2
6. 所以，获胜方，应该是取得奇数个包含d位的数字
7. 如果cnt(num[d] = 1) = 1， 那么 koa获胜
8. cnt(..) / 2 如果是偶数 => koa 或者，否则 koala获胜
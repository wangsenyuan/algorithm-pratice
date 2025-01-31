Takahashi is playing a board game called Sugoroku.

On the board, there are
N+1 squares numbered
0 to
N. Takahashi starts at Square
0, and he has to stop exactly at Square
N to win the game.

The game uses a roulette with the
M numbers from
1 to
M. In each turn, Takahashi spins the roulette. If the number
x comes up when he is at Square
s, he moves to Square
s+x. If this makes him go beyond Square
N, he loses the game.

Additionally, some of the squares are Game Over Squares. He also loses the game if he stops at one of those squares. You
are given a string
S of length
N+1, representing which squares are Game Over Squares. For each
i
(0≤i≤N), Square
i is a Game Over Square if
S[i]=1 and not if
S[i]=0.

Find the sequence of numbers coming up in the roulette in which Takahashi can win the game in the fewest number of turns
possible. If there are multiple such sequences, find the lexicographically smallest such sequence. If Takahashi cannot
win the game, print
−1.

### ideas

1. dp[i] 表示i是否是一个win position
2. dp[n] = true
3. dp[i] = 在i+1...i+m的范围内存在win position,且s[i] = '0'
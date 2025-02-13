Suppose you play a game where the game field looks like a strip of 1×109
 square cells, numbered from 1
 to 109
.

You have 𝑛
 snakes (numbered from 1
 to 𝑛
) you need to place into some cells. Initially, each snake occupies exactly one cell, and you can't place more than one snake into one cell. After that, the game starts.

The game lasts for 𝑞
 seconds. There are two types of events that may happen each second:

snake 𝑠𝑖
 enlarges: if snake 𝑠𝑖
 occupied cells [𝑙,𝑟]
, it enlarges to a segment [𝑙,𝑟+1]
;
snake 𝑠𝑖
 shrinks: if snake 𝑠𝑖
 occupied cells [𝑙,𝑟]
, it shrinks to a segment [𝑙+1,𝑟]
.
Each second, exactly one of the events happens.

If at any moment of time, any snake runs into some obstacle (either another snake or the end of the strip), you lose. Otherwise, you win with the score equal to the maximum cell occupied by any snake so far.

What is the minimum possible score you can achieve?

### ideas
1. 在能够完成所有events的情况下，让所有蛇占据的格子最少？
2. 最大的格子，不是格子的数量
3. 最小化占据的最大的格子的编码
4. 有想法了
5. 
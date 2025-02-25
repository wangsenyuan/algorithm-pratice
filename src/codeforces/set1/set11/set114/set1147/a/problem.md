Alice and Bob are playing a game on a line with 𝑛
 cells. There are 𝑛
 cells labeled from 1
 through 𝑛
. For each 𝑖
 from 1
 to 𝑛−1
, cells 𝑖
 and 𝑖+1
 are adjacent.

Alice initially has a token on some cell on the line, and Bob tries to guess where it is.

Bob guesses a sequence of line cell numbers 𝑥1,𝑥2,…,𝑥𝑘
 in order. In the 𝑖
-th question, Bob asks Alice if her token is currently on cell 𝑥𝑖
. That is, Alice can answer either "YES" or "NO" to each Bob's question.

At most one time in this process, before or after answering a question, Alice is allowed to move her token from her current cell to some adjacent cell. Alice acted in such a way that she was able to answer "NO" to all of Bob's questions.

Note that Alice can even move her token before answering the first question or after answering the last question. Alice can also choose to not move at all.

You are given 𝑛
 and Bob's questions 𝑥1,…,𝑥𝑘
. You would like to count the number of scenarios that let Alice answer "NO" to all of Bob's questions.

Let (𝑎,𝑏)
 denote a scenario where Alice starts at cell 𝑎
 and ends at cell 𝑏
. Two scenarios (𝑎𝑖,𝑏𝑖)
 and (𝑎𝑗,𝑏𝑗)
 are different if 𝑎𝑖≠𝑎𝑗
 or 𝑏𝑖≠𝑏𝑗
.

### ideas
1. 就是检查 (i, i + 1) (i + 1, i)是否valid
2. 对于(i, i+1), 如果i出现了，i+1没有出现，那么就可以在i出现前，换成i+1即可
3. 如果(i, i+1) 如果i没有出现， 那么可以一直使用i，知道i+1最后一次出现后，换成i+1
4. 如果 i没有出现，或者i+1的最后一次出现，是在i的第一次出现前，那么(i, i+1)就是有效的
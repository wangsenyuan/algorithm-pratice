Two beavers, Timur and Marsel, play the following game.

There are n logs, each of exactly m meters in length. The beavers move in turns. For each move a beaver chooses a log and gnaws it into some number (more than one) of equal parts, the length of each one is expressed by an integer and is no less than k meters. Each resulting part is also a log which can be gnawed in future by any beaver. The beaver that can't make a move loses. Thus, the other beaver wins.

Timur makes the first move. The players play in the optimal way. Determine the winner.

Input
The first line contains three integers n, m, k (1 ≤ n, m, k ≤ 109).

Output
Print "Timur", if Timur wins, or "Marsel", if Marsel wins. You should print everything without the quotes.

### ideas. 
1. 考虑n 是偶数时，似乎 Marsel有必胜的策略，就是跟随Timur的操作即可
2. 当n=1时，假设f(i) = true 表示当log的长度是i时，当前的player是否能够获胜
3. f(i) if i <= k => false, 
4. else a * b = i, 且 a >= k, 且 f(a) = false, 或者 f(a) = true, 但是 b是奇数 
There is a country with 𝑛
citizens. The 𝑖
-th of them initially has 𝑎𝑖
money. The government strictly controls the wealth of its citizens. Whenever a citizen makes a purchase or earns some
money, they must send a receipt to the social services mentioning the amount of money they currently have.

Sometimes the government makes payouts to the poor: all citizens who have strictly less money than 𝑥
are paid accordingly so that after the payout they have exactly 𝑥
money. In this case the citizens don't send a receipt.

You know the initial wealth of every citizen and the log of all events: receipts and payouts. Restore the amount of
money each citizen has after all events.

Input
The first line contains a single integer 𝑛
(1≤𝑛≤2⋅105
) — the numer of citizens.

The next line contains 𝑛
integers 𝑎1
, 𝑎2
, ..., 𝑎𝑛
(0≤𝑎𝑖≤109
) — the initial balances of citizens.

The next line contains a single integer 𝑞
(1≤𝑞≤2⋅105
) — the number of events.

Each of the next 𝑞
lines contains a single event. The events are given in chronological order.

Each event is described as either 1 p x (1≤𝑝≤𝑛
, 0≤𝑥≤109
), or 2 x (0≤𝑥≤109
). In the first case we have a receipt that the balance of the 𝑝
-th person becomes equal to 𝑥
. In the second case we have a payoff with parameter 𝑥
.

### thoughts

1. 先考虑暴力处理，依次进行操作
2. 在遇到操作2时，在这之前的操作1中，所有设置p[i] <= x 的，都失效了，所以一个想法是把操作1记录下来
3. 在遇到操作2，或者结尾时，进行处理
4. 但即使这样，操作2本身的操作仍然有可能很多
5. 有一个观察，就是假设之前有个操作2，设置位x，后面设置为更大的值时，前面那个操作2也是无效的
6. 但是如果变小了，似乎还是有效的
7. 倒过来处理，遇到操作1，进行处理，并把它移除出去
8. 并记录到目前为止的操作2的最大值，
9. 操作1 p = max(2, x)
10. 
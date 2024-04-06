Polycarp is in really serious trouble — his house is on fire! It's time to save the most valuable items. Polycarp
estimated that it would take ti seconds to save i-th item. In addition, for each item, he estimated the value of di —
the moment after which the item i will be completely burned and will no longer be valuable for him at all. In
particular, if ti ≥ di, then i-th item cannot be saved.

Given the values pi for each of the items, find a set of items that Polycarp can save such that the total value of this
items is maximum possible. Polycarp saves the items one after another. For example, if he takes item a first, and then
item b, then the item a will be saved in ta seconds, and the item b — in ta + tb seconds after fire started.

Input
The first line contains a single integer n (1 ≤ n ≤ 100) — the number of items in Polycarp's house.

Each of the following n lines contains three integers ti, di, pi (1 ≤ ti ≤ 20, 1 ≤ di ≤ 2 000, 1 ≤ pi ≤ 20) — the time
needed to save the item i, the time after which the item i will burn completely and the value of item i.

Output
In the first line print the maximum possible total value of the set of saved items. In the second line print one integer
m — the number of items in the desired set. In the third line print m distinct integers — numbers of the saved items in
the order Polycarp saves them. Items are 1-indexed in the same order in which they appear in the input. If there are
several answers, print any of them.

### ideas

1. 假设有两个item， a/b, 拯救有没有顺序呢？
2. 如果d[a] < d[b] 感觉先save a更合理，因为救了a以后，还有机会save b
3. dp[x] 表示在时刻x时，前能获得的最大值
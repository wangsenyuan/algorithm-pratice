On his trip to Luxor and Aswan, Sagheer went to a Nubian market to buy some souvenirs for his friends and relatives. The market has some strange rules. It contains n different items numbered from 1 to n. The i-th item has base cost ai Egyptian pounds. If Sagheer buys k items with indices x1, x2, ..., xk, then the cost of item xj is axj + xj·k for 1 ≤ j ≤ k. In other words, the cost of an item is equal to its base cost in addition to its index multiplied by the factor k.

Sagheer wants to buy as many souvenirs as possible without paying more than S Egyptian pounds. Note that he cannot buy a souvenir more than once. If there are many ways to maximize the number of souvenirs, he will choose the way that will minimize the total cost. Can you help him with this task?

### ideas
1. 假设f(k)表示购买k个的产品的最小花费
2. 其中i的贡献是 a[i] + i * k, 那么显然是选择前k个 a[i] + i * k 最小的值
3. 所以binary search + sort
A batch of Christmas trees has arrived at the largest store in Berland. 𝑛
 customers have already come to the store, wanting to buy them.

Before the sales begin, the store needs to determine the price for one tree (the price is the same for all customers). To do this, the store has some information about each customer.

For the 𝑖
-th customer, two integers 𝑎𝑖
 and 𝑏𝑖
 are known, which define their behavior:

if the price of the product is at most 𝑎𝑖
, the customer will buy a tree and leave a positive review;
otherwise, if the price of the product is at most 𝑏𝑖
, the customer will buy a tree but leave a negative review;
otherwise, the customer will not buy a tree at all.
Your task is to calculate the maximum possible earnings for the store, given that it can receive no more than 𝑘
 negative reviews.


 ### ideas
 1. 是不是价格越高越好？似乎不是的，价格越高的时候，不购买的人就多了
 2. 假设价格为x时，那么所有 b[i] > x的人，就放弃了，所有a[i] < x 的，会给出差评
 3. 所以按照 b[i]升序处理
 4. 取x = b[i], 所有 b[?] < x 的，都不购买，且不评价， 那么收入 = x * (n - i)
 5. 此时，要计算，有多少个人的 a[i] < x, 如果这样的人，不超过k个， 那么 x * (n - i)是一个有效的答案
 6. 
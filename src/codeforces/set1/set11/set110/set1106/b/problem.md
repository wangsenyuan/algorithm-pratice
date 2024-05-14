Lunar New Year is approaching, and Bob is planning to go for a famous restaurant — "Alice's".

The restaurant "Alice's" serves 𝑛
 kinds of food. The cost for the 𝑖
-th kind is always 𝑐𝑖
. Initially, the restaurant has enough ingredients for serving exactly 𝑎𝑖
 dishes of the 𝑖
-th kind. In the New Year's Eve, 𝑚
 customers will visit Alice's one after another and the 𝑗
-th customer will order 𝑑𝑗
 dishes of the 𝑡𝑗
-th kind of food. The (𝑖+1)
-st customer will only come after the 𝑖
-th customer is completely served.

Suppose there are 𝑟𝑖
 dishes of the 𝑖
-th kind remaining (initially 𝑟𝑖=𝑎𝑖
). When a customer orders 1
 dish of the 𝑖
-th kind, the following principles will be processed.

If 𝑟𝑖>0
, the customer will be served exactly 1
 dish of the 𝑖
-th kind. The cost for the dish is 𝑐𝑖
. Meanwhile, 𝑟𝑖
 will be reduced by 1
.
Otherwise, the customer will be served 1
 dish of the cheapest available kind of food if there are any. If there are multiple cheapest kinds of food, the one with the smallest index among the cheapest will be served. The cost will be the cost for the dish served and the remain for the corresponding dish will be reduced by 1
.
If there are no more dishes at all, the customer will leave angrily. Therefore, no matter how many dishes are served previously, the cost for the customer is 0
.
If the customer doesn't leave after the 𝑑𝑗
 dishes are served, the cost for the customer will be the sum of the cost for these 𝑑𝑗
 dishes.

Please determine the total cost for each of the 𝑚
 customers.


 [problem](https://codeforces.com/problemset/problem/1106/B)


 ### ideas 
 1. 先理解一下这个题目
 2. 就是customer是依次到来的
 3. 如果能够满足customer j的要求，就会更新对应的库存，否则就跳过他
 4. 如果d[j] > 全部的food，那么就跳过它
 5. 否则就更新
 6. 剩下的就是模拟这个过程
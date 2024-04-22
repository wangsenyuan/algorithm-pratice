JATC loves Banh-mi (a Vietnamese food). His affection for Banh-mi is so much that he always has it for breakfast. This
morning, as usual, he buys a Banh-mi and decides to enjoy it in a special way.

First, he splits the Banh-mi into 𝑛
parts, places them on a row and numbers them from 1
through 𝑛
. For each part 𝑖
, he defines the deliciousness of the part as 𝑥𝑖∈{0,1}
. JATC's going to eat those parts one by one. At each step, he chooses arbitrary remaining part and eats it. Suppose
that part is the 𝑖
-th part then his enjoyment of the Banh-mi will increase by 𝑥𝑖
and the deliciousness of all the remaining parts will also increase by 𝑥𝑖
. The initial enjoyment of JATC is equal to 0
.

For example, suppose the deliciousness of 3
parts are [0,1,0]
. If JATC eats the second part then his enjoyment will become 1
and the deliciousness of remaining parts will become [1,_,1]
. Next, if he eats the first part then his enjoyment will become 2
and the remaining parts will become [_,_,2]
. After eating the last part, JATC's enjoyment will become 4
.

However, JATC doesn't want to eat all the parts but to save some for later. He gives you 𝑞
queries, each of them consisting of two integers 𝑙𝑖
and 𝑟𝑖
. For each query, you have to let him know what is the maximum enjoyment he can get if he eats all the parts with
indices in the range [𝑙𝑖,𝑟𝑖]
in some order.

All the queries are independent of each other. Since the answer to the query could be very large, print it modulo 109+7
.

### ideas

1. 对于一整段，通过先使用1，再使用0可以获得最优的结果
2. 假设有x个1，y个0 m = x + y
3. 第一个1贡献 = pow(2, m - 1)
4. 第x个1的贡献 = pow(2, m - x)
5. 如果全是1 => 2 ** pow(m) - 1
6. 
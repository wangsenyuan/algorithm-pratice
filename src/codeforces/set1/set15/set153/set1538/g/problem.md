Polycarp has 𝑥
 of red and 𝑦
 of blue candies. Using them, he wants to make gift sets. Each gift set contains either 𝑎
 red candies and 𝑏
 blue candies, or 𝑎
 blue candies and 𝑏
 red candies. Any candy can belong to at most one gift set.

Help Polycarp to find the largest number of gift sets he can create.

For example, if 𝑥=10
, 𝑦=12
, 𝑎=5
, and 𝑏=2
, then Polycarp can make three gift sets:

In the first set there will be 5
 red candies and 2
 blue candies;
In the second set there will be 5
 blue candies and 2
 red candies;
In the third set will be 5
 blue candies and 2
 red candies.
Note that in this example there is one red candy that Polycarp does not use in any gift set.



### ideas
1. 假设最后的总数是sum, 且中类型1， a red + b blue 是m个
2. 类型2, a blue + b red 是n个
3. m + n = sum => 最大化sum
4. a * m + b * n <= x
5. b * m + a * n <= y
6. (a + b) * (m + n) <= x + y 这个是成立的
7. 得到了一个 m + n 的上限 (x + y) / (a + b)
8. 另外这两个是一次方程，所以在给定m和n时，所以它们的和也是线性的（增长然后减少)
There are 𝑛
benches near the Main Walkway in Summer Infomatics School. These benches are numbered by integers from 1
to 𝑛
in order they follow. Also there are 𝑚
cookie sellers near the Walkway. The 𝑖
-th (1≤𝑖≤𝑚
) cookie sellers is located near the 𝑠𝑖
-th bench.

Petya is standing in the beginning of the Walkway. He will pass near all benches starting from the 1
-st bench and ending with the 𝑛
-th bench. Petya passes the distance between two consecutive benches in 1
minute. He has a knapsack with an infinite amount of cookies. Petya is going to eat cookies from his knapsack and buy
them from cookie sellers during the walk.

Petya eats cookies only near the benches according to the following rule: he will eat the cookie near the 𝑖
-th (1≤𝑖≤𝑛
) bench if and only if at least one of the following conditions holds:

1. There is a cookie seller near the 𝑖
   -th bench. Then Petya will buy a cookie from cookie seller and eat it immediately.
2. Petya has not yet eaten a cookie. Then Petya will take a cookie from his knapsack and eat it immediately.
3. At least 𝑑 minutes passed since Petya ate the previous cookie. In other words, Petya has not eaten a cookie near the
   benches 𝑖−1,𝑖−2,…,max(𝑖−𝑑+1,1)
   . Then Petya will take a cookie from his knapsack and eat it immediately.
   You may assume that Petya eats cookies instantly. Petya will not eat two or more cookies near the same bench.

You want to minimize the number of cookies Petya will eat during his walk. In order to do this, you will ask the
administration of the Summer Informatics School to remove exactly one cookie seller from the Walkway before Petya starts
his walk.

Please determine the minimum possible number of cookies Petya can eat after removing exactly one cookie seller. Also
determine the number of cookie sellers, such that if you remove one of them, Petya will eat the minimum possible number
of cookies.

### thoughts

1. 考虑在没有取消的情况下，petya怎么能够吃到最多的糖果？
2. petya应该贪心的吃，能吃的时候就吃,(这个比较容易证明)
3. 所以从后往前(sellers), f(i) = petya在i处吃糖的最有结果的话
4. f(i) = f(j) + dist(i, j) / (d + 1) where dist(i, j) >= d + 1
5. 另g(i) = 在最后一个是在i处吃糖的结果（在某些位置是吃不到糖的）
6. 如果把i处的取消了，那么结果 = g(k)  where k 是i的前导 + f(r) + dist(r, k) / (d + 1)
7. 且必须满足 dist(k, i) % (d + 1) != 0 否则，取消i没有效果
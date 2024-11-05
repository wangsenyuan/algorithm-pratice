There are 𝑛
 candies in a candy box. The type of the 𝑖
-th candy is 𝑎𝑖
 (1≤𝑎𝑖≤𝑛
).

You have to prepare a gift using some of these candies with the following restriction: the numbers of candies of each type presented in a gift should be all distinct (i. e. for example, a gift having two candies of type 1
 and two candies of type 2
 is bad).

It is possible that multiple types of candies are completely absent from the gift. It is also possible that not all candies of some types will be taken to a gift.

You really like some of the candies and don't want to include them into the gift, but you want to eat them yourself instead. For each candy, a number 𝑓𝑖
 is given, which is equal to 0
 if you really want to keep 𝑖
-th candy for yourself, or 1
 if you don't mind including it into your gift. It is possible that two candies of the same type have different values of 𝑓𝑖
.

You want your gift to be as large as possible, but you don't want to include too many of the candies you want to eat into the gift. So, you want to calculate the maximum possible number of candies that can be included into a gift, and among all ways to choose maximum number of candies, you want to maximize the number of candies having 𝑓𝑖=1
 in your gift.

You have to answer 𝑞
 independent queries.


 ### ideas
 1. 不考虑是否要保留的限制，如何计算最大的礼物的数量
 2. freq[x] 表示x（类型）的数量, 那么可以从最大的freq[?]开始，每次减少1（或者 freq[?])
 3. 这样加起来后，就是sum
 4. 那么在这个基础上，f[i] = 1 的个数是多少呢？
 5. 同时可以计算一个keep[x] = freq[x] - use[x] 表示可以保留类型x的数量
 6. 那么在这个保留中，尽量的使用f[i] = 0 的
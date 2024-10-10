This is the hard version of this problem. The only difference is the constraint on 𝑘
 — the number of gifts in the offer. In this version: 2≤𝑘≤𝑛
.

Vasya came to the store to buy goods for his friends for the New Year. It turned out that he was very lucky — today the offer "𝑘
 of goods for the price of one" is held in store.

Using this offer, Vasya can buy exactly 𝑘
 of any goods, paying only for the most expensive of them. Vasya decided to take this opportunity and buy as many goods as possible for his friends with the money he has.

More formally, for each good, its price is determined by 𝑎𝑖
 — the number of coins it costs. Initially, Vasya has 𝑝
 coins. He wants to buy the maximum number of goods. Vasya can perform one of the following operations as many times as necessary:

Vasya can buy one good with the index 𝑖
 if he currently has enough coins (i.e 𝑝≥𝑎𝑖
). After buying this good, the number of Vasya's coins will decrease by 𝑎𝑖
, (i.e it becomes 𝑝:=𝑝−𝑎𝑖
).
Vasya can buy a good with the index 𝑖
, and also choose exactly 𝑘−1
 goods, the price of which does not exceed 𝑎𝑖
, if he currently has enough coins (i.e 𝑝≥𝑎𝑖
). Thus, he buys all these 𝑘
 goods, and his number of coins decreases by 𝑎𝑖
 (i.e it becomes 𝑝:=𝑝−𝑎𝑖
).
Please note that each good can be bought no more than once.

For example, if the store now has 𝑛=5
 goods worth 𝑎1=2,𝑎2=4,𝑎3=3,𝑎4=5,𝑎5=7
, respectively, 𝑘=2
, and Vasya has 6
 coins, then he can buy 3
 goods. A good with the index 1
 will be bought by Vasya without using the offer and he will pay 2
 coins. Goods with the indices 2
 and 3
 Vasya will buy using the offer and he will pay 4
 coins. It can be proved that Vasya can not buy more goods with six coins.

Help Vasya to find out the maximum number of goods he can buy.


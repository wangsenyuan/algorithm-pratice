After passing a test, Vasya got himself a box of 𝑛
candies. He decided to eat an equal amount of candies each morning until there are no more candies. However, Petya also
noticed the box and decided to get some candies for himself.

This means the process of eating candies is the following: in the beginning Vasya chooses a single integer 𝑘
, same for all days. After that, in the morning he eats 𝑘
candies from the box (if there are less than 𝑘
candies in the box, he eats them all), then in the evening Petya eats 10%
of the candies remaining in the box. If there are still candies left in the box, the process repeats — next day Vasya
eats 𝑘
candies again, and Petya — 10%
of the candies left in a box, and so on.

If the amount of candies in the box is not divisible by 10
, Petya rounds the amount he takes from the box down. For example, if there were 97
candies in the box, Petya would eat only 9
of them. In particular, if there are less than 10
candies in a box, Petya won't eat any at all.

Your task is to find out the minimal amount of 𝑘
that can be chosen by Vasya so that he would eat at least half of the 𝑛
candies he initially got. Note that the number 𝑘
must be integer.


It's Petya's birthday party and his friends have presented him a brand new "Electrician-𝑛
" construction set, which they are sure he will enjoy as he always does with weird puzzles they give him.

Construction set "Electrician-𝑛
" consists of 2𝑛−1
 wires and 2𝑛
 light bulbs. Each bulb has its own unique index that is an integer from 1
 to 2𝑛
, while all wires look the same and are indistinguishable. In order to complete this construction set one has to use each of the wires to connect two distinct bulbs. We define a chain in a completed construction set as a sequence of distinct bulbs of length at least two, such that every two consecutive bulbs in this sequence are directly connected by a wire. Completed construction set configuration is said to be correct if a resulting network of bulbs and wires has a tree structure, i.e. any two distinct bulbs are the endpoints of some chain.

Petya was assembling different configurations for several days, and he noticed that sometimes some of the bulbs turn on. After a series of experiments he came up with a conclusion that bulbs indexed 2𝑖
 and 2𝑖−1
 turn on if the chain connecting them consists of exactly 𝑑𝑖
 wires. Moreover, the following important condition holds: the value of 𝑑𝑖
 is never greater than 𝑛
.

Petya did his best but was not able to find a configuration that makes all bulbs to turn on, so he seeks your assistance. Please, find out a configuration that makes all bulbs shine. It is guaranteed that such configuration always exists.

### ideas
1. 这个题目比较晦涩难懂。就是 2 * n - 1 和 2 * n 中间需要 di条线
2. 感觉把它们变成一条line也是ok的？
3. 如果一条线不行，再找几个branch就可以了
4. 先把1,3,5,7连在一起（但是得排个序，di大的放在两头）
5. 然后根据di，去放置另外一个端点即可
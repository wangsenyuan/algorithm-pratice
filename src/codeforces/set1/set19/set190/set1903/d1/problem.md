Theofanis really likes to play with the bits of numbers. He has an array 𝑎
of size 𝑛
and an integer 𝑘
. He can make at most 𝑘
operations in the array. In each operation, he picks a single element and increases it by 1
.

He found the maximum bitwise AND that array 𝑎
can have after at most 𝑘
operations.

Theofanis has put a lot of work into finding this value and was very happy with his result. Unfortunately, Adaś, being
the evil person that he is, decided to bully him by repeatedly changing the value of 𝑘
.

Help Theofanis by calculating the maximum possible bitwise AND for 𝑞
different values of 𝑘
. Note that queries are independent.

### thoughts

1. 这个貌似也不是简单的bit by bit（从高到低），假设已经处理完了第x位，那么在增加1次操作的情况下，有可能让x位变成0
2. 应该是在保证当前结果有效的情况下，看需要多少次操作
3. 如果能够让当前位为1，就使用这些操作，否则继续
Monocarp has got two strings 𝑠
 and 𝑡
 having equal length. Both strings consist of lowercase Latin letters "a" and "b".

Monocarp wants to make these two strings 𝑠
 and 𝑡
 equal to each other. He can do the following operation any number of times: choose an index 𝑝𝑜𝑠1
 in the string 𝑠
, choose an index 𝑝𝑜𝑠2
 in the string 𝑡
, and swap 𝑠𝑝𝑜𝑠1
 with 𝑡𝑝𝑜𝑠2
.

You have to determine the minimum number of operations Monocarp has to perform to make 𝑠
 and 𝑡
 equal, and print any optimal sequence of operations — or say that it is impossible to make these strings equal.

 ### ideas
 1. 如果a的个数不是偶数 => false
 2. 在a/b都是偶数的前提下，看怎么分配
 3. 肯定一半的a在s中，一半的b在s中
 4. 排除掉上下一致的情况，剩余的必然是 a/b, b/a这样的结构
 5. 两个 a/b a/b => b/b a/a (在一次操作后)
 6. 会不会出一个 a/b , b/a 种情况
 7. 从个数上是有可能的（似乎就没发进行了？
 8. 如果有一个 b/b 或者 a/a, 那么就可以继续
 9. 如果只有两个是不行的
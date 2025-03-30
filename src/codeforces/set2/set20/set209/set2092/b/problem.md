As soon as Dasha Purova crossed the border of France, the villain Markaron kidnapped her and placed her in a prison under his large castle. Fortunately, the wonderful Lady Bug, upon hearing the news about Dasha, immediately ran to save her in Markaron's castle. However, to get there, she needs to crack a complex password.

The password consists of two bit strings 𝑎
 and 𝑏
, each of which has a length of 𝑛
. In one operation, Lady Bug can choose any index 2≤𝑖≤𝑛
 and perform one of the following actions:

swap(𝑎𝑖
, 𝑏𝑖−1
) (swap the values of 𝑎𝑖
 and 𝑏𝑖−1
), or
swap(𝑏𝑖
, 𝑎𝑖−1
) (swap the values of 𝑏𝑖
 and 𝑎𝑖−1
).
Lady Bug can perform any number of operations. The password is considered cracked if she can ensure that the first string consists only of zeros. Help her understand whether or not she will be able to save the unfortunate Dasha.

### ideas
1. 00/11 
2. 01/01
3. 01/10 (这个好像没法交换)
4. 11/11 
5. 10/01
6. 对于位置i处的0，要么它来自b中的错位，要么来a的相同位置
7. 就是看a中奇数位的1的个数 <= b中偶数位0的个数， good
8. 
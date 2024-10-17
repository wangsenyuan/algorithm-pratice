Ujan has been lazy lately, but now has decided to bring his yard to good shape. First, he decided to paint the path from his house to the gate.

The path consists of 𝑛
 consecutive tiles, numbered from 1
 to 𝑛
. Ujan will paint each tile in some color. He will consider the path aesthetic if for any two different tiles with numbers 𝑖
 and 𝑗
, such that |𝑗−𝑖|
 is a divisor of 𝑛
 greater than 1
, they have the same color. Formally, the colors of two tiles with numbers 𝑖
 and 𝑗
 should be the same if |𝑖−𝑗|>1
 and 𝑛mod|𝑖−𝑗|=0
 (where 𝑥mod𝑦
 is the remainder when dividing 𝑥
 by 𝑦
).

Ujan wants to brighten up space. What is the maximum number of different colors that Ujan can use, so that the path is aesthetic?

### ideas
1. 考虑 n = 6
2. 1, 2, 3, 4, 5, 6
3. 1, 3, 5 (距离为2)
4. 2, 4, 6 (距离为2)
5. 1, 4    (距离为3)
6. 所以只能1种颜色
7. 如果n是奇数, 那么 答案 = n
8. 如果n的质因数，只有一个, 比如2的倍数，但是是2, 如果是3的倍数, 比如9, 就是3
9. 如果有2个以上质因数，那么答案是1
The Master's Assistance Center has announced an entrance exam, which consists of the following.

The candidate is given a set 𝑠
 of size 𝑛
 and some strange integer 𝑐
. For this set, it is needed to calculate the number of pairs of integers (𝑥,𝑦)
 such that 0≤𝑥≤𝑦≤𝑐
, 𝑥+𝑦
 is not contained in the set 𝑠
, and also 𝑦−𝑥
 is not contained in the set 𝑠
.

Your friend wants to enter the Center. Help him pass the exam!


### ideas
1. 0 <=x <= y <= c, and x + y not in the set, y - x not in the set
2. 可以反过来处理吗？就是找哪些在set中,
3. 0 <= x <= y <= c， 不考虑后面的条件, 一共有 (1 + c + 1) * (c + 1) / 2
4. 比如对于3， (0, 0), (0, 1), (0, 2), (0, 3) ... (3, 3) = (4 + 3 + 2 + 1) = 10
5. x + y not in the set and y - x not in the set 的反条件 = x + y， 或者 y - x 在set中
6. 对于set中的数字s, y - x = y + x 的只有当x = 0 的情况
7. 那么这样的数字，正好是set的size (s[0] = 0 可能需要特殊处理)
8. 对于y - x 在集合中的情况，只要对于每个数字s[i], 就是如果 s[i] = x, 那么就排出掉 s[i] + y 就可以了 （选择s[i]的时候)
9. 对于 y + x 在集合中的情况，就是对于数字s[i], 如果数字y =  那么就排除掉 y - s[i]的情况就可以了 （y - s[i] >= 0)
10. 但是可能存在一对 (x, y) (y - x) 和 x + y 都在集合中，这部分，被重复减去了一次，要加回来
11. y - x = a
12. x + y = b => y = (a + b) / 2, x = (b - a) / 2
13. 固定b, 如果b是偶数，那么所有之前的偶数，会产生一对，（x, y)
14. 同理对奇数也是如此
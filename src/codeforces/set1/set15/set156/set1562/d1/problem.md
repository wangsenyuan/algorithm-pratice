Stitch likes experimenting with different machines with his friend Sparky. Today they built another machine.

The main element of this machine are 𝑛
 rods arranged along one straight line and numbered from 1
 to 𝑛
 inclusive. Each of these rods must carry an electric charge quantitatively equal to either 1
 or −1
 (otherwise the machine will not work). Another condition for this machine to work is that the sign-variable sum of the charge on all rods must be zero.

More formally, the rods can be represented as an array of 𝑛
 numbers characterizing the charge: either 1
 or −1
. Then the condition must hold: 𝑎1−𝑎2+𝑎3−𝑎4+…=0
, or ∑𝑖=1𝑛(−1)𝑖−1⋅𝑎𝑖=0
.

Sparky charged all 𝑛
 rods with an electric current, but unfortunately it happened that the rods were not charged correctly (the sign-variable sum of the charge is not zero). The friends decided to leave only some of the rods in the machine. Sparky has 𝑞
 questions. In the 𝑖
th question Sparky asks: if the machine consisted only of rods with numbers 𝑙𝑖
 to 𝑟𝑖
 inclusive, what minimal number of rods could be removed from the machine so that the sign-variable sum of charges on the remaining ones would be zero? Perhaps the friends got something wrong, and the sign-variable sum is already zero. In that case, you don't have to remove the rods at all.

If the number of rods is zero, we will assume that the sign-variable sum of charges is zero, that is, we can always remove all rods.

Help your friends and answer all of Sparky's questions!

### ideas
1. a[i] - a[i+1] + a[i+2] - a[i+3] ... +/- a[j] = 0
2. 那需要知道在i...j中间奇数位有的和是多少，偶数位的和是多少
3. s1, s2, 如果s1 = s2, 那么就不需要删除
4. 如果s1 = 2, s2 = 3, +++++， 这里删掉一个是ok的
5. 但是如果s1 = 3, s2 = -2， 比如+-+-+， 但是这里似乎删掉一个, 中间那个+，  +--+   1 - (-1) + (-1) - 1 = 0
6. +++---   1 - 1 + 1 - （-1） + （-1） - （-1）= 3
7. +++--+   1 - 1 + 1 - （-1） + （-1） - 1 = 0
8. 似乎对于+/-来说，它在奇数位和在偶数位的个数要相同
9. 这个比较容易证明，假设在奇数位有1个+，那么在偶数位必须也要有1个+，否则无法抵消这个奇数位的+
10. 至少得到了一个条件
11. 假设在区间l...r 中间， 奇数位+号的数量是a, 奇数位-号的数量是b
12. 偶数位分别为c/d
13. 那么这里希望的是 a = c, 且 b = d
14. 是不是最多就是2啊？
15. 假设在某个地方删除掉一个字符后，a = c了，是不是一直能够做到的？
16. 删除掉一个字符后，后面的位置就flip了一下，假设这个后面有 a/b,c/d 
17. 现在就变成了 c/d, a/b 了 假设前面 x/y, u/v
18. (x + a)/(b + y), (u + c)/(v + d)
19. 现在变成了 （x + c)/(d + y), (u + a)/(v + b)了
20. 我们希望 x + c = u + a, d + y = v + b
21. x - a = u - c且 y - b = v - d
22. 单独都不好弄
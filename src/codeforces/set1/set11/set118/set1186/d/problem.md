Vus the Cossack has 𝑛
 real numbers 𝑎𝑖
. It is known that the sum of all numbers is equal to 0
. He wants to choose a sequence 𝑏
 the size of which is 𝑛
 such that the sum of all numbers is 0
 and each 𝑏𝑖
 is either ⌊𝑎𝑖⌋
 or ⌈𝑎𝑖⌉
. In other words, 𝑏𝑖
 equals 𝑎𝑖
 rounded up or down. It is not necessary to round to the nearest integer.

For example, if 𝑎=[4.58413,1.22491,−2.10517,−3.70387]
, then 𝑏
 can be equal, for example, to [4,2,−2,−4]
.

Note that if 𝑎𝑖
 is an integer, then there is no difference between ⌊𝑎𝑖⌋
 and ⌈𝑎𝑖⌉
, 𝑏𝑖
 will always be equal to 𝑎𝑖
.

Help Vus the Cossack find such sequence!


### ideas
1. sum(a) = 0
2. a[0] + a[1] + a[2] + ... a[n-1] = 0
3. a[i] = c[i] + d[i] (c[i]是整数部分, d[i]是它的小数部分)
4. 不会做这种～～～
5. 是不是，找出小数绝对值最大的那个数，放在一边，其他的放在另外一边
6. 假设这个数是的小数部分是w, 那么有-w肯定等于右边的部分的小数部分的和吗？
7. 也不大对，因为有正有负
8. 如果右边的小数部分都变成整数的话，必须让他们向离的近的那边移动
9. 
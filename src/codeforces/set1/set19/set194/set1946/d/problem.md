Yarik's birthday is coming soon, and Mark decided to give him an array 𝑎
 of length 𝑛
.

Mark knows that Yarik loves bitwise operations very much, and he also has a favorite number 𝑥
, so Mark wants to find the maximum number 𝑘
 such that it is possible to select pairs of numbers [𝑙1,𝑟1
], [𝑙2,𝑟2
], …
 [𝑙𝑘,𝑟𝑘
], such that:

𝑙1=1
.
𝑟𝑘=𝑛
.
𝑙𝑖≤𝑟𝑖
 for all 𝑖
 from 1
 to 𝑘
.
𝑟𝑖+1=𝑙𝑖+1
 for all 𝑖
 from 1
 to 𝑘−1
.
(𝑎𝑙1⊕𝑎𝑙1+1⊕…⊕𝑎𝑟1)|(𝑎𝑙2⊕𝑎𝑙2+1⊕…⊕𝑎𝑟2)|…|(𝑎𝑙𝑘⊕𝑎𝑙𝑘+1⊕…⊕𝑎𝑟𝑘)≤𝑥
, where ⊕
 denotes the operation of bitwise XOR, and |
 denotes the operation of bitwise OR.
If such 𝑘
 does not exist, then output −1
.

### ideas 
1. 也就是把a分成尽量多的分组，每组xor，然后or后的结果不比x大
2. 从高位到低位,如果x[d] = 1, 那么，似乎对怎么分组没有影响，因为无论怎么分，每组里面也就是出现0/1
3. 但是如果x[d] = 0, 那么，分组后，每组的xor[d]都必须等于0（d假设是x的最高位）
4. 有点眉目了
5. 假设到d位的时候，已经分为了k组，且满足高位到现在 f(k) = x
6. 如果x[d] = 1, 如果按照k组划分后， f(k)[d] = 0, 也就是现在满足了条件，是否k就是最大的数目呢？
7. 如果f(k)[d] = 1, 能否通过再次划分后，得到f(k)[d] = 0?
8. f(d)[d] = 1, 说明有几个分组的xor值=1, 那么此时重新划分，就只能合并不同分组
9. 把一个分组变成2个，是不能出现0的，但是把两个1的合并成一个，是可以得到0的
10. 但是如果对当前的分组进行合并，高位会怎么样？
11. 假设合并最近的两个分组（它们在d位都是1）如果高位也是1，1 =>good
12. 如果高位是（0， 1）=> 合并后还是1，同样是good
13. 如果高位时（0， 0）=》合并后还是0， good
14. 也就是说d位的合并，似乎对高位是没有影响的
15. 从最高位开始，如果x[d] = 1, 且没有分组的情况下，就直接划分位n组，如果这样子 <x => good
16. 否则继续
17. 如果 x[d] = 0, 那么必须将所以相邻的两个1进行合并，以得到0，如果不能 => bad
18. 否则继续
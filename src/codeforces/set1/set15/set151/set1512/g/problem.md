Let us denote by 𝑑(𝑛)
 the sum of all divisors of the number 𝑛
, i.e. 𝑑(𝑛)=∑𝑘|𝑛𝑘
.

For example, 𝑑(1)=1
, 𝑑(4)=1+2+4=7
, 𝑑(6)=1+2+3+6=12
.

For a given number 𝑐
, find the minimum 𝑛
 such that 𝑑(𝑛)=𝑐
.

### ideas
1. 这个也没法二分，比如 f(5) = 1 + 5 = 6 < f(4) = 1 + 2 + 4 = 7
2. 但是这里有个观察是， f(i) >= 1 + i 的 这里 c < 1e7
3. 所以 1+i <= c = 1e7 可以先算出来
4. 但是要求出num的因数，需要sqrt(num)的时间
5. 所以1e7 * 1e3 = 1e10 了
6. 可以反过来
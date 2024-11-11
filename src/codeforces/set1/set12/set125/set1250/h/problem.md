Let's denote the candle representing the digit 𝑑
 as 𝑑
-candle.

Your set contains 𝑐0
 instances of 0
-candles, 𝑐1
 instances of 1
-candles and so on. So, the total number of candles is 𝑐0+𝑐1+⋯+𝑐9
.

These digits are needed to wish your cat a happy birthday. For each birthday, starting with the first, you want to compose the age of the cat using the digits from the set.

Since you light candles for a very short time, candles don't have time to burn out. For this reason you can reuse candles an arbitrary number of times (therefore your set of candles never changes).

For example, if you have one instance of each digit (i.e. 𝑐0=𝑐1=⋯=𝑐9=1
), you can compose any number from 1
 to 10
 using this set, but you cannot compose 11
.

You have to determine the first birthday, on which you cannot compose the age of the cat using the candles from your set. In other words, find the minimum number 𝑦
 such that all numbers from 1
 to 𝑦−1
 can be composed by digits from your set, but 𝑦
 cannot be composed.

 ### ideas
 1. 比如如果有1和0， 那么就可以组合出 10
 2. 如果有两个1，那么可以组合出 11
 3. 类似这样
 4. 所以从长度出发，长度为1的数字，如果缺失任何一个 => find
 5. 所有的数字都有，在两位的情况下, 0必须有1位，但是其他所有的位数，至少有2个
 6. 依次类推
 7. 
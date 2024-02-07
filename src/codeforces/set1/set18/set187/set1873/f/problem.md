Luca is in front of a row of 𝑛
trees. The 𝑖
-th tree has 𝑎𝑖
fruit and height ℎ𝑖
.

He wants to choose a contiguous subarray of the array [ℎ𝑙,ℎ𝑙+1,…,ℎ𝑟]
such that for each 𝑖
(𝑙≤𝑖<𝑟
), ℎ𝑖
is divisible†
by ℎ𝑖+1
. He will collect all the fruit from each of the trees in the subarray (that is, he will collect 𝑎𝑙+𝑎𝑙+1+⋯+𝑎𝑟
fruits). However, if he collects more than 𝑘
fruits in total, he will get caught.

What is the maximum length of a subarray Luca can choose so he doesn't get caught?

†
𝑥
is divisible by 𝑦
if the ratio 𝑥𝑦
is an integer.

### thoughts

1. 定义R[i]表示从i开始，在只考虑h[i] % h[i+1] = 0的情况下，最远能到达的位置
2. R[i] = i if h[i] % h[i+1] != 0 else R[i+1]
3. 假设已经知道了当前l...r是满足条件的，移动到l-1位置时，移动r到合适的位置
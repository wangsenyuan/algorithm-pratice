You're given an array 𝑎
 initially containing 𝑛
 integers. In one operation, you must do the following:

Choose a position 𝑖
 such that 1<𝑖≤|𝑎|
 and 𝑎𝑖=|𝑎|+1−𝑖
, where |𝑎|
 is the current size of the array.
Append 𝑖−1
 zeros onto the end of 𝑎
.
After performing this operation as many times as you want, what is the maximum possible length of the array 𝑎
?

### ideas
1. a[i] + i = n
2. a[j] + j = n + i
3. (a[j] + j) - (a[i] + i) = i
4. a[j] + j = a[i] + 2 * i
5. 
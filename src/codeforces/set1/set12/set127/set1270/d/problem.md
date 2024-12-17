We have hidden an array 𝑎
 of 𝑛
 pairwise different numbers (this means that no two numbers are equal). You can get some information about this array using a new device you just ordered on Amazon.

This device can answer queries of the following form: in response to the positions of 𝑘
 different elements of the array, it will return the position and value of the 𝑚
-th among them in the ascending order.

Unfortunately, the instruction for the device was lost during delivery. However, you remember 𝑘
, but don't remember 𝑚
. Your task is to find 𝑚
 using queries to this device.

You can ask not more than 𝑛
 queries.

Note that the array 𝑎
 and number 𝑚
 are fixed before the start of the interaction and don't depend on your queries. In other words, interactor is not adaptive.

Note that you don't have to minimize the number of queries, and you don't need to guess array 𝑎
. You just have to guess 𝑚
.

### ideas
1. m <= k and k < n
2. 假设询问1....k, 返回i, w 表示 a[i]  = w, 是第m个数 (还不知道m是第几个数)
3. 如果 m = 1, 这种特殊情况，有什么办法吗？
4. 保留i，并替换j一个为k+1。 如果再次返回i,表示什么呢？
5. a[k+1] 和 a[j] 都在i的同一边（要么都大于 a[i]， 要么都小于 a[i])
6. 1...k => i
7. 保留i的时候，但是随便替换一个 1...k+1 => 如果仍然返回i, 
8. 
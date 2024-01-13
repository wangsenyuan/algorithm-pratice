You are given an array of 𝑛
integers, where each integer is either 0
, 1
, or 2
. Initially, each element of the array is blue.

Your goal is to paint each element of the array red. In order to do so, you can perform operations of two types:

    a. pay one coin to choose a blue element and paint it red;
    b. choose a red element which is not equal to 0 and a blue element adjacent to it, decrease the chosen red element by 1, and paint the chosen blue element red.

What is the minimum number of coins you have to spend to achieve your goal?

### thoughts

1. 尽量使用操作2
2. 假设选择了i从蓝色变成了red，
3. 如果a[i] > 0, 那么就可以利用操作2，向左边进行扩展直到向左遇到0
4. 如果a[i]仍然>0, 那么就可以继续利用操作2，先右边进行扩展直到遇到0
5. 那为啥不直接从左端一直处理到右端呢？这样的cost仍然是1
6. 这是因为如果L端是1，那么它没法处理它左边的0
7. 但是如果从某个2开始，向两边，就可以将L-1和R+1给处理掉
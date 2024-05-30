Vasya has two hobbies — adding permutations†
 to arrays and finding the most frequently occurring element. Recently, he found an array 𝑎
 and decided to find out the maximum number of elements equal to the same number in the array 𝑎
 that he can obtain after adding some permutation to the array 𝑎
.

More formally, Vasya must choose exactly one permutation 𝑝1,𝑝2,𝑝3,…,𝑝𝑛
 of length 𝑛
, and then change the elements of the array 𝑎
 according to the rule 𝑎𝑖:=𝑎𝑖+𝑝𝑖
. After that, Vasya counts how many times each number occurs in the array 𝑎
 and takes the maximum of these values. You need to determine the maximum value he can obtain.

†
A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array), and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

### ideas
1. a[i] + p[i] = a[j] + p[j]
2. => a[i] - a[j] = p[i] - p[j]
3. 顺序无关
4. 对a进行排序
5. 假设结果是 a[i] + ? 相等，那么就要看前面有多少个数字和他的差在n-1类的，且不同
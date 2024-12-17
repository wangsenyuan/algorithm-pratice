One common way of digitalizing sound is to record sound intensity at particular time moments. For each time moment intensity is recorded as a non-negative integer. Thus we can represent a sound file as an array of 𝑛
 non-negative integers.

If there are exactly 𝐾
 distinct values in the array, then we need 𝑘=⌈log2𝐾⌉
 bits to store each value. It then takes 𝑛𝑘
 bits to store the whole file.

To reduce the memory consumption we need to apply some compression. One common way is to reduce the number of possible intensity values. We choose two integers 𝑙≤𝑟
, and after that all intensity values are changed in the following way: if the intensity value is within the range [𝑙;𝑟]
, we don't change it. If it is less than 𝑙
, we change it to 𝑙
; if it is greater than 𝑟
, we change it to 𝑟
. You can see that we lose some low and some high intensities.

Your task is to apply this compression in such a way that the file fits onto a disk of size 𝐼
 bytes, and the number of changed elements in the array is minimal possible.

We remind you that 1
 byte contains 8
 bits.

𝑘=⌈𝑙𝑜𝑔2𝐾⌉
 is the smallest integer such that 𝐾≤2𝑘
. In particular, if 𝐾=1
, then 𝑘=0
.

### ideas
1. 假设在区间l...r 中间有m个distinct value
2. 那么 k = [log2 m]
3. 共需要 k * n <= I * 8
4. 
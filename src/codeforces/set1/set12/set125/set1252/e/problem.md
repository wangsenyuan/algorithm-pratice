Andi is a mathematician, a computer scientist, and a songwriter. After spending so much time writing songs, he finally writes a catchy melody that he thought as his best creation. However, the singer who will sing the song/melody has a unique vocal range, thus, an adjustment may be needed.

A melody is defined as a sequence of 𝑁
 notes which are represented by integers. Let 𝐴
 be the original melody written by Andi. Andi needs to adjust 𝐴
 into a new melody 𝐵
 such that for every 𝑖
 where 1≤𝑖<𝑁
:

If 𝐴𝑖<𝐴𝑖+1
, then 𝐵𝑖<𝐵𝑖+1
.
If 𝐴𝑖=𝐴𝑖+1
, then 𝐵𝑖=𝐵𝑖+1
.
If 𝐴𝑖>𝐴𝑖+1
, then 𝐵𝑖>𝐵𝑖+1
.
|𝐵𝑖−𝐵𝑖+1|≤𝐾
, i.e. the difference between two successive notes is no larger than 𝐾
.
Moreover, the singer also requires that all notes are within her vocal range, i.e. 𝐿≤𝐵𝑖≤𝑅
 for all 1≤𝑖≤𝑁
.
Help Andi to determine whether such 𝐵
 exists, and find the lexicographically smallest 𝐵
 if it exists. A melody 𝑋
 is lexicographically smaller than melody 𝑌
 if and only if there exists 𝑗
 (1≤𝑗≤𝑁
) such that 𝑋𝑖=𝑌𝑖
 for all 𝑖<𝑗
 and 𝑋𝑗<𝑌𝑗
.

For example, consider a melody 𝐴={1,3,5,6,7,8,9,10,3,7,8,9,10,11,12,12}
 as shown in the following figure. The diagonal arrow up in the figure implies that 𝐴𝑖<𝐴𝑖+1
, the straight right arrow implies that 𝐴𝑖=𝐴𝑖+1
, and the diagonal arrow down implies that 𝐴𝑖>𝐴𝑖+1
.

### ideas
1. 假设 p[i] = 为了满足后面的条件，而能取到的的范围
2. 假设 a[i] .... a[j]是连续递增的
3. 那么这段如果从 l....l + j - i
4. 如果是连续递减的，同样的也可以得到一个范围
5. 但是a[i]的最小值，假设前一段的最后的值是x, 那么 a[i] >= max(L, x - K)
6. 所以，这里重要的貌似是这些峰顶和峰谷
7. 
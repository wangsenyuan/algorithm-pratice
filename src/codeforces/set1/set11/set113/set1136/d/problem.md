At the big break Nastya came to the school dining room. There are 𝑛
 pupils in the school, numbered from 1
 to 𝑛
. Unfortunately, Nastya came pretty late, so that all pupils had already stood in the queue, i.e. Nastya took the last place in the queue. Of course, it's a little bit sad for Nastya, but she is not going to despond because some pupils in the queue can agree to change places with some other pupils.

Formally, there are some pairs 𝑢
, 𝑣
 such that if the pupil with number 𝑢
 stands directly in front of the pupil with number 𝑣
, Nastya can ask them and they will change places.

Nastya asks you to find the maximal number of places in queue she can move forward.

Input
The first line contains two integers 𝑛
 and 𝑚
 (1≤𝑛≤3⋅105
, 0≤𝑚≤5⋅105
) — the number of pupils in the queue and number of pairs of pupils such that the first one agrees to change places with the second one if the first is directly in front of the second.

The second line contains 𝑛
 integers 𝑝1
, 𝑝2
, ..., 𝑝𝑛
 — the initial arrangement of pupils in the queue, from the queue start to its end (1≤𝑝𝑖≤𝑛
, 𝑝
 is a permutation of integers from 1
 to 𝑛
). In other words, 𝑝𝑖
 is the number of the pupil who stands on the 𝑖
-th position in the queue.

The 𝑖
-th of the following 𝑚
 lines contains two integers 𝑢𝑖
, 𝑣𝑖
 (1≤𝑢𝑖,𝑣𝑖≤𝑛,𝑢𝑖≠𝑣𝑖
), denoting that the pupil with number 𝑢𝑖
 agrees to change places with the pupil with number 𝑣𝑖
 if 𝑢𝑖
 is directly in front of 𝑣𝑖
. It is guaranteed that if 𝑖≠𝑗
, than 𝑣𝑖≠𝑣𝑗
 or 𝑢𝑖≠𝑢𝑗
. Note that it is possible that in some pairs both pupils agree to change places with each other.

Nastya is the last person in the queue, i.e. the pupil with number 𝑝𝑛
.


### ideas
1. Nastya 只能一个个的往前移动，只有当她移动到位置i的时候，它才能使用i-1处的
2. 但是似乎存在一种把前面的某个u换到这里来，这样子，nastya才能移动到前面去的情况
3. 比如她在最后一个位置，p[n-1]不可和她交换，但是如果p[n-2]愿意和她换，且p[n-2]愿意和p[n-1]交换，那么就可以先交换p[n-2], p[n-1]
4. 似乎有点复杂
5. 假设x到到了位置 i,如果i-1不愿意让步，那么就必须移动i-1到某个愿意让步的u那里去
6. 或者找到愿意让步的u下来
7. 应该反过来，考虑有那些人愿意到后面去
8. 假设x愿意到nastya的后面去，那么x必须愿意和 i+1...n都愿意交换
9. 但是这里有两个情况 x  < y 
10. 在y存在的情况下，x有可能不能往后移动，但是y移动完以后，就满足条件了
11. 所以应该先处理y
Nazar, a student of the scientific lyceum of the Kingdom of Kremland, is known for his outstanding mathematical abilities. Today a math teacher gave him a very difficult task.

Consider two infinite sets of numbers. The first set consists of odd positive numbers (1,3,5,7,…
), and the second set consists of even positive numbers (2,4,6,8,…
). At the first stage, the teacher writes the first number on the endless blackboard from the first set, in the second stage — the first two numbers from the second set, on the third stage — the next four numbers from the first set, on the fourth — the next eight numbers from the second set and so on. In other words, at each stage, starting from the second, he writes out two times more numbers than at the previous one, and also changes the set from which these numbers are written out to another.

The ten first written numbers: 1,2,4,3,5,7,9,6,8,10
. Let's number the numbers written, starting with one.

The task is to find the sum of numbers with numbers from 𝑙
 to 𝑟
 for given integers 𝑙
 and 𝑟
. The answer may be big, so you need to find the remainder of the division by 1000000007
 (109+7
).

Nazar thought about this problem for a long time, but didn't come up with a solution. Help him solve this problem.

### ideas
1. 如果只看 1,2,4,3,5,7,9,6,8,10,12,14,16,18,11,13,15,
2. 1,3,5,7,9,11,13,15,17,19,
3. 2,4,6,8,10,12,14,16,18,20
4. 对于l来说，要找到l1,l2的位置
5. 
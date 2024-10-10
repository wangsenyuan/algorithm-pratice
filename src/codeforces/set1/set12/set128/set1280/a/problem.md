We start with a string 𝑠
 consisting only of the digits 1
, 2
, or 3
. The length of 𝑠
 is denoted by |𝑠|
. For each 𝑖
 from 1
 to |𝑠|
, the 𝑖
-th character of 𝑠
 is denoted by 𝑠𝑖
.

There is one cursor. The cursor's location ℓ
 is denoted by an integer in {0,…,|𝑠|}
, with the following meaning:

If ℓ=0
, then the cursor is located before the first character of 𝑠
.
If ℓ=|𝑠|
, then the cursor is located right after the last character of 𝑠
.
If 0<ℓ<|𝑠|
, then the cursor is located between 𝑠ℓ
 and 𝑠ℓ+1
.
We denote by 𝑠left
 the string to the left of the cursor and 𝑠right
 the string to the right of the cursor.

We also have a string 𝑐
, which we call our clipboard, which starts out as empty. There are three types of actions:

The Move action. Move the cursor one step to the right. This increments ℓ
 once.
The Cut action. Set 𝑐←𝑠right
, then set 𝑠←𝑠left
.
The Paste action. Append the value of 𝑐
 to the end of the string 𝑠
. Note that this doesn't modify 𝑐
.
The cursor initially starts at ℓ=0
. Then, we perform the following procedure:

Perform the Move action once.
Perform the Cut action once.
Perform the Paste action 𝑠ℓ
 times.
If ℓ=𝑥
, stop. Otherwise, return to step 1.
You're given the initial string 𝑠
 and the integer 𝑥
. What is the length of 𝑠
 when the procedure stops? Since this value may be very large, only find it modulo 109+7
.

It is guaranteed that ℓ≤|𝑠|
 at any time.

Input
The first line of input contains a single integer 𝑡
 (1≤𝑡≤1000
) denoting the number of test cases. The next lines contain descriptions of the test cases.

The first line of each test case contains a single integer 𝑥
 (1≤𝑥≤106
). The second line of each test case consists of the initial string 𝑠
 (1≤|𝑠|≤500
). It is guaranteed, that 𝑠
 consists of the characters "1", "2", "3".

It is guaranteed that the sum of 𝑥
 in a single file is at most 106
. It is guaranteed that in each test case before the procedure will stop it will be true that ℓ≤|𝑠|
 at any time.

### ideas
1. 光标的位置，只能一直往后走，前面的不会变的
2. 假设现在的长度是l, 位置是p, 且s[p] = x, 
3. f(l, p, x) => f((l - p) * x + p, p + 1, ?)
4. 这里的关键还是？但是，有个简单的问题是，最终的位置是比较短的，前面的部分，是不是可以一直模拟的，超过的部分，直接舍弃掉，记录长度就可以了？
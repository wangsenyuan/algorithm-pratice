Monocarp is opening his own IT company. He wants to hire 𝑛
 programmers and 𝑚
 testers.

There are 𝑛+𝑚+1
 candidates, numbered from 1
 to 𝑛+𝑚+1
 in chronological order of their arriving time. The 𝑖
-th candidate has programming skill 𝑎𝑖
 and testing skill 𝑏𝑖
 (a person's programming skill is different from their testing skill). The skill of the team is the sum of the programming skills of all candidates hired as programmers, and the sum of the testing skills of all candidates hired as testers.

When a candidate arrives to interview, Monocarp tries to assign them to the most suitable position for them (if their programming skill is higher, then he hires them as a programmer, otherwise as a tester). If all slots for that position are filled, Monocarp assigns them to the other position.

Your task is, for each candidate, calculate the skill of the team if everyone except them comes to interview. Note that it means that exactly 𝑛+𝑚
 candidates will arrive, so all 𝑛+𝑚
 positions in the company will be filled.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤104
) — the number of test cases.

Each test case consists of three lines:

the first line contains two integers 𝑛
 and 𝑚
 (0≤𝑛,𝑚≤2⋅105
; 2≤𝑛+𝑚+1≤2⋅105
) — the number of programmers and the number of testers Monocarp wants to hire, respectively;
the second line contains 𝑛+𝑚+1
 integers 𝑎1,𝑎2,…,𝑎𝑛+𝑚+1
 (1≤𝑎𝑖≤109
), where 𝑎𝑖
 is the programming skill of the 𝑖
-th candidate;
the third line contains 𝑛+𝑚+1
 integers 𝑏1,𝑏2,…,𝑏𝑛+𝑚+1
 (1≤𝑏𝑖≤109
; 𝑏𝑖≠𝑎𝑖
), where 𝑏𝑖
 is the testing skill of the 𝑖
-th candidate.
Additional constraint on the input: the sum of (𝑛+𝑚+1)
 over all test cases doesn't exceed 2⋅105
.

Output
For each test case, print 𝑛+𝑚+1
 integers, where the 𝑖
-th integer should be equal to the skill of the team if everyone except the 𝑖
-th candidate comes to interview.

### ideas
1. 考虑n+m个人怎么分配任务
2. 显然目的是使得编程的技巧+测试的技巧之和最大
3. 这个有个贪心的做法，就是 p[i] - t[i]越大的，越应该分配为编程的，否则分配为测试的
4. 这个贪心的做法不是题目的做法。题目的做法就是按照顺序分配就可以了
5. 所以，就是计算出来她们的参与度，但是有个问题，就是第n+1个人，会比较特殊，它有可能递补为程序员
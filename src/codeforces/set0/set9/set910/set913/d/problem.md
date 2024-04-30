You are preparing for an exam on scheduling theory. The exam will last for exactly T milliseconds and will consist of n
problems. You can either solve problem i in exactly ti milliseconds or ignore it and spend no time. You don't need time
to rest after solving a problem, either.

Unfortunately, your teacher considers some of the problems too easy for you. Thus, he assigned an integer ai to every
problem i meaning that the problem i can bring you a point to the final score only in case you have solved no more than
ai problems overall (including problem i).

Formally, suppose you solve problems p1, p2, ..., pk during the exam. Then, your final score s will be equal to the
number of values of j between 1 and k such that k ≤ apj.

You have guessed that the real first problem of the exam is already in front of you. Therefore, you want to choose a set
of problems to solve during the exam maximizing your final score in advance. Don't forget that the exam is limited in
time, and you must have enough time to solve all chosen problems. If there exist different sets of problems leading to
the maximum final score, any of them will do.

### ideas

1. 顺序没有关系
2. 考虑i和j a[i] < a[j], 不考虑时间的情况下， a[j] 越大越好
3. 这是因为，更大的a[j] 更有可能被算入最后的答案
4. 所以可以按照a[i]倒序排，然后处理，在满足总时间一致的情况下，尽量的加入
5. 假设此时选中了k个任务，然后其中a[i] >= k 的部分被计入
6. 但是这里有个问题，那直接选中那些 a[i] >= k 的部分就可以了，多余的不需要进入
7. 因为多选入一个后，有可能会让某个 a[i] == k - 1 的不满足条件
8. 不行，这个地方k的变化会影响到sum，需要固定s (k = s)
9. 然后其中选择那些时间少的部分
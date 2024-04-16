A multi-subject competition is coming! The competition has 𝑚
different subjects participants can choose from. That's why Alex (the coach) should form a competition delegation among
his students.

He has 𝑛
candidates. For the 𝑖
-th person he knows subject 𝑠𝑖
the candidate specializes in and 𝑟𝑖
— a skill level in his specialization (this level can be negative!).

The rules of the competition require each delegation to choose some subset of subjects they will participate in. The
only restriction is that the number of students from the team participating in each of the chosen subjects should be the
same.

Alex decided that each candidate would participate only in the subject he specializes in. Now Alex wonders whom he has
to choose to maximize the total sum of skill levels of all delegates, or just skip the competition this year if every
valid non-empty delegation has negative sum.

(Of course, Alex doesn't have any spare money so each delegate he chooses must participate in the competition).

### ideas

1. 对于所有的人，按照level从高到低进行排序
2. 然后把当前人加入对应的subject后，进行判断
3. 假设这个人参与了，那么他对应subject的人数正好是这个队列的长度，所有其他队列，且人数>=subject的，都应该被选到
4. 但是这里，就有个问题了，如何找到队列人数是x的，都是正数sum的那批
5. 
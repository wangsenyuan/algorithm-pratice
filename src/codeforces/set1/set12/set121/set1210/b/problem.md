Marcin is a coach in his university. There are 𝑛
students who want to attend a training camp. Marcin is a smart coach, so he wants to send only the students that can
work calmly with each other.

Let's focus on the students. They are indexed with integers from 1
to 𝑛
. Each of them can be described with two integers 𝑎𝑖
and 𝑏𝑖
; 𝑏𝑖
is equal to the skill level of the 𝑖
-th student (the higher, the better). Also, there are 60
known algorithms, which are numbered with integers from 0
to 59
. If the 𝑖
-th student knows the 𝑗
-th algorithm, then the 𝑗
-th bit (2𝑗
) is set in the binary representation of 𝑎𝑖
. Otherwise, this bit is not set.

Student 𝑥
thinks that he is better than student 𝑦
if and only if 𝑥
knows some algorithm which 𝑦
doesn't know. Note that two students can think that they are better than each other. A group of students can work
together calmly if no student in this group thinks that he is better than everyone else in this group.

Marcin wants to send a group of at least two students which will work together calmly and will have the maximum possible
sum of the skill levels. What is this sum?

### ideas

1. 对于一个group，如果有一个学生他比其他学生都强，那么这个group就不是calmly的.
2. 但是如果除了x，有另外一个学生和他一样强，那么x就不会有问题
3. 所以，这是一图，最好是有环的图
4. 如果x在一个环上，且选中了x，那么x所在的环，都应该被选中
5. 如果x一个人占了一个位，那么这样子是不行的。但是如果这个位上有至少两个人，就是ok的
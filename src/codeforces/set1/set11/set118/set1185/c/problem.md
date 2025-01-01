A session has begun at Beland State University. Many students are taking exams.

Polygraph Poligrafovich is going to examine a group of 𝑛
 students. Students will take the exam one-by-one in order from 1
-th to 𝑛
-th. Rules of the exam are following:

The 𝑖
-th student randomly chooses a ticket.
if this ticket is too hard to the student, he doesn't answer and goes home immediately (this process is so fast that it's considered no time elapses). This student fails the exam.
if the student finds the ticket easy, he spends exactly 𝑡𝑖
 minutes to pass the exam. After it, he immediately gets a mark and goes home.
Students take the exam in the fixed order, one-by-one, without any interruption. At any moment of time, Polygraph Poligrafovich takes the answer from one student.

The duration of the whole exam for all students is 𝑀
 minutes (max𝑡𝑖≤𝑀
), so students at the end of the list have a greater possibility to run out of time to pass the exam.

For each student 𝑖
, you should count the minimum possible number of students who need to fail the exam so the 𝑖
-th student has enough time to pass the exam.

For each student 𝑖
, find the answer independently. That is, if when finding the answer for the student 𝑖1
 some student 𝑗
 should leave, then while finding the answer for 𝑖2
 (𝑖2>𝑖1
) the student 𝑗
 student does not have to go home.

 ### ideas
 1. 如果要让i能够完成考试，那么它前面所有参与作答的人的时间sum <= M - t[i]
 2. 那么只需要让前面花费时间越长的，离开知道这个时间足够
 3. AVL tree是比较合适的
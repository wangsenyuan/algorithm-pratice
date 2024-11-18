There are 𝑛
students standing in a row. Two coaches are forming two teams — the first coach chooses the first team and the second
coach chooses the second team.

The 𝑖
-th student has integer programming skill 𝑎𝑖
. All programming skills are distinct and between 1
and 𝑛
, inclusive.

Firstly, the first coach will choose the student with maximum programming skill among all students not taken into any
team, and 𝑘
closest students to the left of him and 𝑘
closest students to the right of him (if there are less than 𝑘
students to the left or to the right, all of them will be chosen). All students that are chosen leave the row and join
the first team. Secondly, the second coach will make the same move (but all students chosen by him join the second
team). Then again the first coach will make such move, and so on. This repeats until the row becomes empty (i. e. the
process ends when each student becomes to some team).

Your problem is to determine which students will be taken into the first team and which students will be taken into the
second team.

### ideas

1. just simulate the process
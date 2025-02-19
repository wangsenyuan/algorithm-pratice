After the educational reform Polycarp studies only two subjects at school, Safety Studies and PE (Physical Education). During the long months of the fourth term, he received n marks in them. When teachers wrote a mark in the journal, they didn't write in what subject the mark was for, they just wrote the mark.

Now it's time to show the journal to his strict parents. Polycarp knows that recently at the Parent Meeting the parents were told that he received a Safety Studies marks and b PE marks (a + b = n). Now Polycarp wants to write a subject's name in front of each mark so that:

there are exactly a Safety Studies marks,
there are exactly b PE marks,
the total average score in both subjects is maximum.
An average subject grade is the sum of all marks in it, divided by the number of them. Of course, the division is performed in real numbers without rounding up or down. Polycarp aims to maximize the x1 + x2, where x1 is the average score in the first subject (Safety Studies), and x2 is the average score in the second one (Physical Education).


### ideas
1. 假设得分分别是x1, x2
2. x1 * a + x2 * b = sum
3. a + b = n
4. x1 * a + x2 * (n - b) = sum
5. x1 * (a - b) + x2 * n = sum
6. 假设a <= b, 如果不是，交换就可以
7. 这种情况下， 高分似乎分配给SE更好，因为平摊后，得分更高
8. 
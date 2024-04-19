In a galaxy far, far away Lesha the student has just got to know that he has an exam in two days. As always, he hasn't
attended any single class during the previous year, so he decided to spend the remaining time wisely.

Lesha knows that today he can study for at most 𝑎
hours, and he will have 𝑏
hours to study tomorrow. Note that it is possible that on his planet there are more hours in a day than on Earth. Lesha
knows that the quality of his knowledge will only depend on the number of lecture notes he will read. He has access to
an infinite number of notes that are enumerated with positive integers, but he knows that he can read the first note in
one hour, the second note in two hours and so on. In other words, Lesha can read the note with number 𝑘
in 𝑘
hours. Lesha can read the notes in arbitrary order, however, he can't start reading a note in the first day and finish
its reading in the second day.

Thus, the student has to fully read several lecture notes today, spending at most 𝑎
hours in total, and fully read several lecture notes tomorrow, spending at most 𝑏
hours in total. What is the maximum number of notes Lesha can read in the remaining time? Which notes should he read in
the first day, and which — in the second?

### ideas

1. 如果lesha可以准备x, 但不准备 y, 且x > y, 那么通过交换 x, y, 可以得到至少不更差的结果
2. 所以可知lesha 是阅读了 (1, 2, .... x) (1 + x) * x <= a + b
3. 因为对于某个数num <= a + b, 通过使用1....x 中的某些数字，始终可以获得它，所以可以不浪费的获得这个结果
4. 假设 a <= b
5. 如果 a <= x， 那么直接只用a即可，其他的都放在第二天
6. a > x， 那么通过使用尽可能大的数，最后就可以获得一个小于a的数
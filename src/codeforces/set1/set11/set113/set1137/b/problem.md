The new camp by widely-known over the country Spring Programming Camp is going to start soon. Hence, all the team of
friendly curators and teachers started composing the camp's schedule. After some continuous discussion, they came up
with a schedule 𝑠
, which can be represented as a binary string, in which the 𝑖
-th symbol is '1' if students will write the contest in the 𝑖
-th day and '0' if they will have a day off.

At the last moment Gleb said that the camp will be the most productive if it runs with the schedule 𝑡
(which can be described in the same format as schedule 𝑠
). Since the number of days in the current may be different from number of days in schedule 𝑡
, Gleb required that the camp's schedule must be altered so that the number of occurrences of 𝑡
in it as a substring is maximum possible. At the same time, the number of contest days and days off shouldn't change,
only their order may change.

Could you rearrange the schedule in the best possible way?

### ideas

1. kmp
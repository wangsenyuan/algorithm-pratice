Petya studies at university. The current academic year finishes with 𝑛
 special days. Petya needs to pass 𝑚
 exams in those special days. The special days in this problem are numbered from 1
 to 𝑛
.

There are three values about each exam:

𝑠𝑖
 — the day, when questions for the 𝑖
-th exam will be published,
𝑑𝑖
 — the day of the 𝑖
-th exam (𝑠𝑖<𝑑𝑖
),
𝑐𝑖
 — number of days Petya needs to prepare for the 𝑖
-th exam. For the 𝑖
-th exam Petya should prepare in days between 𝑠𝑖
 and 𝑑𝑖−1
, inclusive.
There are three types of activities for Petya in each day: to spend a day doing nothing (taking a rest), to spend a day passing exactly one exam or to spend a day preparing for exactly one exam. So he can't pass/prepare for multiple exams in a day. He can't mix his activities in a day. If he is preparing for the 𝑖
-th exam in day 𝑗
, then 𝑠𝑖≤𝑗<𝑑𝑖
.

It is allowed to have breaks in a preparation to an exam and to alternate preparations for different exams in consecutive days. So preparation for an exam is not required to be done in consecutive days.

Find the schedule for Petya to prepare for all exams and pass them, or report that it is impossible.

### ideas
1. 一门课越早发布，越早准备？还是越早考试，越早准备呢？
2. 如果一门课发布了，在已发布的课程中，先准备马上要考试的
3. 
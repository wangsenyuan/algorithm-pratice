Every December, VK traditionally holds an event for its employees named "Secret Santa". Here's how it happens.

𝑛
 employees numbered from 1
 to 𝑛
 take part in the event. Each employee 𝑖
 is assigned a different employee 𝑏𝑖
, to which employee 𝑖
 has to make a new year gift. Each employee is assigned to exactly one other employee, and nobody is assigned to themselves (but two employees may be assigned to each other). Formally, all 𝑏𝑖
 must be distinct integers between 1
 and 𝑛
, and for any 𝑖
, 𝑏𝑖≠𝑖
 must hold.

The assignment is usually generated randomly. This year, as an experiment, all event participants have been asked who they wish to make a gift to. Each employee 𝑖
 has said that they wish to make a gift to employee 𝑎𝑖
.

Find a valid assignment 𝑏
 that maximizes the number of fulfilled wishes of the employees.

For each test case, print two lines.

In the first line, print a single integer 𝑘
 (0≤𝑘≤𝑛
) — the number of fulfilled wishes in your assignment.

In the second line, print 𝑛
 distinct integers 𝑏1,𝑏2,…,𝑏𝑛
 (1≤𝑏𝑖≤𝑛
; 𝑏𝑖≠𝑖
) — the numbers of employees assigned to employees 1,2,…,𝑛
.

𝑘
 must be equal to the number of values of 𝑖
 such that 𝑎𝑖=𝑏𝑖
, and must be as large as possible. If there are multiple answers, print any.

### ideas
1. 貌似挺简单的呐
2. 还有个限制不能匹配到自己
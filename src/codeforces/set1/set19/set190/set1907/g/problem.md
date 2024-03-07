In the end of the day, Anna needs to turn off the lights in the office. There are 𝑛
lights and 𝑛
light switches, but their operation scheme is really strange. The switch 𝑖
changes the state of light 𝑖
, but it also changes the state of some other light 𝑎𝑖
(change the state means that if the light was on, it goes off and vice versa).

Help Anna to turn all the lights off using minimal number of switches, or say it is impossible.

Input
The first line of input contains a single integer 𝑡
(1≤𝑡≤104
) — the number of test cases. Descriptions of test cases follow.

The first line of each test case contains the integer 𝑛
(2≤𝑛≤105
) — the number of lights.

The second line of each test case contains the string of 𝑛
characters, the initial state of the lights. Character "0" means that the corresponding light is off, and "1" means that
it is on.

The third line of each test case contains 𝑛
integers 𝑎𝑖
(1≤𝑎𝑖≤𝑛
, 𝑎𝑖≠𝑖
) — the switch 𝑖
changes the states of light 𝑖
and light 𝑎𝑖
.

It is guaranteed that sum of 𝑛
over all test cases does not exceed 2⋅105
Output
For each test case output the integer 𝑘
, the minimal number of switches to use, then in the separate line output the list of 𝑘
switches.

If it is impossible to turn off all the lights, output single integer −1
.
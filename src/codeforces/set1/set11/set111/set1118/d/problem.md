Polycarp has to write a coursework. The coursework consists of 𝑚
 pages.

Polycarp also has 𝑛
 cups of coffee. The coffee in the 𝑖
-th cup Polycarp has 𝑎𝑖
 caffeine in it. Polycarp can drink some cups of coffee (each one no more than once). He can drink cups in any order. Polycarp drinks each cup instantly and completely (i.e. he cannot split any cup into several days).

Surely, courseworks are not being written in a single day (in a perfect world of Berland, at least).

Let's consider some day of Polycarp's work. Consider Polycarp drinks 𝑘
 cups of coffee during this day and caffeine dosages of cups Polycarp drink during this day are 𝑎𝑖1,𝑎𝑖2,…,𝑎𝑖𝑘
. Then the first cup he drinks gives him energy to write 𝑎𝑖1
 pages of coursework, the second cup gives him energy to write 𝑚𝑎𝑥(0,𝑎𝑖2−1)
 pages, the third cup gives him energy to write 𝑚𝑎𝑥(0,𝑎𝑖3−2)
 pages, ..., the 𝑘
-th cup gives him energy to write 𝑚𝑎𝑥(0,𝑎𝑖𝑘−𝑘+1)
 pages.

If Polycarp doesn't drink coffee during some day, he cannot write coursework at all that day.

Polycarp has to finish his coursework as soon as possible (spend the minimum number of days to do it). Your task is to find out this number of days or say that it is impossible.

### ideas
1. 二分，
2. 假设需要d天，
3. 假设都喝完了，那么每天最好喝平均的杯数（这样损失最小）
4. avg = n / d, rem = n % d
5. 然后拍成一个矩形，尽量不要浪费a[i]比较多的那些（否则的话，有可能变成0）
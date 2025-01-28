Kevin used to be a participant of Codeforces. Recently, the KDOI Team has developed a new Online Judge called
Forcescode.

Kevin has participated in 𝑛
contests on Forcescode. In the 𝑖
-th contest, his performance rating is 𝑎𝑖
.

Now he has hacked into the backend of Forcescode and will select an interval [𝑙,𝑟]
(1≤𝑙≤𝑟≤𝑛
), then skip all of the contests in this interval. After that, his rating will be recalculated in the following way:

Initially, his rating is 𝑥=0
;
For each 1≤𝑖≤𝑛
, after the 𝑖
-th contest,
If 𝑙≤𝑖≤𝑟
, this contest will be skipped, and the rating will remain unchanged;
Otherwise, his rating will be updated according to the following rules:
If 𝑎𝑖>𝑥
, his rating 𝑥
will increase by 1
;
If 𝑎𝑖=𝑥
, his rating 𝑥
will remain unchanged;
If 𝑎𝑖<𝑥
, his rating 𝑥
will decrease by 1
.
You have to help Kevin to find his maximum possible rating after the recalculation if he chooses the interval [𝑙,𝑟]
optimally. Note that Kevin has to skip at least one contest.

### ideas

1. 删除一段后，增加的 - 减少的要最多
2. diff[i] = 1 if a[i] > a[i-1],
3.         = 0, if a[i] = a[i-1]
4.         = -1 if a[i] < a[i-1]
5. pref[i] = diff[0] + diff[1] + ... diff[i]
6. 假设选择了一段 l...r suf[r] + pref[l] 最大
7. 要找到任何一个位置， 最大的 pref[l]
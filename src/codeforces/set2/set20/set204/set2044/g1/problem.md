A group of 𝑛
 spiders has come together to exchange plushies. Initially, each spider has 1
 plushie. Every year, if spider 𝑖
 has at least one plushie, he will give exactly one plushie to spider 𝑟𝑖
. Otherwise, he will do nothing. Note that all plushie transfers happen at the same time. In this version, if any spider has more than 1
 plushie at any point in time, they will throw all but 1
 away.

The process is stable in the current year if each spider has the same number of plushies (before the current year's exchange) as he did the previous year (before the previous year's exchange). Note that year 1
 can never be stable.

Find the first year in which the process becomes stable.

### ideas
1. 如果有一个环，那么这个环会一直处在稳定状态
2. 最长的不是环的边+1
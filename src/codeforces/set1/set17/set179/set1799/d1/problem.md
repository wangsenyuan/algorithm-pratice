You have a device with two CPUs. You also have 𝑘
programs, numbered 1
through 𝑘
, that you can run on the CPUs.

The 𝑖
-th program (1≤𝑖≤𝑘
) takes 𝑐𝑜𝑙𝑑𝑖
seconds to run on some CPU. However, if the last program we ran on this CPU was also program 𝑖
, it only takes ℎ𝑜𝑡𝑖
seconds (ℎ𝑜𝑡𝑖≤𝑐𝑜𝑙𝑑𝑖
). Note that this only applies if we run program 𝑖
multiple times consecutively — if we run program 𝑖
, then some different program, then program 𝑖
again, it will take 𝑐𝑜𝑙𝑑𝑖
seconds the second time.

You are given a sequence 𝑎1,𝑎2,…,𝑎𝑛
of length 𝑛
, consisting of integers from 1
to 𝑘
. You need to use your device to run programs 𝑎1,𝑎2,…,𝑎𝑛
in sequence. For all 2≤𝑖≤𝑛
, you cannot start running program 𝑎𝑖
until program 𝑎𝑖−1
has completed.

Find the minimum amount of time needed to run all programs 𝑎1,𝑎2,…,𝑎𝑛
in sequence.

### thoughts

1. 加入当前两个cpu的在处理(i, j) (i > j), 新加入i+1时，
2. 转移到 (i + 1, j) 或者(i, i + 1)
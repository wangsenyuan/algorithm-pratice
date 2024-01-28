Ntarsis has received two integers 𝑛
and 𝑘
for his birthday. He wonders how many fibonacci-like sequences of length 𝑘
can be formed with 𝑛
as the 𝑘
-th element of the sequence.

A sequence of non-decreasing non-negative integers is considered fibonacci-like if 𝑓𝑖=𝑓𝑖−1+𝑓𝑖−2
for all 𝑖>2
, where 𝑓𝑖
denotes the 𝑖
-th element in the sequence. Note that 𝑓1
and 𝑓2
can be arbitrary.

For example, sequences such as [4,5,9,14]
and [0,1,1]
are considered fibonacci-like sequences, while [0,0,0,1,1]
, [1,2,1,3]
, and [−1,−1,−2]
are not: the first two do not always satisfy 𝑓𝑖=𝑓𝑖−1+𝑓𝑖−2
, the latter does not satisfy that the elements are non-negative.

Impress Ntarsis by helping him with this task.

### thoughts

1. 序列被第一个和第二个数确定
2. 也就是说，f[0]= x, f[1] = y, 那么 整个序列就已经被确定了
3. num正常情况下都远大于k
4. 可以倒过来验证

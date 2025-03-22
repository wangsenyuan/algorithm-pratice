Learn, learn and learn again — Valera has to do this every day. He is studying at mathematical school, where math is the main discipline. The mathematics teacher loves her discipline very much and tries to cultivate this love in children. That's why she always gives her students large and difficult homework. Despite that Valera is one of the best students, he failed to manage with the new homework. That's why he asks for your help. He has the following task. A sequence of n numbers is given. A prefix of a sequence is the part of the sequence (possibly empty), taken from the start of the sequence. A suffix of a sequence is the part of the sequence (possibly empty), taken from the end of the sequence. It is allowed to sequentially make two operations with the sequence. The first operation is to take some prefix of the sequence and multiply all numbers in this prefix by  - 1. The second operation is to take some suffix and multiply all numbers in it by  - 1. The chosen prefix and suffix may intersect. What is the maximum total sum of the sequence that can be obtained by applying the described operations?

### ideas
1. 重叠的可以不考虑，相当于没有变化
2. 所以分三段，第一段 -1, 第二段1, 第三段 -1
3. 假设迭代r，知道r后面乘以-1后的结果
4. 假设在[0...l]是第一段
5. sum[r] - sum[l] (l+1...r)的和
6. + -sum[l]
7. sum[r] - 2 * sum[l]的最大值
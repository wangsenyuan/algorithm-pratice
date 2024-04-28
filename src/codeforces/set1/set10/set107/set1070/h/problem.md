Polycarp is working on a new operating system called BerOS. He asks you to help with implementation of a file suggestion
feature.

There are 𝑛
files on hard drive and their names are 𝑓1,𝑓2,…,𝑓𝑛
. Any file name contains between 1
and 8
characters, inclusive. All file names are unique.

The file suggestion feature handles queries, each represented by a string 𝑠
. For each query 𝑠
it should count number of files containing 𝑠
as a substring (i.e. some continuous segment of characters in a file name equals 𝑠
) and suggest any such file name.

For example, if file names are "read.me", "hosts", "ops", and "beros.18", and the query is "os", the number of matched
files is 2
(two file names contain "os" as a substring) and suggested file name can be either "hosts" or "beros.18".

Input
The first line of the input contains integer 𝑛
(1≤𝑛≤10000
) — the total number of files.

The following 𝑛
lines contain file names, one per line. The 𝑖
-th line contains 𝑓𝑖
— the name of the 𝑖
-th file. Each file name contains between 1
and 8
characters, inclusive. File names contain only lowercase Latin letters, digits and dot characters ('.'). Any sequence of
valid characters can be a file name (for example, in BerOS ".", ".." and "..." are valid file names). All file names are
unique.

The following line contains integer 𝑞
(1≤𝑞≤50000
) — the total number of queries.

The following 𝑞
lines contain queries 𝑠1,𝑠2,…,𝑠𝑞
, one per line. Each 𝑠𝑗
has length between 1
and 8
characters, inclusive. It contains only lowercase Latin letters, digits and dot characters ('.').

### ideas

1. 因为query的字符更少（小写字母和.）以及位置固定， 可以用query构建trie，用f去匹配query
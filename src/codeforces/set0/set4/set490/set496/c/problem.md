You are given an n × m rectangular table consisting of lower case English letters. In one operation you can completely remove one column from the table. The remaining parts are combined forming a new table. For example, after removing the second column from the table


abcd
edfg
hijk
 

we obtain the table:


acd
efg
hjk
 

A table is called good if its rows are ordered from top to bottom lexicographically, i.e. each row is lexicographically no larger than the following one. Determine the minimum number of operations of removing a column needed to make a given table good.

### ideas
1. just do it from left to right greedy
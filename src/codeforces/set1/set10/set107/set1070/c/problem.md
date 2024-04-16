Buber is a Berland technology company that specializes in waste of investor's money. Recently Buber decided to transfer
its infrastructure to a cloud. The company decided to rent CPU cores in the cloud for 𝑛
consecutive days, which are numbered from 1
to 𝑛
. Buber requires 𝑘
CPU cores each day.

The cloud provider offers 𝑚
tariff plans, the 𝑖
-th tariff plan is characterized by the following parameters:

𝑙𝑖
and 𝑟𝑖
— the 𝑖
-th tariff plan is available only on days from 𝑙𝑖
to 𝑟𝑖
, inclusive,
𝑐𝑖
— the number of cores per day available for rent on the 𝑖
-th tariff plan,
𝑝𝑖
— the price of renting one core per day on the 𝑖
-th tariff plan.
Buber can arbitrarily share its computing core needs between the tariff plans. Every day Buber can rent an arbitrary
number of cores (from 0 to 𝑐𝑖
) on each of the available plans. The number of rented cores on a tariff plan can vary arbitrarily from day to day.

Find the minimum amount of money that Buber will pay for its work for 𝑛
days from 1
to 𝑛
. If on a day the total number of cores for all available tariff plans is strictly less than 𝑘
, then this day Buber will have to work on fewer cores (and it rents all the available cores), otherwise Buber rents
exactly 𝑘
cores this day.

Input
The first line of the input contains three integers 𝑛
, 𝑘
and 𝑚
(1≤𝑛,𝑘≤106,1≤𝑚≤2⋅105
) — the number of days to analyze, the desired daily number of cores, the number of tariff plans.

The following 𝑚
lines contain descriptions of tariff plans, one description per line. Each line contains four integers 𝑙𝑖
, 𝑟𝑖
, 𝑐𝑖
, 𝑝𝑖
(1≤𝑙𝑖≤𝑟𝑖≤𝑛
, 1≤𝑐𝑖,𝑝𝑖≤106
), where 𝑙𝑖
and 𝑟𝑖
are starting and finishing days of the 𝑖
-th tariff plan, 𝑐𝑖
— number of cores, 𝑝𝑖
— price of a single core for daily rent on the 𝑖
-th tariff plan.

Output
Print a single integer number — the minimal amount of money that Buber will pay.

### ideas

1. 假设在day i, 有m组tariff plan能够使用
2. 那么最省钱的方案是按照单价p进行排序，尽量使用便宜的cpu core，直到k
3. 所以需要一个数据结构，要直到前k个cpu core的最小值
4. avl tree？
5. segment tree也是可以的 （on p <= 10e6)
6. 貌似简单点
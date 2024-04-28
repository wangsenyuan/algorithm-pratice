One department of some software company has 𝑛
servers of different specifications. Servers are indexed with consecutive integers from 1
to 𝑛
. Suppose that the specifications of the 𝑗
-th server may be expressed with a single integer number 𝑐𝑗
of artificial resource units.

In order for production to work, it is needed to deploy two services 𝑆1
and 𝑆2
to process incoming requests using the servers of the department. Processing of incoming requests of service 𝑆𝑖
takes 𝑥𝑖
resource units.

The described situation happens in an advanced company, that is why each service may be deployed using not only one
server, but several servers simultaneously. If service 𝑆𝑖
is deployed using 𝑘𝑖
servers, then the load is divided equally between these servers and each server requires only 𝑥𝑖/𝑘𝑖
(that may be a fractional number) resource units.

Each server may be left unused at all, or be used for deploying exactly one of the services (but not for two of them
simultaneously). The service should not use more resources than the server provides.

Determine if it is possible to deploy both services using the given servers, and if yes, determine which servers should
be used for deploying each of the services.

### ideas

1. 考虑最后满足条件时，S1部署在a个server上，S2部署在b个server上
2. 那么S1的server要满足 c[i] >= x1 / a
3. S2的server要满足 c[i] >= x2 / b
4. 然后可以肯定的是，高容量的被用来部署，是个更合理的额选择
5. 假设按照c降序排， c[i]是最后一个可以被使用的机器
6. 那么有 a >= x1 / c[i], b >= x2 / c[i]
7. 那么 a + b <= i就可以了
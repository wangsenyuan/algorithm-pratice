In distant future on Earth day lasts for n hours and that's why there are n timezones. Local times in adjacent timezones
differ by one hour. For describing local time, hours numbers from 1 to n are used, i.e. there is no time "0 hours",
instead of it "n hours" is used. When local time in the 1-st timezone is 1 hour, local time in the i-th timezone is i
hours.

Some online programming contests platform wants to conduct a contest that lasts for an hour in such a way that its
beginning coincides with beginning of some hour (in all time zones). The platform knows, that there are ai people from
i-th timezone who want to participate in the contest. Each person will participate if and only if the contest starts no
earlier than s hours 00 minutes local time and ends not later than f hours 00 minutes local time. Values s and f are
equal for all time zones. If the contest starts at f hours 00 minutes local time, the person won't participate in it.

Help platform select such an hour, that the number of people who will participate in the contest is maximum.

### ideas

1. 如果选取时区i，要看多少个时区的(f, s)能够覆盖它，然后计算这些时区的和
2. 看错了，比赛始终是在第一个时区的（1），但是要确定在哪个时刻比赛
3. 从而满足足够多人的要求
4. 假设在时刻x时进行 （x > 0, x < n), 那么是时区i处的时刻是 y = (x + n - i) % n
5. 那么要求 y >= f, y <= f
6. 如果 x = 1, 在时区1的时间为0，在时区2的时间为n-1
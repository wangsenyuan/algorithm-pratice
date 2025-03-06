In Walrusland public transport tickets are characterized by two integers: by the number of the series and by the number of the ticket in the series. Let the series number be represented by a and the ticket number — by b, then a ticket is described by the ordered pair of numbers (a, b).

The walruses believe that a ticket is lucky if a * b = rev(a) * rev(b). The function rev(x) reverses a number written in the decimal system, at that the leading zeroes disappear. For example, rev(12343) = 34321, rev(1200) = 21.

The Public Transport Management Committee wants to release x series, each containing y tickets, so that at least w lucky tickets were released and the total number of released tickets (x * y) were minimum. The series are numbered from 1 to x inclusive. The tickets in each series are numbered from 1 to y inclusive. The Transport Committee cannot release more than maxx series and more than maxy tickets in one series.

### ideas
1. a * b = rev(a) * rev(b)
2. 在给定a的情况下， a 和 rev(a)是确定的
3. 也就是 a / rev(a)是确定的， rev(b) / b 是确定的
4. b的范围是确定（maxy）
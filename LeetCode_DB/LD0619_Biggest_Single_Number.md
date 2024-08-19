# 619. 只出现一次的最大数字

### 简单

MyNumbers 表：

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | num         | int  |
    +-------------+------+
    该表可能包含重复项（换句话说，在SQL中，该表没有主键）。
    这张表的每一行都含有一个整数。
 

单一数字 是在 MyNumbers 表中只出现一次的数字。

找出最大的 单一数字 。如果不存在 单一数字 ，则返回 null 。

### 示例 1：

    输入：
    MyNumbers 表：
    +-----+
    | num |
    +-----+
    | 8   |
    | 8   |
    | 3   |
    | 3   |
    | 1   |
    | 4   |
    | 5   |
    | 6   |
    +-----+
    输出：
    +-----+
    | num |
    +-----+
    | 6   |
    +-----+
    解释：单一数字有 1、4、5 和 6 。
    6 是最大的单一数字，返回 6 。

### 示例 2：

    输入：
    MyNumbers table:
    +-----+
    | num |
    +-----+
    | 8   |
    | 8   |
    | 7   |
    | 7   |
    | 3   |
    | 3   |
    | 3   |
    +-----+
    输出：
    +------+
    | num  |
    +------+
    | null |
    +------+
    解释：输入的表中不存在单一数字，所以返回 null 。

### 解：

```sql
select 
    if(n.cnt = 1, n.num, null) as num
from(    
    select 
        num,
        count(*) as cnt
    from MyNumbers
    group by num
    order by cnt asc, num desc
    limit 1
) n
```
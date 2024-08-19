# 602. 好友申请 II ：谁有最多的好友

### 中等

    RequestAccepted 表：

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | requester_id   | int     |
    | accepter_id    | int     |
    | accept_date    | date    |
    +----------------+---------+
    (requester_id, accepter_id) 是这张表的主键(具有唯一值的列的组合)。
    这张表包含发送好友请求的人的 ID ，接收好友请求的人的 ID ，以及好友请求通过的日期。
 

编写解决方案，找出拥有最多的好友的人和他拥有的好友数目。

生成的测试用例保证拥有最多好友数目的只有 1 个人。

### 示例 1：

    输入：
    RequestAccepted 表：
    +--------------+-------------+-------------+
    | requester_id | accepter_id | accept_date |
    +--------------+-------------+-------------+
    | 1            | 2           | 2016/06/03  |
    | 1            | 3           | 2016/06/08  |
    | 2            | 3           | 2016/06/08  |
    | 3            | 4           | 2016/06/09  |
    +--------------+-------------+-------------+
    输出：
    +----+-----+
    | id | num |
    +----+-----+
    | 3  | 3   |
    +----+-----+
    解释：
    编号为 3 的人是编号为 1 ，2 和 4 的人的好友，所以他总共有 3 个好友，比其他人都多。

### 进阶：
在真实世界里，可能会有多个人拥有好友数相同且最多，你能找到所有这些人吗？

### 解：

UNION 和 UNION ALL 之间的最主要的区别在于：
- UNION 操作符去除重复的行，UNION ALL 不去除重复的行。
- UNION 操作符默认按照第一个 SELECT 语句的顺序进行排序，UNION ALL 不进行排序。

```sql
select 
    t.id,
    count(*) as num
from(
    select requester_id as id from RequestAccepted
    union all
    select accepter_id as id from RequestAccepted
) t
group by t.id
order by num desc
limit 1
```


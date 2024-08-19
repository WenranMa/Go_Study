# 1045. 买下所有产品的客户

### 中等

Customer 表：

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | customer_id | int     |
    | product_key | int     |
    +-------------+---------+
    该表可能包含重复的行。
    customer_id 不为 NULL。
    product_key 是 Product 表的外键(reference 列)。

Product 表：

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | product_key | int     |
    +-------------+---------+
    product_key 是这张表的主键（具有唯一值的列）。

编写解决方案，报告 Customer 表中购买了 Product 表中所有产品的客户的 id。

### 示例 1：

    输入：
    Customer 表：
    +-------------+-------------+
    | customer_id | product_key |
    +-------------+-------------+
    | 1           | 5           |
    | 2           | 6           |
    | 3           | 5           |
    | 3           | 6           |
    | 1           | 6           |
    +-------------+-------------+
    Product 表：
    +-------------+
    | product_key |
    +-------------+
    | 5           |
    | 6           |
    +-------------+
    输出：
    +-------------+
    | customer_id |
    +-------------+
    | 1           |
    | 3           |
    +-------------+
    解释：
    购买了所有产品（5 和 6）的客户的 id 是 1 和 3 。

### 解：

```sql
select 
    t.customer_id as customer_id
from (
select 
    customer_id,
    count(distinct product_key) as cnt
from Customer
group by customer_id
) t
where t.cnt = (select count(*) as cnt from Product)
```
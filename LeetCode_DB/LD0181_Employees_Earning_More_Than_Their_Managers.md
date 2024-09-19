# 181. 超过经理收入的员工

### 简单

表：Employee 

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    | salary      | int     |
    | managerId   | int     |
    +-------------+---------+
    id 是该表的主键（具有唯一值的列）。
    该表的每一行都表示雇员的ID、姓名、工资和经理的ID。
    
编写解决方案，找出收入比经理高的员工。

以 任意顺序 返回结果表。

### 示例 1:

    输入: 
    Employee 表:
    +----+-------+--------+-----------+
    | id | name  | salary | managerId |
    +----+-------+--------+-----------+
    | 1  | Joe   | 70000  | 3         |
    | 2  | Henry | 80000  | 4         |
    | 3  | Sam   | 60000  | Null      |
    | 4  | Max   | 90000  | Null      |
    +----+-------+--------+-----------+
    输出: 
    +----------+
    | Employee |
    +----------+
    | Joe      |
    +----------+
    解释: Joe 是唯一挣得比经理多的雇员。

### 解：
经理id不为空表示是普通员工，以及经理是谁。
```sql
select e.name as Employee
from Employee e 
left join Employee m on e.managerId = m.id
where e.salary > m.salary
```
# Write your MySQL query statement below
select e.*,
case
    when e.operator = '>' and v1.value > v2.value then 'true'
    when e.operator = '<' and v1.value < v2.value then 'true'
    when e.operator = '=' and v1.value = v2.value then 'true'
    else 'false'
end value
from Expressions as e
join Variables as v1
on e.left_operand = v1.name
join Variables as v2
on e.right_operand = v2.name
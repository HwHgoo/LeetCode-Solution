#!/bin/bash

# 从命令行获取n的值
n=$1

# 遍历输出n行内容
for ((i=1; i<=n; i++))
do
    # 生成数组的长度，范围在[1, 1000]
    array_length=$((RANDOM % 1000 + 1))
    
    # 初始化数组字符串
    array_string="["

    # 遍历生成数组元素
    for ((j=1; j<=array_length; j++))
    do
        # 生成一个随机整数，范围在[-10, 10]
        num=$((RANDOM % 21 - 10))

        # 拼接数组元素到数组字符串中
        if [ $j -eq $array_length ]; then
            array_string="$array_string $num"
        else
            array_string="$array_string $num,"
        fi
    done

    # 结束数组字符串
    array_string="$array_string ]"

    # 输出数组字符串
    echo $array_string
done
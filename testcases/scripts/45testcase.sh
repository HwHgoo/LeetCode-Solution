#!/bin/bash

length=$1

# 验证命令行参数是否有效
if ! [[ $length =~ ^[0-9]+$ ]]; then
    echo "请输入一个有效的数字作为长度参数。"
    exit 1
fi

# 生成数据
data="["
for ((i=1; i<=$length; i++)); do
    # 生成随机数字
    number=$((RANDOM % 1000 + 1))

    # 将数字追加到数据字符串中
    data+=" $number,"
done

# 移除最后一个逗号并添加右括号
data=${data%,}"]"

# 输出生成的数据
echo $data
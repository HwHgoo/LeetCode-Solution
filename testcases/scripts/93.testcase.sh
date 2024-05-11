#!/bin/bash

if [ "$#" -ne 1 ]; then
  echo "Usage: $0 number_of_lines" >&2
  exit 1
fi

n=$1

for (( i=1; i<=$n; i++ ))
do
  num=$(shuf -i 4-12 -n 1) # 生成4到12之间的随机数
  line=""
  for (( j=1; j<=$num; j++ ))
  do
    line="$line$((RANDOM % 10))"
  done
  echo "\"\"$line\"\"" # 输出行并用双引号引用
done
#!/bin/bash

read -r -a val1
read -r -a val2
res1=0
res2=0
for i in $(seq 1 3);
do
    if [[ ${val1[i]} -gt ${val2[i]} ]]; then
	res1=$(($res1 + 1))
    elif [[ ${val2[i]} -gt ${val1[i]} ]]; then
	res2=$(($res2 + 1))
    fi
    done
echo "$res1 $res2"


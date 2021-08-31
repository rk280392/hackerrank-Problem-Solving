#!/bin/bash

read a
if [[ ! $a =~ ^[0-9]+$ ]]; then
    echo "non integer"
    exit
fi
read -r -a ar
for i in ${ar[@]}; do
    sum=$(( $sum + $i ))
done
echo $sum

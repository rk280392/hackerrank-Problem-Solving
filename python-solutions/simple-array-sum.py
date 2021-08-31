#!/usr/bin/python

def simpleArraySum(ar):
    sum = 0
    for i in range(len(ar)):
        sum += ar[i]
    return sum

a = int(input())
ar = []
ar = list(map(int, input().rstrip().split()))

result = simpleArraySum(ar)
print(result)

#!/usr/bin/python3


def compareTriplets(a,b):
    ar = 0
    br = 0
    for i in range(3):
        if a[i] > b[i]:
            ar += 1
        elif b[i] > a[i]:
            br += 1
    return ar,br

a = list(map(int, input().rstrip().split()))
b = list(map(int, input().rstrip().split()))

res = compareTriplets(a, b)
print(' '.join(map(str, res)))


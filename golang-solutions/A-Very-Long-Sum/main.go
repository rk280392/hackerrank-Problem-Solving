package main

import (
    "fmt"
)

func aVeryBigSum(ar []int64) int64 {
    sum := int64(0)
    for i := 0; i < len(ar); i ++ {
	sum += ar[i]
    }
    return sum
}

func main() {
    var a, res int64
    var ar []int64
    fmt.Scan(&a)
    for i := int64(0); i < a; i++ {
	fmt.Scan(&ar)
    }
    res = aVeryBigSum(ar)
    fmt.Println(res)
}

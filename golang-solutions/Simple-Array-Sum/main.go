package main

import (
    "fmt"
)

func simpleArraySum(ar []int) int {
    sum := 0
    for i :=0; i<len(ar); i++ {
	sum += ar[i]
    }
    return sum
}

func main() {
    var a,temp int
    ar := []int{}
    fmt.Scan(&a)
    for i :=0; i<a; i++ {
	fmt.Scan(&temp)
	ar = append(ar, temp)
    }
    res := simpleArraySum(ar)
    fmt.Println(res)
}

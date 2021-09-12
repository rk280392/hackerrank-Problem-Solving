package main

import (
	"fmt"
	"sort"
)

func main() {
    ar := make([]int, 5)
    fmt.Scan(&ar[0], &ar[1], &ar[2], &ar[3], &ar[4])
    sort.Ints(ar)
    x, y := MinMaxSum(ar)
    fmt.Println(x,  y)
}

func MinMaxSum(ar []int) (int,int) {
    min := ar[0] + ar[1] + ar[2] + ar[3]
    max := ar[1] + ar[2] + ar[3] + ar[4]
    return min, max
}

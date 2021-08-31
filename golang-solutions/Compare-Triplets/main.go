package main

import (
    "fmt"
)

func compareTriplets(a [3]int, b [3]int) (int,int) {
    var ar, br int
    for i := 0; i < 3; i++ {
	if a[i] > b[i] {
	    ar++
	} else if b[i] > a[i] {
	    br++
	}
    }
    return ar,br
}

func main() {
    var a,b [3]int
    fmt.Scan(&a[0], &a[1], &a[2])
    fmt.Scan(&b[0], &b[1], &b[2])
    res1,res2:= compareTriplets(a,b)
    fmt.Println(res1,res2)
}



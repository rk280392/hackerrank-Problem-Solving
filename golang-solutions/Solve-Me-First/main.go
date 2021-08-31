package main

import (
    "fmt"
)

func solveMeFirst(a uint32,b uint32) uint32 {
	return a + b
}

func main() {
    var a, b, sum uint32
    fmt.Scanf("%v\n%v",&a,&b)
    sum = solveMeFirst(a,b)
    fmt.Println(sum)
}

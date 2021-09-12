package main

import (
    "fmt"
)

func plusMinus(ar []int) (float64, float64, float64) {
    minusnum := 0
    plusnum := 0
    zeronum := 0
    for i := 0; i < len(ar); i++ {
	if ar[i] < 0 {
	    minusnum++
	} else if ar[i] > 0 {
	    plusnum++
	} else {
	    zeronum++
	}
    }
    minusratio := float64(minusnum) / float64(len(ar))
    plusratio := float64(plusnum) / float64(len(ar))
    zeroratio := float64(zeronum) / float64(len(ar))
    return plusratio, minusratio, zeroratio


}
func main() {
    var a,temp int
    ar := []int{}
    fmt.Scan(&a)
    for i := 0; i < a; i++ {
	fmt.Scan(&temp)
	ar = append(ar, temp)
    }
    b,c,d := plusMinus(ar)
    fmt.Printf("%.6f\n",b)
    fmt.Printf("%.6f\n",c)
    fmt.Printf("%.6f\n",d)

}

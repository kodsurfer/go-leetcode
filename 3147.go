package main

import "fmt"

func main() {
  fmt.Println(maximumEnergy(]int{-2,-3,-1},2))
  fmt.Println(maximumEnergy([]int{5,2,-10,-5,1},3))
}

func maximumEnergy(energy []int, k int) int {
    res := -1<<31;
    for i:=len(energy)-k;i<len(energy);i++{
        s := 0
        for j:=i;j>=0;j-=k {
            s += energy[j]
            res = max(res,s)
        }
    }
    return res
}

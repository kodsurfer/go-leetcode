package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
    res := 0
    subs := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] == 0 {
            subs += 1
        } else {
            subs = 0
        }
        res += subs
    }
    return int64(res)
}

func main() {
  fmt.Println(zeroFilledSubarray([]int{1,3,0,0,2,0,0,4}))
  fmt.Println(zeroFilledSubarray([]int{0,0,0,2,0,0}))
  fmt.Println(zeroFilledSubarray([]int{2,10,2019}))
}

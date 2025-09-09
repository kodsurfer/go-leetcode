package main

import "fmt"

func main() {
  fmt.Println(largestUniqueNumber([]int{5,7,3,9,4,9,8,3,1}))
  fmt.Println(largestUniqueNumber([]int{9,9,8,8}))
}

func largestUniqueNumber(nums []int) int {
    freq := make(map[int]int, len(nums))
    for _, n := range nums {
        _, ok := freq[n]
        if !ok {
            freq[n] = 1
        } else {
            freq[n]++
        }
    }
    maxi := -1
    for k, v := range freq {
        if v == 1 {
            if k > maxi {
                maxi = k
            }
        }
    }
    return maxi
}

package main

import "fmt"

func main() {
  fmt.Println(judgeSquareSum(0))
  fmt.Println(judgeSquareSum(2))
  fmt.Println(judgeSquareSum(4))
}


func judgeSquareSum(c int) bool {
    for i:=0; i*i<=c; i++ {
        b := math.Sqrt(float64(c - i*i))
        if float64(int(b)) == b {
            return true
        }
    }
    return false
}

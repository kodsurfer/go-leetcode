package main

import "fmt"

func main() {
  fmt.Println(iPowrOfFour(16))
  fmt.Println(iPowrOfFour(5))
  fmt.Println(iPowrOfFour(1))
}

func isPowerOfFour(n int) bool {
    if n<=0 {
        return false
    }
    p := math.Log(float64(n))/math.Log(float64(4))
    return math.Abs(p-math.Round(p))<0.0000000001
}

package main

import "fmt"

func main() {
  fmt.Println(isPowerOfThree(0))
  fmt.Println(isPowerOfThree(1))
  fmt.Println(isPowerOfThree(2))
  fmt.Println(isPowerOfThree(3))
}

func isPowerOfThree(n int) bool {
    if n<=0 {
        return false
    }
    r := math.Log(float64(n))/math.Log(float64(3))
    return math.Abs(r - math.Round(r))<0.0000000001
}

package main

import "fmt"

func main() {
  fmt.Println(numWaterBottles(15,4))
  fmt.Println(numWaterBottles(9,3))
}

func numWaterBottles(numBottles int, numExchange int) int {
    drunk := numBottles
    emptys := numBottles
    var neos, remains int
    for ;emptys>=numExchange; {
        neos = int(emptys/numExchange)
        remains = emptys%numExchange
        drunk += neos
        emptys = remains + neos
    }
    return drunk
}

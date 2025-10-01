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

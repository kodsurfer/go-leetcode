func smallestNumber(n int) int {
    x := 1
    for ;; {
        x *= 2
        if x > n {
            break
        }
    }
    return x-1
}

func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }
    p := math.Log(float64(n))/math.Log(float64(2))
    return math.Abs(p - math.Round(p)) < 0.0000000001
}

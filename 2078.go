func maxDistance(colors []int) int {
    res := 0
    for i:=0;i<len(colors)-1;i++ {
        for j:=i+1;j<len(colors);j++ {
            if colors[i]!=colors[j] {
                res = int(math.Max(float64(res), float64(j-i)))
            }
        }
    }
    return res
}

func minTimeToVisitAllPoints(points [][]int) int {
    res := 0
    var tx,ty,cx,cy int
    for i:=0;i<len(points)-1;i++ {
        cx = points[i][0]
        cy = points[i][1]
        tx = points[i+1][0]
        ty = points[i+1][1]
        res += int(math.Max(math.Abs(float64(tx-cx)), math.Abs(float64(ty-cy))))
    }
    return res
}

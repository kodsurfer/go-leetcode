func minMoves(nums []int, limit int) int {
    n := len(nums)
    dif := make([]int, 2*limit+2)
    for i:=0;i<n/2;i++{
        x := min(nums[i],nums[n-1-i])
        y := max(nums[i],nums[n-1-i])
        dif[2] += 2
        dif[x+1] -= 1
        dif[x+y] -= 1
        dif[x+y+1] += 1
        dif[y+limit+1] += 1
    }
    minops := n
    curops := 0
    for j:=2;j<=2*limit;j++{
        curops += dif[j]
        minops = min(minops, curops)
    }
    return minops
}

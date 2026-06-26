func twoSum(nums []int, target int) []int {
    di := make(map[int]int, len(nums))
    for i, n := range nums {
        _, ok := di[target-n]
        if ok {
            return []int{di[target-n], i}
        } else {
            di[n] = i
        }
    }
    return []int{}
}

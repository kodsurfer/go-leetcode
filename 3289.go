func getSneakyNumbers(nums []int) []int {
    m := make(map[int]int, len(nums))
    for i:=0; i<len(nums); i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = 1
        } else {
            m[nums[i]]++
        }
    }
    var res []int
    for k,v := range m {
        if v==2 {
            res = append(res,k)
        }
    }
    return res
}

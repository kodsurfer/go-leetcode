func kLengthApart(nums []int, k int) bool {
    cnt := k
    for _, n := range nums {
        if n == 1 {
            if cnt<k {
                return false
            }
            cnt = 0
        } else {
            cnt += 1
        }
    }
    return true
}

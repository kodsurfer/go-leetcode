func intersect(nums1 []int, nums2 []int) []int {
    var small, big []int
    if len(nums1)<len(nums2){
        small, big = nums1, nums2
    } else {
        small, big = nums2, nums1
    }
    var res []int
    for _, n := range small {
        for i, b := range big {
            if n == b {
                res = append(res, n)
                big = append(big[:i],big[i+1:]...)
                break
            }
        }
    }
    return res
}

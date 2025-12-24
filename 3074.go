func minimumBoxes(apple []int, capacity []int) int {
    sum := 0
    for _, v:= range apple {
        sum += v
    }
    sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
    needs := 0
    for sum > 0 {
        sum -= capacity[needs]
        needs++
    }
    return needs
}

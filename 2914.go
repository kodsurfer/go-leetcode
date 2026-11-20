func minChanges(s string) int {
    cur := s[0]
    conscnt := 0
    minchg := 0
    for i := range s {
        if s[i] == cur {
            conscnt++
            continue
        }
        if conscnt % 2 == 0 {
            conscnt = 1
        } else {
            conscnt = 0
            minchg++
        }
        cur = s[i]
    }
    return minchg
}

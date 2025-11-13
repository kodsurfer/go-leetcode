func maxOperations(s string) int {
    res := 0
    cnt1 := 0
    for i:=0; i<len(s); i++ {
        if s[i]=='0' {
            for ; i+1<len(s) && s[i+1]=='0'; {
                i++
            }
            res += cnt1
        } else {
            cnt1++
        }
    }
    return res
}

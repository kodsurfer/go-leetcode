func sumAndMultiply(n int) int64 {
    if n == 0 {
        return 0
    }
    st := ""
    var summ int
    ns := strconv.Itoa(n)
    for i := range ns {
        if ns[i] != '0' {
            st += string(ns[i])
            tmp, _ := strconv.Atoi(string(ns[i]))
            summ += tmp
        }
    }
    sn, _ := strconv.Atoi(st)
    return int64(sn * summ)
}

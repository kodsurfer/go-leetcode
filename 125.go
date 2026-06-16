func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    le := 0
    ri := len(s)-1
    for ; le<ri; {
        for le<ri && !isalnum(s[le]) {
            le++
        }
        for le<ri && !isalnum(s[ri]) {
            ri--
        }
        if s[le] != s[ri] {
            return false
        }
        le++
        ri--
    }
    return true
}

func isalnum(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

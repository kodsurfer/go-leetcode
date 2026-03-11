func bitwiseComplement(n int) int {
        if n == 0 {
            return 1
        }
        
        x, b := n, 1
        
        for ; x>0; {
            n = n ^ b
            b = b << 1
            x = x >> 1
        }

        return n
}

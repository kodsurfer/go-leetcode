package main

import "fmt"

func main() {
        fmt.Println(hasSameDigits("3902))
        fmt.Println(hasSameDigits("34789))
}

func hasSameDigits(s string) bool {
        x := []byte(s)
        n := len(s)
        for i:=1; i<=n-2; i++ {
            for j:=0; j<=n-i-1; j++ {
                x[j]=byte(((int(x[j]-'0')+int(x[j+1]-'0'))%10) +'0')
            }
        }
        return x[0]==x[1]
}

package main

import (
    "fmt"
    "strconv"
)

func maximum69Number(num int) int {
    s := []rune(fmt.Sprintf("%d",num))
    for i:=0; i<len(s); i++ {
        if s[i] == '6' {
            s[i] = '9'
            break
        }
    }
    res, _ := strconv.Atoi(string(s))
    return res
}

func main(){
  fmt.Println(maximum69Number(9669))
  fmt.Println(maximum69Number(9996)
  fmt.Println(maximum69Number(9999)
}

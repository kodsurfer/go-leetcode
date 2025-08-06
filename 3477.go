package main

func main() {
    fmt.Println(numOfUnplacedFruits([1,2,3],[3,2,1]))
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
    cnt := 0
    for _, fru := range fruits {
        unst := 1
        for i:=0;i<len(baskets);i++{
            if fru <=baskets[i] {
                baskets[i] = 0
                unst = 0
                break
            }
        }
        cnt += unst
    }
    return cnt
}

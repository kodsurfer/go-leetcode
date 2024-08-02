package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
	fmt.Println(countSeniors([]string{"1313579440F2036", "2921522980M5644"}))
}

func countSeniors(details []string) int {
	cnt := 0
	var tens, ones int
	for _, d := range details {
		tens, _ = strconv.Atoi(string(d[11]))
		if tens < 6 {
			continue
		}
		ones, _ = strconv.Atoi(string(d[12]))
		if tens*10+ones > 60 {
			cnt++
		}
	}
	return cnt
}

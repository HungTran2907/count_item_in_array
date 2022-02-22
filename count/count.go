package main

import (
	"fmt"
)

func main() {
	ar := [...]int{10, 20, 30, 40, 50, 60, 10, 30, 20, 10, 20, 30}
	b := make([]int32, len(ar))
	for i := 0; i < len(ar); i++ {
		b[i] = 1
	}
	for i := 0; i < len(ar); i++ {
		count := 1
		if b[i] == 1 {
			b[i] = 0
			for j := i + 1; j < len(ar); j++ {
				if ar[j] == ar[i] {
					count++
					b[j] = 0
				}
			}
			fmt.Printf("number %v appear :%v\n", ar[i], count)
		}
	}
}

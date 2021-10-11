package main

import (
	"fmt"
	"math/rand"
)

func Genslice(n int) []int {
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
	}
	return arr
}

func main() {
	arr := Genslice(100)
	fmt.Println(arr)
}

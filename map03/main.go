package main

import (
	"fmt"
	"math/rand"
)

func Genslice(n int) []int {
	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr
}

func Getdiff(arr []int) int {
	m := make(map[int]bool)
	for _, i := range arr {
		fmt.Println(len(m))
		m[i] = true
	}
	return len(m)
}

func main() {
	arr := Genslice(100)
	fmt.Println(Getdiff(arr))
}

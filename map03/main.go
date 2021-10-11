package main

import (
	"fmt"
	"math/rand"
)

func Genslice(n int) []int {
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(128))
		fmt.Println(cap(arr))
	}
	return arr
}

func Getdiff(arr []int) int {
	m := make(map[int]bool)
	for _, i := range arr {
		m[i] = true
	}
	return len(m)
}

func main() {
	arr := Genslice(100)
	fmt.Println(Getdiff(arr))

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
}

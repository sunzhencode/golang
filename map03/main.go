package main

import (
	"fmt"
	"math/rand"
)

//slice扩容规则实验
func Genslice(n int) []int {
	arr := make([]int, 0, 10)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(128))
		//fmt.Println(cap(arr))
	}
	return arr
}

//1.创建一个初始长度为0、容量为10的int型切片，调用rand.Intn(128)100次，往切片里面添加100个元素，利用map统计该切片里有多少个互不相同的元素。
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

package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s1 := a[:4] // [0,1,2,3,4]
	s2 := a[:4] // [0,1,2,3,4]
	s1[0] = 10
	fmt.Println(a[:]) // [10,2,3,4,5]  修改切片会影响到底级数组，反之亦然
	s2[0] = 20
	fmt.Println(s1) // [20,2,3,4]  修改切片会影响到基于同一底级数组的切片

	s1 = append(s1, []int{6, 7}...) //append 时容量不够时会导致重新分配内存，与原底级数组关系断开
	s1 = append(s1,8,9)
	s1[0] = 100
	fmt.Println(s1)   // [100,2,3,4,6,7]
	fmt.Println(a[:]) //[20,2,3,4,5]
	fmt.Println(s2)   //[20,2,3,4]
}

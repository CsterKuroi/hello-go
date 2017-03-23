package main

import "fmt"

func main() {
     slice := []int32{1, 2, 3, 4, 5, 6}
     slice2 := slice[:2]
     _ = append(slice2, 50, 60, 70, 80, 90)
     fmt.Printf("slice为：%v\n", slice)
     fmt.Printf("操作的切片：%v\n", slice2)
     _ = append(slice2, 50, 60)
     fmt.Printf("slice为：%v\n", slice)
     fmt.Printf("操作的切片：%v\n", slice2)
}

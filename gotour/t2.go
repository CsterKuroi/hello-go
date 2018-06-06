package main

import "fmt"

func main() {
	t1 := []int{11, 12, 13, 14, 15}
	t1 = append(t1, 0)

	for i, _ := range t1 {
		if i < len(t1)-1 {
			t2 := t1[i+1]
			t1[i+1] = t1[0]
			t1[0] = t2
		}
	}
	t1[0] = 0
	fmt.Println(t1)
}

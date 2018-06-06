package main

import "fmt"

func main() {
	t1 := make([]int, 2, 3)
	t1[0] = 1
	t1[1] = 2
	fmt.Println(t1)
	for i, _ := range t1 {
		if i < cap(t1)-1 {
			if i == len(t1)-1 {
				t1 = append(t1, t1[0])
			}
			t2 := t1[i+1]
			t1[i+1] = t1[0]
			t1[0] = t2
		}
	}
	t1[0] = 0
	fmt.Println(t1)
}

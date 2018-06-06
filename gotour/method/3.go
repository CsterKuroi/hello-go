package main

//import "fmt"
//
//type I interface {
//	M()
//}
//
//type J interface {
//	M()
//}
//
//type T struct {
//	S string
//}
//
//// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
//func (t T) I() {
//	fmt.Println("1")
//}
//func (t T) M() {
//	fmt.Println("1")
//}
//
//func main() {
//	var i I = T{"hello"}
//	i.M()
//	var j J = T{"heJllo"}
//	j.M()
//}
package main

import "fmt"

/**
 * @Author wangzhs
 * @create 2021/4/6 15:39
* slice 创建
*/
func main() {

	slice1 := []int{1, 23, 3}
	fmt.Println("len = ", len(slice1), "slice1 = ", slice1)

	var slice2 []int
	slice2 = make([]int, 3)
	fmt.Println("len = ", len(slice2), "slice2 = ", slice2)

	slice3 := make([]int, 3)
	fmt.Println("len = ", len(slice3), "slice3 = ", slice3)

	slice4 := append(slice2, 2)
	fmt.Println(slice4)

	slice5 := slice4[0:2]
	fmt.Println(slice5)

	var slice6 []int = make([]int, 3)
	copy(slice6, slice5)
	fmt.Println(slice6)

	var slice7 []int
	fmt.Println("cap = ", cap(slice7))

	slice7 = append(slice7, 1)
	fmt.Println("cap = ", cap(slice7))

}

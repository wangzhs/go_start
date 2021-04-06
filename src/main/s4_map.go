package main

import "fmt"

/**
 * @Author wangzhs
 * @create 2021/4/6 16:13
 */
func main() {

	var map1 map[string]string
	fmt.Printf("%T", map1)

	map1 = make(map[string]string, 1)
	map1["one"] = "123123"

	fmt.Println(map1)

	var map2 map[string]string = make(map[string]string, 1)
	//map2 := make(map[string]string, 1)
	map2["dd"] = "sdadas"
	fmt.Println(map2)

	map3 := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}
	fmt.Println(map3)

	delete(map3, "0")
	fmt.Println(map3)
	map3["3"] = "4"
	fmt.Println(map3)

}

package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author wangzhs
 * @create 2021/4/7 11:32
 */
func main() {
	user := User{1, "eli"}
	reflectProcess(user)
}

func reflectProcess(org interface{}) {
	of := reflect.TypeOf(org)

	fmt.Println(of)

	value := reflect.ValueOf(org)
	fmt.Println(value)

	for i := 0; i < of.NumField(); i++ {
		fmt.Println("type =>", of.Field(i).Name)
		fmt.Println("value =>", value.Field(i).Interface())
	}
}

type User struct {
	Id   int
	Name string
}

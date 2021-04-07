package main

import "fmt"

/**
 * @Author wangzhs
 * @create 2021/4/7 10:03
 */
func main() {
	// printInterface("111")

	var animal Animal
	animal = &Cat{"red"}
	printAnimal(animal)
}

func printAnimal(animal Animal) {
	animal.Sleep()
}

func printInterface(org interface{}) {

	fmt.Printf("%v\n %T", org, org)

}

type Animal interface {
	Sleep()
	Type()
}

type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("this cat sleep")
}

func (cat *Cat) Type() {
	fmt.Println("this cat")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("this dog sleep")
}

func (dog *Dog) Type() {
	fmt.Println("this dog")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

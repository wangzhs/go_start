package main

/**
 * @author  wangzhs
 * @date  2021/4/6 13:44
 */
func main() {
	print(deferRerunFunc())
}

func deferRerunFunc() int {
	defer deferFunc()
	return returnFunc()
}

func returnFunc() int {
	return 10
}

func deferFunc() int {
	return 20
}

package main

import "fmt"

func main() {
	test01()
	test02()
	test03()
}

func test01() {
	// 假如还有97天放假，问：xx个星期零xx天
	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Println(week, "个星期零", day, "天")
}

func test02() {
	// 定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.5
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度 %v 对应的摄氏温度为 %v \n", huashi, sheshi)
}
func test03() {
	// 有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
}

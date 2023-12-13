package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成5个数，并将其反转打印
	test41()

	// 将斐波那契数列的前50个数放到切片中
	test42(50)

	// 定义一个二维数组用来保存三个班，每个班五名同学的成绩，并求出每个班的平均分和所有班级的平均分
	test43()

	// 随机生成10个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，求最大值、最大值的下标、最小值、最小值的下标，并查找里面是否有55
	test44()

	// 有个升序好的数组，现在插入一个数，要求此数组依然是升序
	test49()

	// 定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
	test50()
}

func test41() {
	// 随机生成5个数，并将其反转打印
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	var arr [5]int
	for index, _ := range arr {
		arr[index] = rand.Intn(10000) + 1
	}
	fmt.Println("此时生成是数组是：", arr)

	// 将生成的数组反转
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("此时生成是数组是：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test42(num int) {
	// 将斐波那契数列的前50个数放到切片中
	var slice1 []uint64 = make([]uint64, num)
	slice1[0] = 1
	slice1[1] = 1
	for i := 2; i < len(slice1); i++ {
		slice1[i] = slice1[i-1] + slice1[i-2]
	}
	fmt.Println("此时生成的切片是：", slice1)
	// 分割线
	println("-------------分割线--------------")
}

func test43() {
	// 定义一个二维数组用来保存三个班，每个班五名同学的成绩，并求出每个班的平均分和所有班级的平均分
	var arr = [3][5]float64{{18, 29, 73, 45, 78}, {17, 24, 34, 46, 65}, {14, 52, 63, 44, 20}}
	// 每个班的成绩和
	var arr2 [3]float64
	// 每个班级的平均分
	var arr3 [3]float64
	for index1, value1 := range arr {
		var sum float64 = 0
		for _, value2 := range value1 {
			sum += value2
		}
		arr2[index1] = sum
		arr3[index1] = sum / float64(len(value1))
	}
	fmt.Println("每个班的成绩和是：", arr2)
	fmt.Println("每个班的平均分是：", arr3)
	// 分割线
	println("-------------分割线--------------")
}

func test44() {
	// 随机生成10个整数(1-100的范围)保存到数组，并倒序打印以及求平均值，求最大值、最大值的下标，并查找里面是否有55
	// 设置纳秒级种子
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	// 组装数组
	for index, _ := range arr {
		arr[index] = rand.Intn(100) + 1
	}
	fmt.Println("此时生成是数组是：", arr)
	// 倒序打印
	test45(arr)
	// 求最大值
	test46(arr)
	// 求平均值
	test47(arr)
	// 查找里面是否有55
	test48(arr, 55)
}
func test45(arr [10]int) {
	// 倒序打印
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	fmt.Println("倒叙打印数组为：", arr)
	// 分割线
	println("-------------分割线--------------")
}

func test46(arr [10]int) {
	// 求最大值及其坐标
	var max int = arr[0]
	var maxIndex int = 0
	for index, value := range arr {
		if value > max {
			max = value
			maxIndex = index
		}
	}
	fmt.Printf("最大值是：%v,最大值的下标是：%v \n", max, maxIndex)
	// 分割线
	println("-------------分割线--------------")
}

func test47(arr [10]int) {
	// 求平均值
	var sum int = 0
	for _, value := range arr {
		sum += value
	}
	fmt.Printf("平均值是：%v \n", float64(sum)/float64(len(arr)))
	// 分割线
	println("-------------分割线--------------")
}

func test48(arr [10]int, num int) {
	// 查找数组是否含有某个数 使用二分法
	// 将数组进行从小到大排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("从小到大排序后的数组是: ", arr)
	// 使用二分法查找
	var leftIndex int = 0
	var rightIndex int = len(arr) - 1
	var midIndex int = (leftIndex + rightIndex) / 2
	for leftIndex <= rightIndex {
		if arr[midIndex] == num {
			fmt.Printf("找到了，下标是：%v \n", midIndex)
			break
		} else if arr[midIndex] > num {
			rightIndex = midIndex - 1
			midIndex = (leftIndex + rightIndex) / 2
		} else {
			leftIndex = midIndex + 1
			midIndex = (leftIndex + rightIndex) / 2
		}
	}
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
	}
	// 分割线
	println("-------------分割线--------------")
}

func test49() {
	// 有个升序好的数组，现在插入一个数，要求此数组依然是升序
	var arr = [5]int{1, 2, 3, 4, 5}
	var num int = 3
	var index int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > num {
			index = i
			break
		}
	}
	var arr2 [6]int
	for i := 0; i < len(arr2); i++ {
		if i < index {
			arr2[i] = arr[i]
		} else if i == index {
			arr2[i] = num
		} else {
			arr2[i] = arr[i-1]
		}
	}
	fmt.Println("插入后的数组是：", arr2)
	// 分割线
	println("-------------分割线--------------")
}

func test50() {
	// 定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
	var arr [3][4]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%v行第%v列的值：", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println("此时生成的数组是：", arr)
	// 将四周的数据清0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == len(arr)-1 || j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}
	fmt.Println("此时生成的数组是：", arr)
	// 分割线
	println("-------------分割线--------------")
}

package main

import (
	"fmt"
	zhoulin "github.com.qq4215279/StudyGo/GoTeaCode/day05/10calc"

)



var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行!")
	fmt.Println(x, pi)
}

func main() {
	ret := zhoulin.Add(10, 20)
	fmt.Println(ret)

}

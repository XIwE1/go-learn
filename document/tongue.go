package main

import (
    // 标准库导入
    "fmt"
	"math"
	// 本地导入
    "myproject/middleware"
	// "./middleware" x 必须指定项目名与 go.mod 里的 module一样
	// "./middleware/hello.go" x 只能以package的形式统一引入不能单独使用
)

// 接口是 Go 最强大的特性之一，提供了一种定义行为而不涉及实现细节的方式。

// 接口定义
type Measurable interface {
	Area() float64
}

// 用于实现接口的结构体
type Circle struct {
	Radius float64
}

// func (接收者) 方法名(参数列表) 返回值类型
// (c Circle) 是接收者，表示这个方法属于 Circle 类型，Area 是方法名。
// 相当于给circle类型的对象绑定了一个area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 任意实现了Measurable接口的类型都可以调用的函数
func processShape(m Measurable) {
	fmt.Printf("Area: %f\n", m.Area())
}

func main() {
	fmt.Println(middleware.Greet("xiwei"))

	circle := Circle{Radius: 5}
	processShape(circle)

	rectangle := Rectangle{Width: 10, Height: 5}
	processShape(rectangle)
}
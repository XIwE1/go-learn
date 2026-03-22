// 声明包
package middleware

import "fmt"  // 导入 fmt 包用于格式化 I/O

// 以函数的形式调用
// Go 使用大小写来控制可见性
func Greet(name string) string {
	return fmt.Sprintf("hello,%s!", name)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为0")
	}
	return a / b, nil
}


func multiply(a, b int) (result int) {
	result = a * b
	return // 裸返回 = 自动补全result
}

// 可变参数
func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

// 自定义错误类型
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("验证错误 %s: %s", e.Field, e.Message)
}

// 带自定义错误的函数
func validateAge(age int) error {
    if age < 0 {
        return ValidationError{Field: "age", Message: "年龄不能为负数"}
    }
    if age > 150 {
        return ValidationError{Field: "age", Message: "年龄似乎不现实"}
    }
    return nil
}

// main
func main() {
	// 基本类型
	var name string = "xiwei"
	age := 25 // 短变量声明 类型推断
	isActive := true

	// 数组-固定大小
	// var fixedArray [5]int = [5]int{1, 2, 3, 4, 5}
	// 元素类型为 interface{} = （可以容纳任意类型）
	// var mixedArray [4]interface{} = [4]interface{}{1, isActive, age, name}

	// 切片-动态数组
	scores := []int{85, 92, 78}
	scores = append(scores, 95)  // 切片添加新元素

	// Map
    // map[string]interface{} // Go 的内建映射类型（类似于 JS 的对象或 Map）。
    // string // 表示 key 必须是字符串类型。
    // interface{} // 表示 value 可以是任意类型（因为 interface{} 是空接口，任何类型都实现了它）。
	person := map[string]interface{}{"name": "Bob", "age": 25}
	student := map[string]interface{}{
		"name": name,
		"age": age,
		"working": isActive,
		"score": scores[0],
	}

    // 结构体（类似对象但有定义的结构）
    // type Person struct {
    //     Name     string
    //     Age      int
    //     IsActive bool
    // }
    
    // person := Person{
    //     Name:     "Bob",
    //     Age:      25,
    //     IsActive: true,
    // }

	// if else 语句
	temperature := 12

	if temperature > 30 {
		fmt.Println("好热")
	} else if temperature > 20 {
		fmt.Println("不热不冷")
	} else {
		fmt.Println("好冷")
	}

	// for 循环
	for i := 0; i < 5; i++ {
		fmt.Printf("循环%d次\n", i + 1)
	}

	// 类似for ... of ...
	for i, score := range scores {
		fmt.Printf("第%d人 成绩为%d\n", i + 1, score)
	}

	for key, value := range person {
		fmt.Printf("键%s: 值%v\n", key, value)
	}

	fmt.Println("这里有个学生，它信息是-", student)
	fmt.Printf("成绩单里一共有%d个信息\n", len(scores))

	fmt.Println("sum=总成绩", sum(scores...))
	fmt.Println("sum=年龄+总成绩", sum(append(scores, age)...))
	fmt.Println("multiply=年龄*总成绩", multiply(sum(scores...), age))

	fmt.Printf("平均成绩为：%d / %d = %f \n", sum(scores...), len(scores), float64(sum(scores...)) / float64(len(scores)))

	// Go 没有像 JavaScript 那样的 typeof，但可以使用反射
    fmt.Printf("age 类型: %T\n", age) // int
	fmt.Println(Greet(name))

	// Go 使用显式的错误返回值
	if err := validateAge(-5); err != nil {
		fmt.Printf("验证错误: %v\n", err)
    }

	if result, err := divide(float64(age), 2); err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)

	}

}


// 函数变量 约等于 变量表达式
var add = func(a, b int) int {
	return a + b
}



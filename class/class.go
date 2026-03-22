// Go 使用组合而不是继承，通过结构体嵌入实现
package class

import (
	"fmt"
)

// 基础结构体
type Animal struct {
    Name string
}

type User struct {
    Name string `json:"name"`   // `xx`结构体标签（Struct Tag）。它是一种元数据（关于数据的数据），用来为结构体的字段提供额外的信息
	// 当把 User 结构体转换成 JSON 字符串（序列化）时，字段 Name 在 JSON 中应该使用键名 "name"，而不是默认的字段名 "Name"。
	// 当把 JSON 字符串解析成 User 结构体（反序列化）时，JSON 中的 "name" 键对应的值应该赋值给结构体的 Name 字段。
    Age  int    `json:"age"`

	// 类似的还有
	// gorm:"column:user_name;type:varchar(100)" 用来指定数据库表中的列名和类型
}

// 给 Animal 结构体绑定了一个名为 Speak 的方法
func (a Animal) Speak() string {
    return fmt.Sprintf("%s makes a sound", a.Name)
}

// 在 Dog 中嵌入 Animal
type Dog struct {
    Animal // 嵌入结构体
}

// 覆盖Animal的Speak方法
func (d Dog) Speak() string {
    return fmt.Sprintf("%s barks", d.Name)
}

func (d Dog) Fetch() string {
    return fmt.Sprintf("%s fetches the ball", d.Name)
}

func main() {
    dog := Dog{Animal{Name: "Rex"}}
	fmt.Println(dog.Speak()) // "Rex barks"
    fmt.Println(dog.Fetch()) // "Rex fetches the ball"

	// 通过嵌入实现接口满足
	type Speaker interface {
		Speak() string
	}
	
	var speakers []Speaker = []Speaker{dog}
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}
}
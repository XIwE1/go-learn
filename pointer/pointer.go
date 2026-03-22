package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	person1 := Person{Name: "John", Age: 20}
	// 这样获取的是值拷贝，修改person2不会影响person1
	person2 := person1

	person2.Name = "Jane"
	fmt.Println(person1)  // {John 20}
		fmt.Println(person2)  // {Jane 20}

	// 获取指针地址
	person3 := &person1
	// 指针拷贝
	person4 := person3
	person3.Name = "Jim"  // 通过指针修改person1的Name
	fmt.Println(person1)  // {Jim 20}
	fmt.Println(person3)  // &{Jim 20}
	fmt.Println(person4)  // &{Jim 20}

    modifyByValue(person1)     // 值传递
    modifyByPointer(&person1)  // 指针传递

	// 基本指针操作
	var value int = 42
	var ptr *int = &value // 获取 value 的地址
	// 通过指针修改值
	*ptr = 100 // 解引用并赋值
	// 多级指针
	var ptrToPtr **int = &ptr
	// 不同类型的指针
	var str string = "Hello"
	var strPtr *string = &str
}

func modifyByValue(person Person) {
	person.Name = "modified"
	fmt.Println("函数内值传递: ", person)
}

func modifyByPointer(person *Person) {
	person.Name = "pointer"
	fmt.Println("函数内指针传递: ", person)
}





// 最佳实践与陷阱
// 指针比较
// 陷阱：指针比较比较的是地址，不是值
fmt.Printf("指针相等: %t\n", user1 == user2) // false
    
// 正确：比较值
fmt.Printf("值相等: %t\n", *user1 == *user2) // true
// 小结构体 使用值传递或指针传递都可以
func (c *Counter) Increment() {}
func (c Counter) GetCount() int {}

// 大结构体 使用指针传递避免拷贝
func (c *Counter) ProcessData() int {}

// 函数创建变量
// 陷阱1: 返回局部变量的指针
func CreateUser() *User {
    // 危险：返回栈上变量的指针
    // user := User{Name: "John", Age: 30}
    // return &user // Go 编译器会自动将其移到堆上，但这不是最佳实践

	return &User{Name: "John", Age: 30} // 直接在堆上创建
}

// 使用指针前判空
func goodProcessUser(user *User) {
    // 最佳实践：总是检查 nil
    if user == nil {
        fmt.Println("用户为空")
        return
    }
    fmt.Println(user.Name)
}

// 在循环中创建指针切片
func goodCreateUserPointers() []*User {
    var users []*User
    names := []string{"Alice", "Bob", "Charlie"}
    
    for _, name := range names {
		// 陷阱：所有指针指向同一个变量
		// user := User{Name: name, Age: 25}
        // users = append(users, &user) 

        // 方法1：使用局部变量副本
        userName := name
        user := User{Name: userName, Age: 25}
        users = append(users, &user)
        
        // 方法2：直接在堆上创建
        // users = append(users, &User{Name: name, Age: 25})
    }
    return users
}

// 指针接收者 - 可以修改原结构体
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor  // 修改原结构体
    r.Height *= factor // Go 自动解引用，等价于 (*r).Width
}

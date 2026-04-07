简洁性、高效性和并发性设计的现代编程语言

### 特点

- **并发编程**：掌握 Goroutines 和 Channels 编写高效的并发代码
- **系统编程**：构建高性能的系统级应用程序
- **微服务**：创建可扩展的微服务架构
- **云原生开发**：在现代云环境中部署应用程序
- **性能优化**：在垃圾回收的基础上实现接近 C 语言的性能

静态语言 强类型 编译执行（ast -> 机器码）

**并发处理** Goroutines 和 Channels

### 运行 go

`go run`  命令有两种使用方式:

1. `go run hello.go` - 运行单个文件
2. `go run .` - **推荐** 运行整个包（当前目录）

`go build hello.go` - **构建可执行文件**

### 包 依赖

通过 go mod 的 .mod 文件管理依赖

- **`go run`**：编译并运行 Go 程序
- **`go build`**：将 Go 程序编译为可执行文件
- **`go test`**：运行测试
- **`go fmt`**：格式化 Go 代码
- **`go mod`**：管理依赖项
- **`go get`**：下载并安装包
- **`go vet`**：分析代码中的潜在问题
- **`go doc`**：生成文档

// 初始化模块

// go mod init my-project

```
// Go: go mod 命令
// 初始化模块
go mod init my-project

// 添加依赖
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql@v1.7.1

// 移除未使用的依赖
go mod tidy

// 下载依赖
go mod download

// 验证依赖
go mod verify

// 列出依赖
go list -m all

// 更新依赖
go get -u github.com/gin-gonic/gin
go get -u all
```

// 添加依赖

格式: {host}/{owner}/{repo}

```go
// go get github.com/gin-gonic/gin

// go get github.com/go-sql-driver/mysql
```

Go: 标准项目结构

```
my-project/
├── go.mod
├── go.sum
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   └── user.go
│   └── models/
│       └── user.go
├── pkg/
│   ├── math/
│   │   └── math.go
│   └── utils/
│       └── helpers.go
├── api/
│   └── routes.go
├── tests/
│   └── math_test.go
└── README.md
```

// 同一文件夹的所有文件使用相同包名（通常与文件夹名相同 但不强制）

```go
// file: middleware/auth.go
package middleware
// file: middleware/cors.go
package middleware
// file: middleware/logging.go
package middleware
```

// 使用时只需要导入一次包

```
// 不需要指定具体文件
import "my-project/internal/middleware"

import (
    // 标准库导入
    "fmt"
    "net/http"
    "time"

    // 第三方导入
    "github.com/gin-gonic/gin"
    "github.com/go-sql-driver/mysql"

    // 本地导入
    "./internal/handlers"
    "./pkg/math"
)
```

// 带别名的导入

```
import (
    mymath "./pkg/math"
)
```

// 带点的导入（将所有名称带入当前作用域）

```
import (
    . "./pkg/math" // 现在可以直接使用 Add() 而不是 math.Add()
)
```

// 带下划线的导入（仅用于副作用）

```
import (
    _ "github.com/go-sql-driver/mysql" // 注册 MySQL 驱动
)
```

![QQ_1755320487908.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/75ab9018de514258b91640aec4d03b65~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAg6KW_57u0:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiMTM3MjY1NDM4OTQzMzQ5NiJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1766861773&x-orig-sign=FhfHRG1q51HKMAS7l9Pv%2F0k0aoo%3D)

### 变量 函数

Go 使用静态类型，具有类型推断能力

```go
var name string = "Go"
age := 25 // 短变量声明（类型推断 局部作用域）
var x int = 10 // 显示类型声明
{ x := 10 } // 块级作用域


func add(a, b int) int {
    return a + b
}
```

### 语法

#### 控制语句

**if** 后直接跟表达式，不需要分号

`if age > 18`

且需要显式的表达式

`if age > 0` 不支持 `if age { xxx }`

**for**

```go
// 传统 for 循环（Go 的唯一循环结构）
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 基于 range 的 for 循环（类似 for...of）
arr := []int{1, 2, 3, 4, 5}
for i, val := range arr {
    fmt.Printf("索引 %d: %d\n", i, val)
}
// 遍历映射
for key, value := range person {
    fmt.Printf("%s: %v\n", key, value)
}
```

**... 操作符** 将切片展开为独立的参数

**append(scores, age)** 创建新切片

**len(scrores)** 获取切片长度

```go
scores := []int{85, 92, 78}
fmt.Println("sum=总成绩", sum(scores...))
```

#### 指针

&符号跟变量 等于获取目标内存地址

`var ptr *int = &value`

#### 断言

错误的断言会导致使用出现 panic

```go
// 带逗号 ok 惯用法的类型断言
if str, ok := value.(string); ok {
    return str
}
```

### 类型 接口

go 的数字类型更精细化 `int float uint`...

go 的数组分为固定和切片 `[n]T` 和 `[]T`

go 的对象预先定义好了构造所以更明确 使用 `struct` `map[K]V`

```go
    // 结构体（类似对象但有定义的结构）
    type Person struct {
        Name     string
        Age      int
        IsActive bool
    }

    person := Person{
        Name:     "Bob",
        Age:      25,
        IsActive: true,
    }
```

go 的函数名为 `func`

go 独有的类型 通道(并发通信)`chan T` 接口(多态支持)`interface {}` 指针(内存地址引用)`*T`

go 的空值为 `nil`

![QQ_1755320526183.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/d22848b390614317b787fbb6c5a63297~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAg6KW_57u0:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiMTM3MjY1NDM4OTQzMzQ5NiJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1766861773&x-orig-sign=UxSano6eBa4OB7t6PiRWyD5gCTE%3D)

```
import "fmt"

func main() {
    // 映射声明和初始化
    var m map[string]int = make(map[string]int)
    m["name"] = 1
    m["age"] = 30

    // 或使用字面量语法
    person := map[string]interface{}{
        "name": "John",
        "age":  30,
        "city": "New York",
    }

    // 访问值
    fmt.Println(person["name"]) // "John"

    // 检查键是否存在
    if age, exists := person["age"]; exists {
        fmt.Printf("Age: %v\n", age)
    }

    // 添加/更新值
    person["country"] = "USA"

    // 删除值
    delete(person, "age")

    // 遍历映射
    for key, value := range person {
        fmt.Printf("%s: %v\n", key, value)
    }
}

// Go: 结构体
package main

import "fmt"

// 结构体定义
type Person struct {
    Name string
    Age  int
    City string
}

// 结构体上的方法
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// 带指针接收者的方法 (可以修改结构体)
func (p *Person) SetAge(age int) {
    p.Age = age
}

func main() {
    // 创建结构体实例
    person := Person{
        Name: "John",
        Age:  30,
        City: "New York",
    }

    // 访问字段
    fmt.Println(person.Name) // "John"

    // 调用方法
    fmt.Println(person.Greet()) // "Hello, I'm John"

    // 修改结构体
    person.SetAge(31)
    fmt.Println(person.Age) // 31

    // 结构体嵌入 (组合)
    type Employee struct {
        Person
        Salary float64
    }

    emp := Employee{
        Person: Person{Name: "Jane", Age: 25, City: "Boston"},
        Salary: 50000,
    }

    // 访问嵌入结构体的方法
    fmt.Println(emp.Greet()) // "Hello, I'm Jane"
}
```

#### 接口

接口是 Go 最强大的特性之一，提供了一种定义行为而不涉及实现细节的方式。**提供了一些动态特性**

```go
// Go: 显式接口
package main

import (
    "fmt"
    "math"
)

// 接口定义
type Measurable interface {
    Area() float64
}

// 实现接口的结构体
type Circle struct {
    Radius float64
}

// 方法实现
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

// 与任何实现 Measurable 的类型一起工作的函数
func processShape(m Measurable) {
    fmt.Printf("Area: %.2f\n", m.Area())
}

func main() {
    circle := Circle{Radius: 5}
    rectangle := Rectangle{Width: 4, Height: 6}

    // 两种类型都实现了 Measurable 接口
    processShape(circle)    // Area: 78.54
    processShape(rectangle) // Area: 24.00

    // 接口满足是隐式的
    var m Measurable = circle
    fmt.Printf("Type: %T, Area: %.2f\n", m, m.Area())
}


// Go: 接口组合
package main

import "fmt"

// 基础接口
type Movable interface {
    Move() string
}

type Drawable interface {
    Draw() string
}

// 组合接口
type GameObject interface {
    Movable
    Drawable
}

// 实现
type Sprite struct {
    X, Y int
}

func (s *Sprite) Move() string {
    s.X += 1
    s.Y += 1
    return fmt.Sprintf("Moving to (%d, %d)", s.X, s.Y)
}

func (s Sprite) Draw() string {
    return fmt.Sprintf("Drawing at (%d, %d)", s.X, s.Y)
}

// 与任何 GameObject 一起工作的函数
func updateGameObject(obj GameObject) {
    fmt.Println(obj.Move())
    fmt.Println(obj.Draw())
}

func main() {
    sprite := &Sprite{X: 0, Y: 0}
    updateGameObject(sprite)

    // 接口组合允许灵活的设计
    var movable Movable = sprite
    var drawable Drawable = sprite
    var gameObj GameObject = sprite

    fmt.Println(movable.Move())  // 只能调用 Move()
    fmt.Println(drawable.Draw()) // 只能调用 Draw()
    fmt.Println(gameObj.Move())  // 可以调用 Move() 和 Draw()
    fmt.Println(gameObj.Draw())
}
```

#### 错误处理

区别 js 的 try catch ，go 的错误显式返回

```go
func main() {
    safeDivide(10, 2) // 结果: 5
    safeDivide(10, 0) // 错误: 除零错误

    if err := validateAge(-5); err != nil {
        fmt.Printf("验证错误: %v\n", err)
    }

    if err := validateAge(200); err != nil {
        fmt.Printf("验证错误: %v\n", err)
    }
}
```

#### 指针

重要特性，它允许程序直接访问和操作内存地址。

关键原则：

- 小对象传值，大对象传指针
- 需要修改时使用指针
- 总是检查 nil 指针
- 在循环中小心指针的生命周期
- 使用指针提高性能，但要注意内存安全

go 的 = 号传值是深度的，如果不加指针标识相当于一个新的数据

```go
    // 值类型
    person1 := Person{Name: "John", Age: 30}
    person2 := person1 // 值拷贝

    fmt.Println("person1:", person1.Name) // "John"
    fmt.Println("person2:", person2.Name) // "John"

    person2.Name = "Jane"
    fmt.Println("person1:", person1.Name) // "John" - person1 未改变
    fmt.Println("person2:", person2.Name) // "Jane"

    // 指针类型
    ptr1 := &person1 // 获取 person1 的地址
    ptr2 := ptr1     // 指针拷贝，指向同一内存地址

    fmt.Println("通过 ptr1:", ptr1.Name) // "John"
    fmt.Println("通过 ptr2:", ptr2.Name) // "John"

    ptr2.Name = "Alice" // 通过指针修改
    fmt.Println("person1:", person1.Name) // "Alice" - 原变量被修改
    fmt.Println("通过 ptr1:", ptr1.Name) // "Alice"
```

存在多级指针

```go
// 基本指针操作
var value int = 42
var ptr *int = &value // 获取 value 的地址

// 多级指针
var ptrToPtr **int = &ptr
```

函数接收指针类型的参数才能修改原结构体，不然修改的是副本

```go
// 值接收者 - 不能修改原结构体
func (r Rectangle) ScaleByValue(factor float64) Rectangle {
    r.Width *= factor  // 只修改副本
    r.Height *= factor
    return r // 返回修改后的副本
}

// 指针接收者 - 可以修改原结构体
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor  // 修改原结构体
    r.Height *= factor // Go 自动解引用，等价于 (*r).Width
}
```

常见陷阱

```go
// 陷阱1: 返回局部变量的指针
func badCreateUser() *User {
    // 危险：返回栈上变量的指针
    user := User{Name: "John", Age: 30}
    return &user // Go 编译器会自动将其移到堆上，但这不是最佳实践
}

// 最佳实践：明确使用堆分配
func goodCreateUser() *User {
    return &User{Name: "John", Age: 30} // 直接在堆上创建
}

// 陷阱2: 空指针解引用
func badProcessUser(user *User) {
    // 危险：没有检查 nil
    fmt.Println(user.Name) // 如果 user 是 nil，程序会 panic
}

func goodProcessUser(user *User) {
    // 最佳实践：总是检查 nil
    if user == nil {
        fmt.Println("用户为空")
        return
    }
    fmt.Println(user.Name)
}

// 陷阱3: 在循环中创建指针切片
func badCreateUserPointers() []*User {
    var users []*User
    names := []string{"Alice", "Bob", "Charlie"}

    for _, name := range names {
        user := User{Name: name, Age: 25}
        users = append(users, &user) // 陷阱：所有指针指向同一个变量
    }
    return users
}

func goodCreateUserPointers() []*User {
    var users []*User
    names := []string{"Alice", "Bob", "Charlie"}

    for _, name := range names {
        // 方法1：使用局部变量副本
        userName := name
        user := User{Name: userName, Age: 25}
        users = append(users, &user)

        // 方法2：直接在堆上创建
        // users = append(users, &User{Name: name, Age: 25})
    }
    return users
}

// 陷阱：指针比较比较的是地址，不是值
fmt.Printf("指针相等: %t\n", user1 == user2) // false

// 正确：比较值
fmt.Printf("值相等: %t\n", *user1 == *user2) // true
```

### 并发和 Goroutine

Go 提供了一种简单而强大的并发编程方法，其座右铭是"不要通过共享内存来通信；要通过通信来共享内存"。

Goroutine 是 Go 处理并发操作的方式。它们是由 Go 运行时管理的轻量级线程

Goroutine 使用  `go`  关键字后跟函数调用来创建

```go
// Go: Goroutine
package main

import (
	"fmt"
	"time"
)

func simulateTask(name string, duration time.Duration) {
	fmt.Printf("%s 开始\n", name)
	time.Sleep(duration)
	fmt.Printf("%s 完成\n", name)
}

func main() {
	fmt.Println("主程序开始")

	// 启动 goroutine（并发执行）
	go simulateTask("任务 A", 2*time.Second)
	go simulateTask("任务 B", 1*time.Second)

	// 等待 goroutine 完成
	time.Sleep(3 * time.Second)

	fmt.Println("主程序结束")
}

// Go: 基于 channel 的同步
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, done chan<- int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("工作线程 %d 完成\n", id)
	done <- id // 发送结果到 channel
}

func main() {
	fmt.Println("启动工作线程...")

	// 创建用于同步的 channel
	done := make(chan int, 3)

	// 启动多个工作线程
	for i := 0; i < 3; i++ {
		go worker(i, done)
	}

	// 等待所有线程完成
	var results []int
	for i := 0; i < 3; i++ {
		result := <-done // 从 channel 接收
		results = append(results, result)
	}

	fmt.Println("所有工作线程完成:", results)
}
```

Channel 是 Go 中 goroutine 间通信的主要机制。它们提供了一种安全的数据共享方式，无需显式锁定。

// 创建一个 channel

ch := make(chan string)

```
// Go: Channel 通信
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 channel
	ch := make(chan string)

	// 发送数据的 goroutine
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "来自发送者的问候" // 发送数据到 channel
	}()

	// 主 goroutine 接收数据
	message := <-ch // 从 channel 接收数据
	fmt.Println("收到:", message)
}
```

缓冲

```
// Go: 缓冲 vs 无缓冲 channel
package main

import (
	"fmt"
	"time"
)

func main() {
	// 无缓冲 channel（同步）
	unbuffered := make(chan string)

	// 缓冲 channel（在缓冲区大小内异步）
	buffered := make(chan string, 2)

	// 无缓冲 channel 示例
	go func() {
		fmt.Println("发送到无缓冲 channel...")
		unbuffered <- "Hello" // 这会阻塞直到有人接收
		fmt.Println("已发送到无缓冲 channel")
	}()

	time.Sleep(100 * time.Millisecond)
	message := <-unbuffered
	fmt.Println("从无缓冲 channel 收到:", message)

	// 缓冲 channel 示例
	fmt.Println("发送到缓冲 channel...")
	buffered <- "第一条消息"  // 不会阻塞
	buffered <- "第二条消息" // 不会阻塞
	fmt.Println("已发送到缓冲 channel")

	// 从缓冲 channel 接收
	fmt.Println("收到:", <-buffered)
	fmt.Println("收到:", <-buffered)
}
```

select 等待第一个完成的任务

```go

	// Select 等待第一个可用的 channel
	select {
	case msg := <-chA:
		fmt.Println("从 A 收到:", msg)
	case msg := <-chB:
		fmt.Println("从 B 收到:", msg)
	case msg := <-chC:
		fmt.Println("从 C 收到:", msg)
	}
```

### kratos

```
/api         # Protobuf 定义
/cmd         # 入口文件
/configs     # 配置文件
/internal    # 私有包
  /data      # 数据层（DB/Cache）
  /biz       # 业务逻辑
  /service   # 服务实现
/pkg         # 公共库
```

### rpc

远程过程调用 RPC：
A 服务器 （方法调用服务） 调用 B 服务器的方法 （方法部署服务），

### gRPC

实现 RPC 数据传输协议 protobuff（纯二进制 数据压缩），跟语言和平台无关，基于 http2.0

gRPC 是一种**高性能、跨语言的开源 RPC（远程过程调用）框架**，核心作用是为分布式系统提供**高效通信**。

- 基于  **HTTP/2**（多路复用、头部压缩）
- 默认使用  **Protobuf（Protocol Buffers）**   二进制编码 + 数据压缩

1.  **微服务通信**

    - 替代 REST，解决服务间**高频调用性能瓶颈**（延迟降低 30%-80%）

1.  **移动端与后端交互**

    - 节省带宽（二进制压缩），提升弱网环境效率

1.  **实时数据流**

    - 流式传输支持**实时日志/消息推送/金融行情**等场景

1.  **云原生基础设施**

    - Kubernetes、Etcd 等核心组件使用 gRPC 通信

**典型场景**：支付系统（微服务间交易处理）、游戏服务器（实时状态同步）、IoT 设备（低带宽数据传输）。

### 依赖注入

**依赖注入（Dependency Injection）**   是一种**对象间解耦的设计模式**，其核心思想是：**由外部容器（而非对象自身）负责创建并注入对象所需的依赖项**。

是构建**可维护、可测试、可扩展**系统的基石

想象成汽车制造工厂的装配线——零件（依赖项）由传送带（容器）自动送到组装工位（对象），而非工人四处找零件。

- 厨师专注烹饪（业务逻辑）
- 传送带自动送食材（依赖注入）
- 管理员随时更换食材供应商（更换实现）

![QQ_1755338451627.png](https://p0-xtjj-private.juejin.cn/tos-cn-i-73owjymdk6/27a0ffe8544d4807a5e5174c9ca296a8~tplv-73owjymdk6-jj-mark-v1:0:0:0:0:5o6Y6YeR5oqA5pyv56S-5Yy6IEAg6KW_57u0:q75.awebp?policy=eyJ2bSI6MywidWlkIjoiMTM3MjY1NDM4OTQzMzQ5NiJ9&rk3s=f64ab15b&x-orig-authkey=f32326d3454f2ac7e96d3d06cdbb035152127018&x-orig-expires=1766861773&x-orig-sign=cnDKKN2ytq2IrNTQQgQUPvHS5og%3D)

```go
// 传统紧耦合写法（难测试、难替换）
type UserController struct {
    userService *UserService // 直接依赖具体实现
}

// DI解耦写法（依赖接口）
type UserController struct {
    userService UserServiceInterface // 依赖抽象
}


// 声明依赖关系
func NewUserController(svc UserServiceInterface) *UserController {
    return &UserController{userService: svc}
}

// 容器自动组装（Wire工具）
var ProviderSet = wire.NewSet(
    NewUserService,    // 创建Service
    NewUserController, // 创建Controller（自动注入Service）
)

// ### 1. 分层注入
// data层（基础设施）
wire.NewSet(NewMySQL, NewRedis)

// biz层（业务逻辑）
wire.NewSet(NewUserRepo, NewOrderRepo)

// service层（服务实现）
wire.NewSet(NewUserService, NewOrderService)

// server层（传输层）
wire.NewSet(NewHTTPServer, NewGRPCServer)

// ### 2. 接口绑定
// 定义接口
type UserRepo interface {
    GetUser(id int) (*User, error)
}

// 实现接口
type mysqlUserRepo struct {
    db *gorm.DB
}

// 容器绑定
var UserRepoSet = wire.NewSet(
    NewMySQLUserRepo,      // 提供具体实现
    wire.Bind(new(UserRepo), new(*mysqlUserRepo)), // 绑定接口
)
```

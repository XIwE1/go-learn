package main

import (
	"errors"
	"fmt"
)

// Go 将错误视为必须显式处理的值。
// 错误必须显式检查和处理

func main() {
	err := processFile("data.txt")
	// 最常见的是函数调用后立即检查错误
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	}

	user, err := dbQuery("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		// %w 用于包装原始错误，保留上下文
		return nil, fmt.Errorf("获取用户失败: %w", err)
	}

	user, err := findUser(2000)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			fmt.Println("用户未找到")
		case errors.Is(err, ErrInvalidInput):
			fmt.Println("无效的用户ID")
		default:
			fmt.Printf("未知错误: %v\n", err)
		}
		return
	}

	// 使用 errors.Is 和 errors.As 进行错误比较
	var netErr NetworkError
	if errors.As(err, &netErr) {
		fmt.Printf("网络错误，状态码: %d\n", netErr.StatusCode)
	} else {
		fmt.Printf("其他错误: %v\n", err)
	}

	// 收集多个错误并一起返回
	var errors []string
	
	if user["name"] == "" {
		errors = append(errors, "姓名是必需的")
	}
	if user["email"] == "" {
		errors = append(errors, "邮箱是必需的")
	}
	if age, ok := user["age"].(int); ok && age < 0 {
		errors = append(errors, "年龄必须为正数")
	}
}

// 哨兵错误 表示特定错误条件的**预定义错误值**
var (
	ErrNotFound          = errors.New("资源未找到")
	ErrPermissionDenied  = errors.New("权限被拒绝")
	ErrInvalidInput      = errors.New("无效输入")
)
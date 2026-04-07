package type

import (
	"fmt"
)

// 空接口 (interface{}) 接受任何类型
func describeType(value interface{}) string{
	switch type := value.(type) {
	case int:
		return fmt.Sprintf("int: %d", value)
	case string:
		return fmt.Sprintf("string: %s", value)
	case bool:
		return fmt.Sprintf("bool: %t", value)
	default:
		return fmt.Sprintf("unknown type: %T", value)
	}
}

func processValue(value interface{}) string {
    // 带逗号 ok 惯用法的类型断言
    if str, ok := value.(string); ok {
        return str
    }
	if arr, ok := value.([]int); ok {
        return fmt.Sprintf("%v", arr)
    }
    return fmt.Sprintf("unknown type: %T", value)
}
func main() {
	fmt.Println("Hello, World!")
}